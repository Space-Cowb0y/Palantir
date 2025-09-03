FROM golang:1.22 AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/api-agent ./cmd/api-agent

FROM gcr.io/distroless/static:nonroot
WORKDIR /app
COPY --from=build /out/api-agent /app/api-agent
USER nonroot:nonroot
EXPOSE 8082
ENTRYPOINT ["/app/api-agent"]
