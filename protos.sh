mkdir proto/admin proto/common proto/customdim proto/indexing proto/indexingcore proto/lists proto/services proto/serving proto/sstatus
protoc -I proto/protoc-gen-openapiv2 -I proto -I proto/google/api -I proto/status/ \
    --go_out=paths=source_relative:./proto/status \
    --go-grpc_out=paths=source_relative:./proto/status \
    proto/status.proto
protoc -I proto/protoc-gen-openapiv2 -I proto -I proto/google/api -I proto/status/ \
    --go_out=paths=source_relative:./proto/admin \
    --go-grpc_out=paths=source_relative:./proto/admin \
    proto/admin*.proto
protoc -I proto/protoc-gen-openapiv2 -I proto -I proto/google/api -I proto/status/ \
    --go_out=paths=source_relative:./proto/common \
    --go-grpc_out=paths=source_relative:./proto/common \
    proto/attribute.proto proto/common.proto
protoc -I proto/protoc-gen-openapiv2 -I proto -I proto/google/api -I proto/status/ \
    --go_out=paths=source_relative:./proto/services \
    --go-grpc_out=paths=source_relative:./proto/services \
    proto/services.proto
protoc -I proto/protoc-gen-openapiv2 -I proto -I proto/google/api -I proto/status/ \
    --go_out=paths=source_relative:./proto/serving \
    --go-grpc_out=paths=source_relative:./proto/serving \
    proto/serving.proto
protoc -I proto/protoc-gen-openapiv2 -I proto -I proto/google/api -I proto/status/ \
    --go_out=paths=source_relative:./proto/indexing \
    --go-grpc_out=paths=source_relative:./proto/indexing \
    proto/indexing.proto
protoc -I proto/protoc-gen-openapiv2 -I proto -I proto/google/api -I proto/status/ \
    --go_out=paths=source_relative:./proto/indexingcore \
    --go-grpc_out=paths=source_relative:./proto/indexingcore \
    proto/indexing_core.proto
protoc -I proto/protoc-gen-openapiv2 -I proto -I proto/google/api -I proto/status/ \
    --go_out=paths=source_relative:./proto/lists \
    --go-grpc_out=paths=source_relative:./proto/lists \
    proto/list_documents.proto
protoc -I proto/protoc-gen-openapiv2 -I proto -I proto/google/api -I proto/status/ \
    --go_out=paths=source_relative:./proto/customdim \
    --go-grpc_out=paths=source_relative:./proto/customdim \
    proto/custom_dim.proto