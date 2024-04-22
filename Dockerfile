FROM golang:1.22.1

ENV GOPROXY=direct

WORKDIR /app

# Copia el código fuente al contenedor
COPY . .

# Instala las dependencias (si es necesario)
# RUN go mod download

# Compila la aplicación
RUN go build -o main .

# Expone el puerto en el que la aplicación se ejecutará
EXPOSE 8080

# Especifica el comando de inicio
CMD ["./main"]