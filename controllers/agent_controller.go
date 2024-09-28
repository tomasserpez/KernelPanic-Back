package controllers

import (
	"KernelPanic-Back/db"
	"KernelPanic-Back/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AgentController struct {
	db           *db.DB
	agentService *services.AgentService
}

func NewAgentController(db *db.DB, agentService *services.AgentService) *AgentController {
	return &AgentController{db: db, agentService: agentService}
}

// RegisterAgent maneja la creación de un nuevo agente
func (ac *AgentController) RegisterAgent(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Faction  string `json:"faction" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	regResp, err := ac.agentService.RegisterAgent(strings.ToUpper(input.Username), strings.ToUpper(input.Faction))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Guardar el agente en la base de datos
	regResp.Agent.Token = regResp.Token
	if err := ac.db.SaveAgent(&regResp.Agent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": regResp.Token,
		"agent": regResp.Agent,
	})
}

// GetAgents devuelve la lista de agentes
func (ac *AgentController) GetAgents(c *gin.Context) {
	agents, err := ac.db.GetAgentsAndTokens()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, agents)
}

// GetAgentByName maneja la obtención de un agente por su nombre
func (ac *AgentController) GetAgentByName(c *gin.Context) {
	name := c.Param("name")

	agent, err := ac.db.GetAgentByName(name)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, agent)
}

// GetAgentByToken maneja la obtención de un agente por su token
func (ac *AgentController) GetAgentByToken(c *gin.Context) {
	token := c.Param("token")

	agentInfo, err := ac.agentService.GetAgentInfo(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, agentInfo)
}
