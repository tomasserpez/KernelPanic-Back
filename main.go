package main

import (
	"KernelPanic-Back/db"
	"KernelPanic-Back/services"
	"fmt"
	"log"
)

func main() {
	// Inisializamos la DB
	database := db.NewDB("./agents.db")
	defer database.Conn.Close()

	//Inicializamos el servicio de agente
	agentService := services.NewAgentService()

	username := "PREGONI"
	//faction := "COSMIC"
	//
	//regResp, err := agentService.RegisterAgent(username, faction)
	//if err != nil {
	//	log.Fatalf("Error registrando el agente: %v", err)
	//}
	//regResp.Agent.Token = regResp.Token
	//err = database.SaveAgent(&regResp.Agent)
	//if err != nil {
	//	log.Fatalf("Error guardando al agente en la base de datos: %v", err)
	//}
	//fmt.Printf("Agente nuevo registrado con el token: %s\n", regResp.Token)
	//fmt.Printf("Detalles del agente: %+v\n", regResp.Agent)

	//Ejemplo
	savedAgent, err := database.GetAgentByName(username)
	if err != nil {
		log.Fatalf("Error al hacer un fetching del agente en la DB: %v", err)
	}

	agentInfo, err := agentService.GetAgentInfo(savedAgent.Token)
	if err != nil {
		log.Fatalf("Error haciendo fetching de la información del agente: %v", err)
	}

	fmt.Printf("Info del agente: %+v\n", agentInfo)
}
