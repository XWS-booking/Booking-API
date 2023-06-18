FROM golang:alpine
WORKDIR ./api-gateway
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o gateway .
EXPOSE 8085
EXPOSE 8086
CMD ["./gateway"]