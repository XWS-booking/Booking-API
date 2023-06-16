FROM golang:alpine
WORKDIR ./recommendation-service
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o recommendation-app .
EXPOSE 9000
CMD ["./recommendation-app"]