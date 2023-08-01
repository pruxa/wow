FROM golang:1.20
WORKDIR /go/src/github.com/pruxa/wow

COPY ./ ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build  -o /wow ./cmd

EXPOSE 8080:8080

# Run
CMD ["/wow"]