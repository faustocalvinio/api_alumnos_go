# Usa una imagen base que contenga el entorno de ejecución de Go
FROM golang:latest

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /go/src/app

# Copia el contenido del directorio actual al directorio de trabajo del contenedor
COPY . .

# Descarga e instala cualquier dependencia si es necesario (utiliza go.mod y go.sum si estás usando módulos)
RUN go mod download

# !Compila el código de Go en un binario ejecutable
# !RUN go build -o main .

# Expone el puerto en el que se ejecuta tu aplicación
EXPOSE 8080


# Comando por defecto para ejecutar tu aplicación cuando se inicie el contenedor
CMD ["go","run","."]
