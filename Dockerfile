FROM golang:1.18 as builder

WORKDIR /go/src/app
COPY main.go .

RUN go mod init
RUN go get -d -v ./...
RUN go vet -v
RUN go test -v

RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static

COPY --from=builder /go/bin/app /
CMD ["/app"]
