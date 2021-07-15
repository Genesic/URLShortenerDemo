FROM golang:1.14 as build

COPY . /app
WORKDIR /app
RUN chmod +x /app/wait-for-it.sh
RUN go mod tidy && go mod vendor
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/main main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/migrate /app/service/database/cmd/migrate.go
RUN chmod +x /app/start.sh

FROM alpine as final
WORKDIR /app
RUN apk add --no-cache bash
COPY --from=build /app /app