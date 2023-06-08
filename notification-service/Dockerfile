FROM golang:alpine
WORKDIR ./notification-service
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o booking-notification-app .
EXPOSE 9000
EXPOSE 8080
CMD ["./booking-notification-app"]