protoc --go_out=./reservation --go_opt=paths=source_relative \
--go-grpc_out=./reservation --go-grpc_opt=paths=source_relative \
reservation_service.proto