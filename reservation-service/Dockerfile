FROM golang:alpine
WORKDIR ./reservation-service
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o booking-reservation-app .
EXPOSE 9000
CMD ["./booking-reservation-app"]