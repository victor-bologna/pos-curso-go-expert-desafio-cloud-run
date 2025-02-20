
FROM golang:1.22 AS build-stage

WORKDIR /app

COPY . .

RUN go mod download

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o weather-zipcode ./cmd/server/

FROM build-stage AS run-test-stage
RUN go test -v ./...

FROM alpine:latest

WORKDIR /root/

COPY --from=build-stage /app/weather-zipcode .

CMD ["./weather-zipcode"]