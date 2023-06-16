protoc --go_out=./recommendation --go_opt=paths=source_relative \
--go-grpc_out=./recommendation --go-grpc_opt=paths=source_relative \
recommendation_service.proto