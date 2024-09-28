package main

import (
	"KernelPanic-Back/controllers"
	"KernelPanic-Back/db"
	"KernelPanic-Back/services"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Inicializamos la DB
	database := db.NewDB("./agents.db")
	defer database.Conn.Close()

	//Inicializamos el servicio de agente
	agentService := services.NewAgentService()

	// Configuramos el router de GIN
	router := gin.Default()
	agentController := controllers.NewAgentController(database, agentService)

	// Definimos las rutas
	router.POST("/agents/register", agentController.RegisterAgent)
	router.GET("/agents", agentController.GetAgents)
	router.GET("/agents/name/:name", agentController.GetAgentByName)
	router.GET("/agents/token/:token", agentController.GetAgentByToken)

	// Iniciamos el servidor
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
