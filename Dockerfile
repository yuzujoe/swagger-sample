FROM golang:1.13 as builder

ENV GO111MODULE=on

WORKDIR /app/

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o build ./main.go

FROM alpine

RUN apk update \
  && apk --no-cache add tzdata \
  && apk add --no-cache ca-certificates

ENV GIN_MODE=release
COPY --from=builder /app/build /app/build

ENTRYPOINT ["app/build"]
