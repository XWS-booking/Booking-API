cd reservation-service/proto &&
protoc --go_out=./reservation --go_opt=paths=source_relative --go-grpc_out=./reservation --go-grpc_opt=paths=source_relative reservation_service.proto &&
cd ../../auth-service/proto &&
protoc --go_out=./auth --go_opt=paths=source_relative --go-grpc_out=./auth --go-grpc_opt=paths=source_relative auth_service.proto &&
cd ../../accomodation-service/proto &&
protoc --go_out=./accomodation --go_opt=paths=source_relative --go-grpc_out=./accomodation --go-grpc_opt=paths=source_relative accomodation_service.proto &&
cd ../../notification-service/proto &&
protoc --go_out=./notification --go_opt=paths=source_relative --go-grpc_out=./notification --go-grpc_opt=paths=source_relative notification_service.proto &&
cd ../../rating-service/proto &&
protoc --go_out=./rating --go_opt=paths=source_relative --go-grpc_out=./rating --go-grpc_opt=paths=source_relative rating_service.proto
cd ../../recommendation-service/proto &&
protoc --go_out=./recommendation --go_opt=paths=source_relative --go-grpc_out=./recommendation --go-grpc_opt=paths=source_relative recommendation_service.proto &&
cd ../../api-gateway/proto &&
protoc --go_out=./gateway --go_opt=paths=source_relative --go-grpc_out=./gateway --go-grpc_opt=paths=source_relative --grpc-gateway_out=./gateway --grpc-gateway_opt=paths=source_relative gateway_service.proto