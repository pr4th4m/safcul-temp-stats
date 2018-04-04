# # Build image
# docker build -t "safcul-temp-stats:dev" .
#
# # Run container
# docker run --rm --name safcul-temp-stats -it safcul-temp-stats:dev
FROM golang:1.10.0-stretch

# Copy binary and config file
WORKDIR /go/src/github.com/pratz/safcul-temp-stats
COPY . .

# Run unit tests
RUN go test

# Build binary
RUN go build -o stats *.go

# Run app
ENTRYPOINT ["./stats"]
