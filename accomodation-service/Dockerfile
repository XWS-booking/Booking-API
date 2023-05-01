FROM golang:alpine
WORKDIR ./accomodation-service
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o booking-accomodation-app .
EXPOSE 9000
CMD ["./booking-accomodation-app"]