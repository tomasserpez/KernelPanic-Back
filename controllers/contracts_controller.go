package controllers

import (
	"KernelPanic-Back/db"
	"KernelPanic-Back/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ContractsController struct {
	db               *db.DB
	contractsService *services.ContractsService
}

func NewContractsController(db *db.DB, contractsService *services.ContractsService) *ContractsController {
	return &ContractsController{
		db:               db,
		contractsService: contractsService,
	}
}

func (cCont *ContractsController) GetAgentContractsByName(ginContext *gin.Context) {

	agentName := ginContext.Param("agentName")

	agent, err := cCont.db.GetAgentByName(agentName)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			ginContext.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	contracts, err := cCont.contractsService.GetContractsForAgent(agent)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginContext.JSON(http.StatusOK, contracts)
}

func (cCont *ContractsController) AcceptAgentContract(ginContext *gin.Context) {
	//Extraemos los parametros de la URL
	contractId := ginContext.Param("contractId")
	agentName := ginContext.Param("agentName")

	//Buscamos el agente en la db
	agent, err := cCont.db.GetAgentByName(agentName)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			ginContext.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	//Aceptamos el contrato
	respData, err := cCont.contractsService.AcceptContract(agent, contractId)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ginContext.JSON(http.StatusOK, respData)

}
