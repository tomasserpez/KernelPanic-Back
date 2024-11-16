package services

import (
	"KernelPanic-Back/entities"
	"fmt"
	"github.com/go-resty/resty/v2"
)

// AgentService provee metodos de interacción con la API
type AgentService struct {
	client *resty.Client
}

// NewAgentService crea una nueva instancia de agentService
func NewAgentService(client *resty.Client) *AgentService {
	return &AgentService{
		client: client,
	}
}

// RegisterAgent registra un nuevo agente con el endpoint de SpaceTraders
func (s *AgentService) RegisterAgent(username string, faction string) (*entities.RegisterData, error) {
	response := new(entities.RegisterResponse)
	resp, err := s.client.R().
		SetBody(map[string]string{
			"symbol":  username,
			"faction": faction,
		}).
		SetResult(response).
		Post(BASEURL + "/register")

	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error al registrar el agente: %s", resp.Status())
	}
	return &response.Data, nil
}

// GetAgentInfo trae y procesa toda la información del agente usando un token de proveedor
func (s *AgentService) GetAgentInfo(token string) (*entities.Agent, error) {
	response := new(entities.AgentInfoResponse)
	resp, err := s.client.R().
		SetAuthToken(token).
		SetResult(response).
		Get(BASEURL + "/my/agent")
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("error al procesar la información del agente: %s", resp.Status())
	}
	return &response.Data, nil
}
