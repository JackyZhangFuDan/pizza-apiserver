FROM golang:1.17 as build
WORKDIR /go/src/pizza-apiserver
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=build /go/src/pizza-apiserver/pizza-apiserver /
ENTRYPOINT ["/pizza-apiserver"]