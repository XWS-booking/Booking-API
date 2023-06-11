protoc --go_out=./notification --go_opt=paths=source_relative \
--go-grpc_out=./notification --go-grpc_opt=paths=source_relative \
notification_service.proto