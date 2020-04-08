FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/app/
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o /go/bin/app

##############################

FROM scratch
COPY --from=builder /go/bin/app /go/bin/app
CMD ["/go/bin/app"]
