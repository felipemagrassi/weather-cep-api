FROM golang:1.22 as build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o weather-cep-api cmd/server/main.go

FROM scratch
WORKDIR /app
COPY --from=build .env .
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/weather-cep-api .
ENTRYPOINT ["./weather-cep-api"]

