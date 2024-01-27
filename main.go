// main.go

package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"vectara.com/public/proto/services"
	"vectara.com/public/proto/serving"
)

var (
	conn           *grpc.ClientConn
	addr           = flag.String("addr", "serving.vectara.io:443", "the address to connect to")
	customerID     = "2523211369" // Replace with your actual customer ID
	int_customerID = 2523211369
	xAPIKey        = "zwt_lmUmaZXsE_PkhRRjH7ryw8a9dNXYgDXtavC1cw" // Replace with your actual API key
	corpusID       = "12"                                         // Replace with your actual corpus ID
	int_corpusID   = 12
	Lambda         = .025
	certFile       = "key/vectara.cer" // Replace with your actual certificate file path
	serverAddr     = ":5050"
	shutdownSec    = 10
	htmlContent    string
	numResults     = 5
	maxResults     = 10
	samantics      = 0
)

type QueryPayload struct {
	SearchQuery string `json:"searchQuery"`
}

type QueryResponse struct {
	Result string `json:"result"`
}

func main() {
	// Initialize gRPC client
	var err error
	conn, err = initializeGRPCClient()
	if err != nil {
		log.Fatalf("Failed to initialize gRPC client: %v", err)
	}

	// Register HTTP handler
	http.HandleFunc("/query", queryHandler)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "5050"
	}

	log.Printf("Server listening on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
func initializeGRPCClient() (*grpc.ClientConn, error) {
	// Load TLS credentials
	creds, err := credentials.NewClientTLSFromFile(certFile, "")
	if err != nil {
		log.Printf("Error loading TLS credentials: %v", err)
		return nil, fmt.Errorf("failed to load TLS credentials: %v", err)
	}

	// Set up a connection to the server
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(tokenAuth{
		customerID: customerID,
		xAPIKey:    xAPIKey,
		corpusID:   corpusID,
	}))
	if err != nil {
		log.Printf("Error connecting to gRPC server: %v", err)
		return nil, fmt.Errorf("did not connect: %v", err)
	}

	return conn, nil
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		// Respond to preflight requests
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Add CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	var payload QueryPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Error decoding JSON payload", http.StatusBadRequest)
		return
	}

	// Access the searchQuery from the payload
	searchQuery := payload.SearchQuery
	log.Printf("searchQuery?" + searchQuery)
	result, err := callVectara(searchQuery)
	if err != nil {
		log.Printf("Error calling Vectara: %v", err)
		http.Error(w, fmt.Sprintf("Error calling Vectara: %v", err), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(result)
	if err != nil {
		log.Printf("Error marshaling JSON response: %v", err)
		http.Error(w, fmt.Sprintf("Error marshaling JSON response: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

type tokenAuth struct {
	token      string
	apiKey     string
	customerID string
	xAPIKey    string
	corpusID   string
}

func (t tokenAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	headers := map[string]string{
		"customer-id": t.customerID,
		"x-api-key":   t.xAPIKey,
		"corpus-id":   t.corpusID,
	}
	return headers, nil
}

func (tokenAuth) RequireTransportSecurity() bool {
	return true
}

// CallVectara function
func callVectara(searchString string) (*serving.ResponseSet, error) {
	flag.Parse()

	rlambda := serving.LinearInterpolation{
		Lambda: float32(Lambda),
	}

	corpKey := serving.CorpusKey{
		CustomerId:                 uint32(int_customerID),
		CorpusId:                   uint32(int_corpusID),
		Semantics:                  serving.CorpusKey_Semantics(samantics),
		LexicalInterpolationConfig: &rlambda,
	}

	summaryReq := serving.SummarizationRequest{
		MaxSummarizedResults: uint32(numResults),
		ResponseLang:         "auto",
		SummarizerPromptName: "vectara-summary-ext-v1.2.0",
	}

	query := serving.QueryRequest{
		Query:      searchString,
		Start:      0,
		NumResults: uint32(maxResults),
		CorpusKey:  []*serving.CorpusKey{&corpKey},
		Summary:    []*serving.SummarizationRequest{&summaryReq},
	}

	client := services.NewQueryServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()

	response, err := client.Query(ctx, &serving.BatchQueryRequest{Query: []*serving.QueryRequest{&query}})
	if err != nil {
		return nil, fmt.Errorf("could not get answer: %v", err)
	}
	if len(response.ResponseSet) > 0 {
		// Return the proto response directly
		return response.ResponseSet[0], nil
	}

	return nil, fmt.Errorf("response structure does not contain the expected data")
}
