protoc --go_out ./gateway --go_opt paths=source_relative \
--go-grpc_out ./gateway --go-grpc_opt paths=source_relative \
--grpc-gateway_out ./gateway --grpc-gateway_opt paths=source_relative \
gateway_service.proto