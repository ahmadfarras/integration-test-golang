FROM golang:1.21.1 as build

ENV GOPATH=/go

WORKDIR /app

COPY . .

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/myapi main.go

FROM alpine:3.18.3

WORKDIR /app

COPY --from=build /app/myapi /app/myapi

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
RUN chmod +x /app/myapi
USER appuser

EXPOSE 8080

HEALTHCHECK --interval=60s --timeout=3s --start-period=5s --retries=3 CMD [ "wget", "-q", "http://localhost:8080/healthz", "-O", "-" ]

ENTRYPOINT ["/app/myapi"]