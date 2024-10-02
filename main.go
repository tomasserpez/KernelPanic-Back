package main

import (
	"KernelPanic-Back/controllers"
	"KernelPanic-Back/db"
	"KernelPanic-Back/services"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"log"
)

func main() {
	// Inicializamos la DB
	database := db.NewDB("./agents.db")
	defer database.Conn.Close()

	restyClient := resty.New()

	//Inicializamos el servicio de agente
	agentService := services.NewAgentService(restyClient)
	contractsService := services.NewContractService(restyClient)

	// Configuramos el router de GIN
	router := gin.Default()
	agentController := controllers.NewAgentController(database, agentService)
	contractsController := controllers.NewContractsController(database, contractsService)

	// Definimos las rutas
	//Agent
	router.POST("/agents/register", agentController.RegisterAgent)
	router.GET("/agents", agentController.GetAgents)
	router.GET("/agents/name/:name", agentController.GetAgentByName)
	router.GET("/agents/token/:token", agentController.GetAgentByToken)

	//Contracts
	router.GET("/:agentName/contracts", contractsController.GetAgentContractsByName)
	router.POST("/:agentName/contracts/:contractId/accept", contractsController.AcceptAgentContract)

	// Iniciamos el servidor
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
