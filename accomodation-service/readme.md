protoc --go_out=./accomodation --go_opt=paths=source_relative \
--go-grpc_out=./accomodation --go-grpc_opt=paths=source_relative \
accomodation_service.proto