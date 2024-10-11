# Imagen base
FROM golang:1.19-alpine

# Configurar el directorio de trabajo
WORKDIR /app

# Copiar archivos del proyecto
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Construir la aplicación
RUN go build -o /go-image-text cmd/main.go

# Exponer el puerto
EXPOSE 8080

# Ejecutar la aplicación
CMD ["/go-image-text"]
