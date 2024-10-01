package services

import (
	"KernelPanic-Back/entities"
	"fmt"
	"github.com/go-resty/resty/v2"
)

// ContractService provee metodos de interacción con la API para contratos
type ContractsService struct {
	Agent  entities.Agent
	client *resty.Client
}

// NewContractService crea una nueva instancia de ContractService
func NewContractService(agent entities.Agent) *ContractsService {
	return &ContractsService{
		Agent:  agent,
		client: resty.New(),
	}
}

// SetAgent permite establecer el agente para el ContractService
func (service *ContractsService) SetAgent(agent entities.Agent) {
	service.Agent = agent
}

// GetAgent devuelve el agente actual del ContractService
func (service *ContractsService) GetAgent() entities.Agent {
	return service.Agent
}

func (service *ContractsService) GetContracts() (*entities.Contracts, error) {
	response := new(entities.ContractsForAgentResponse)
	resp, err := service.client.R().
		SetAuthToken(service.Agent.Token).
		SetResult(response).
		Get(BASEURL + "/my/contracts")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("Error al procesar el pedido de contratos: %s", resp.Status())
	}
	return &response.Data, nil
}
