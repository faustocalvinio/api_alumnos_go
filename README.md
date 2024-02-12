# Api Alumnos en Go

> [!NOTE]
> Primera API basica en Go desarrollada a modo de aprendizaje. Rutas disponibles:

- /alumnos (GET)
- /alumnos/nuevo (POST)
    parametros de url: 
        nombre,	apellido, edadStr, carrera,
- /alumnos/set-initials (GET)

#### env variables
```bash
SERVIDOR_BD=mi_servidor_bd
PUERTO=8080
```

## Paquetes utilizados:
1) fmt -> formatter
2) github.com/gin-gonic/gin -> framework ligero para manejar peticiones http
3) github.com/joho/godotenv -> paquete para leer variables de entorno en Go
4) net/http  -> HTTP client and server implementations.
4) strconv -> conversions to and from string representations of basic data types.
5) os -> platform-independent interface to operating system functionality

## To-do
1) No carga localhost, hay que usar 127.0.1 en su lugar