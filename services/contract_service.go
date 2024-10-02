package services

import (
	"KernelPanic-Back/entities"
	"fmt"
	"github.com/go-resty/resty/v2"
)

// ContractsService provee metodos de interacción con la API para contratos
type ContractsService struct {
	client *resty.Client
}

// NewContractService crea una nueva instancia de ContractService
func NewContractService(client *resty.Client) *ContractsService {
	return &ContractsService{
		client,
	}
}

func (service *ContractsService) GetContractsForAgent(agent *entities.Agent) (*entities.Contracts, error) {
	response := new(entities.ContractsForAgentResponse)
	resp, err := service.client.R().
		SetAuthToken(agent.Token).
		SetResult(response).
		Get(BASEURL + "/my/contracts")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error en respuesta de GET a contratos: %s", resp.Status())
	}
	return &response.Data, nil
}

func (service *ContractsService) AcceptContract(agent *entities.Agent, contractId string) (*entities.ContractOnAcceptData, error) {
	response := new(entities.ContractOnAcceptResponse)

	resp, err := service.client.R().
		SetAuthToken(agent.Token).
		SetResult(response).
		Post(BASEURL + "/my/contracts/" + contractId + "/accept")

	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error al aceptar el contrato: %s", resp.Status())
	}
	return &response.Data, nil
}
