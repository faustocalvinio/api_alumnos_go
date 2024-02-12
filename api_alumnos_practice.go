package main

// ! IMPORTACIONES NECESARIAS
import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// ! TIPADO DE UN ALUMNO
type alumno struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad     int    `json:"edad"`
	Carrera  string `json:"carrera"`
}

// ! ARREGLO DE ALUMNOS UNR (UNIVERSIDAD NACIONAL DE RIOJA)
var alumnosUNR = []alumno{
	{ID: 1, Nombre: "Juan", Apellido: "Perez", Edad: 20, Carrera: "Ingenieria"},
	{ID: 2, Nombre: "Maria", Apellido: "Gomez", Edad: 22, Carrera: "Medicina"},
	{ID: 3, Nombre: "Pedro", Apellido: "Rodriguez", Edad: 24, Carrera: "Derecho"},
	{ID: 4, Nombre: "Ana", Apellido: "Lopez", Edad: 26, Carrera: "Arquitectura"},
}

// ! FUNCION PARA RETORNAR TODOS LOS ALUMNOS
func getAllAlumns(c *gin.Context) {
	// ?leer un x-token enviado mediante cookies del cliente en un futuro
	cookieValue, err := c.Cookie("x-token")
	// ?imprimir el valor de esa cookie en caso de que no haya error
	if err == nil {
		println("Valor de la cookie x-token: ", cookieValue)
	}
	// ?enviar un json como respuesta
	c.JSON(http.StatusAccepted, gin.H{"ok": true, "alumnosUNR": alumnosUNR})
}

// ! FUNCION PARA RETORNAR TODOS LOS ALUMNOS MEDIANTE URL PARAMS
func addAlumno(c *gin.Context) {
	// ?leer un x-token enviado mediante cookies del cliente en un futuro
	cookieValue, err := c.Cookie("x-token")

	if err == nil {
		fmt.Println("Valor de la cookie x-token: ", cookieValue)
	}
	// ?imprimir el valor de esa cookie en caso de que no haya error

	queryValue := c.Query("q")
	if queryValue != "" {
		fmt.Println("Valor del parámetro 'q':", queryValue)
	}
	
	nombre := c.Query("nombre")
	apellido := c.Query("apellido")
	edadStr := c.Query("edad")
	carrera := c.Query("carrera")

	edad, err := strconv.Atoi(edadStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Edad debe ser un número entero"})
		return
	}

	newAlumno := alumno{
		ID:       len(alumnosUNR) + 1, // ID único
		Nombre:   nombre,
		Apellido: apellido,
		Edad:     edad,
		Carrera:  carrera,
	}

	alumnosUNR = append(alumnosUNR, newAlumno)

	c.JSON(http.StatusAccepted, gin.H{"ok": true, "alumnoNuevo": newAlumno})

	fmt.Println("Añadido el nuevo alumno:", newAlumno)
}

// ! FUNCION "SEED" PARA CARGAR ALUMNOS INICIALES
func setInitials(c *gin.Context) {
	alumnosUNR = []alumno{
		{ID: 1, Nombre: "Juan", Apellido: "Perez", Edad: 20, Carrera: "Ingenieria"},
		{ID: 2, Nombre: "Maria", Apellido: "Gomez", Edad: 22, Carrera: "Medicina"},
		{ID: 3, Nombre: "Pedro", Apellido: "Rodriguez", Edad: 24, Carrera: "Derecho"},
		{ID: 4, Nombre: "Ana", Apellido: "Lopez", Edad: 26, Carrera: "Arquitectura"},
	}
	c.JSON(http.StatusAccepted, gin.H{"ok": true, "alumnos": alumnosUNR})
}

// ! FUNCION PARA CARGAR VARIABLES DE ENTORNO CON GODOTENV LIBRARY

func loadEnvVars() error {
	// ?cargar las variables de entorno desde el archivo .env
	err := godotenv.Load()
	// ?si hay algun error imprimirlo en consola
	if err != nil {
		return fmt.Errorf("Error al cargar el archivo .env: %v", err)
	}
	return nil
}

// ! FUNCION PARA CONFIGURAR EL CORS

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// ! FUNCION GENERAL
func main() {

	loadEnvVars()

	servidorBD := os.Getenv("SERVIDOR_BD")
	if servidorBD == "" {
		fmt.Println("La variable de entorno SERVIDOR_BD no está definida.")
	} else {
		fmt.Println("Valor de SERVIDOR_BD:", servidorBD)
	}

	// ? Obtener y mostrar el valor de variable de entorno PORT
	puerto := os.Getenv("PORT")
	if puerto == "" {
		fmt.Println("La variable de entorno PUERTO no está definida.")
	} else {
		fmt.Println("Valor de PUERTO:", puerto)
	}
	// ! DEFINICION DEL ROUTER
	router := gin.Default()
	// ! USAR LA CONFIGURACION DEL CORS
	router.Use(CORSMiddleware())
	direccion := fmt.Sprintf("localhost:%s", "8080")

	router.SetTrustedProxies([]string{"127.0.0.1"})
	// ? METODOS GET
	router.GET("/alumnos", getAllAlumns)
	router.GET("/alumnos/set-initials", setInitials)
	// ? METODOS POST
	router.POST("/alumnos/nuevo", addAlumno)
	// ? CORRER EL ROUTER
	router.Run(direccion)
}
