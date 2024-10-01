package main

import (
	"KernelPanic-Back/db"
	"KernelPanic-Back/entities"
	"KernelPanic-Back/services"
	"fmt"
	"log"
)

func main() {
	// Inicializamos la DB
	database := db.NewDB("./agents.db")
	defer database.Conn.Close()
	//Inicializamos los servicios de agente y contratos
	agentService := services.NewAgentService()
	contractsService := services.NewContractService(entities.Agent{}) //le pasamos un agente vacio

	username := "TESTKP4"
	token := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGlmaWVyIjoiVEVTVEtQNCIsInZlcnNpb24iOiJ2Mi4yLjAiLCJyZXNldF9kYXRlIjoiMjAyNC0wOS0wMSIsImlhdCI6MTcyNzc1MjUxNiwic3ViIjoiYWdlbnQtdG9rZW4ifQ.PLGuNBTmlAZmg8XZoG0Ti-_y0qN4o38UUyP67xnHt8xJmBTOQvRtbw_lP6BAbW96RAbOLTW4pMmz1mg10_PRy9b09CT1g15qmLas0gO99GClvcQkCqvvjg2mxkHUv0yVD8nNX4gxWgFvG1Zjcm5d_pb9HaPDmR4FWE77OfZJYRnT17Y1eTs99RbvgZJmk1JKqXJ-157IXNYO6IbdqCEAdknXjSfdRNuxopzs4Y-jBv7rdF0C5RhWoZmHEZCuvAwOO97VziMxpKlvl7-VgtA8jkPLHg1ynLI_CO_GaZLnW-4ZigY8FEfGIm-pBVY5ar3NAndrJPCorMRbZE8zA595fQ"
	//faction := "COSMIC"
	//regResp, err := agentService.RegisterAgent(username, faction)

	resp, err := agentService.GetAgentInfo(token)
	if err != nil {
		log.Fatalf("Error registrando el agente: %v", err)
	}
	resp.Token = token
	err = database.SaveAgent(resp)
	if err != nil {
		log.Fatalf("Error guardando al agente en la base de datos: %v", err)
	}
	fmt.Printf("Agente nuevo registrado con el token: %s\n", resp.Token)
	fmt.Printf("Detalles del agente: %+v\n", *resp)

	//Ejemplo

	savedAgent, err := database.GetAgentByName(username)
	if err != nil {
		log.Fatalf("Error al hacer un fetching del agente en la DB: %v", err)
	}

	contractsService.SetAgent(*savedAgent) //cargamos el agente en el servicio
	contracts, err := contractsService.GetContracts()
	if err != nil {
		log.Fatalf("Error en GET a contratos: %v", err)
	}
	fmt.Printf("Contratos del agente %s: %+v\n", (*savedAgent).Name, contracts)

	agentInfo, err := agentService.GetAgentInfo(savedAgent.Token)
	if err != nil {
		log.Fatalf("Error haciendo fetching de la información del agente: %v", err)
	}

	fmt.Printf("Info del agente: %+v\n", agentInfo)
}
