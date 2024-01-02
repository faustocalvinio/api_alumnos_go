package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"github.com/joho/godotenv"
	"os"
)

type alumno struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad     int    `json:"edad"`
	Carrera  string `json:"carrera"`
}

var alumnosUNR = []alumno{
	{ID: 1, Nombre: "Juan", Apellido: "Perez", Edad: 20, Carrera: "Ingenieria"},
	{ID: 2, Nombre: "Maria", Apellido: "Gomez", Edad: 22, Carrera: "Medicina"},
	{ID: 3, Nombre: "Pedro", Apellido: "Rodriguez", Edad: 24, Carrera: "Derecho"},
	{ID: 4, Nombre: "Ana", Apellido: "Lopez", Edad: 26, Carrera: "Arquitectura"},
}

func getAllAlumns(c *gin.Context) {
	cookieValue, err := c.Cookie("x-token")
	if err == nil {
		println("Valor de la cookie x-token: ", cookieValue)
	}

	c.JSON(http.StatusAccepted, gin.H{"ok": true, "alumnosUNR": alumnosUNR})
}

func addAlumno(c *gin.Context) {
	cookieValue, err := c.Cookie("x-token")
	if err == nil {
		fmt.Println("Valor de la cookie x-token: ", cookieValue)
	}

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

func setInitials(c *gin.Context) {
	alumnosUNR = []alumno{
		{ID: 1, Nombre: "Juan", Apellido: "Perez", Edad: 20, Carrera: "Ingenieria"},
		{ID: 2, Nombre: "Maria", Apellido: "Gomez", Edad: 22, Carrera: "Medicina"},
		{ID: 3, Nombre: "Pedro", Apellido: "Rodriguez", Edad: 24, Carrera: "Derecho"},
		{ID: 4, Nombre: "Ana", Apellido: "Lopez", Edad: 26, Carrera: "Arquitectura"},
	}
	c.JSON(http.StatusAccepted, gin.H{"ok": true, "alumnos": alumnosUNR})
}

func cargarVariablesEnv() error {
    err := godotenv.Load()
    if err != nil {
        return fmt.Errorf("Error al cargar el archivo .env: %v", err)
    }
    return nil
}

func main() {

	cargarVariablesEnv()

	servidorBD := os.Getenv("SERVIDOR_BD")
    if servidorBD == "" {
        fmt.Println("La variable de entorno SERVIDOR_BD no está definida.")
    } else {
        fmt.Println("Valor de SERVIDOR_BD:", servidorBD)
    }

    // Obtener y mostrar el valor de otra variable de entorno
    puerto := os.Getenv("PUERTO")
    if puerto == "" {
        fmt.Println("La variable de entorno PUERTO no está definida.")
    } else {
        fmt.Println("Valor de PUERTO:", puerto)
    }

	router := gin.Default()
	direccion := fmt.Sprintf("localhost:%s", puerto)
	
	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.GET("/alumnos", getAllAlumns)
	router.GET("/alumnos/set-initials", setInitials)

	router.POST("/alumnos/nuevo", addAlumno)
	router.Run(direccion)
}
