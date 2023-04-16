FROM golang:alpine
WORKDIR ./auth-service
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o booking-auth-app .
EXPOSE 9000
CMD ["./booking-auth-app"]