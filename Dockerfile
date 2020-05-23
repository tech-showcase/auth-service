FROM golang:1.14.2-alpine3.11 AS builder

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/app/
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o /go/bin/app

RUN mkdir -p /log \
    && touch http.log

##############################

FROM scratch
COPY --from=builder /go/bin/app /go/bin/app
COPY --from=builder /log /log
ENTRYPOINT ["/go/bin/app"]
