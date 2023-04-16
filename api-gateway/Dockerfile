FROM golang:alpine
WORKDIR ./auth-service
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o gateway .
EXPOSE 8085
CMD ["./gateway"]