# Etapa de build
FROM golang:1.22 AS build
WORKDIR /app

# Copie go.mod/go.sum primeiro p/ cache de deps
COPY go.mod go.sum ./
RUN go mod download

# Copie o resto do código
COPY . .

# Compile o binário
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/api-admin ./cmd/api-admin

# Etapa final (imagem mínima)
FROM gcr.io/distroless/static:nonroot
WORKDIR /app
COPY --from=build /out/api-admin /app/api-admin
USER nonroot:nonroot

# Var. que seu main lê
ENV ADMIN_API_PORT=8081
EXPOSE 8081
ENTRYPOINT ["/app/api-admin"]
