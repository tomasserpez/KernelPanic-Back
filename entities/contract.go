package entities

// ContractsForAgentResponse respuesta de API para el pedido de contratos de un agente
type ContractsForAgentResponse struct {
	Data Contracts `json:"data"`
}

type ContractOnAcceptResponse struct {
	Data ContractOnAcceptData `json:"data"`
}

type ContractOnAcceptData struct {
	Contract Contract `json:"contract"`
	Agent    Agent    `json:"agent"`
}

type Contracts []Contract

type Contract struct {
	ID               string `json:"id" db:"id"` // Primary Key in DB
	FactionSymbol    string `json:"factionSymbol" db:"factionSymbol"`
	Type             string `json:"type" db:"type"`
	Terms            Terms  `json:"terms" db:"terms"`
	Accepted         bool   `json:"accepted" db:"accepted"`
	Fulfilled        bool   `json:"fulfilled" db:"fulfilled"`
	Expiration       string `json:"expiration" db:"expiration"`
	DeadlineToAccept string `json:"deadlineToAccept" db:"deadlineToAccept"`
}

type Terms struct {
	Deadline string          `json:"deadline" db:"deadline"`
	Payment  ContractPayment `json:"payment" db:"payment"`
	Deliver  []ContractCargo `json:"deliver" db:"deliver"`
}

type ContractPayment struct {
	OnAccepted  int `json:"onAccepted" db:"onAccepted"`
	OnFulfilled int `json:"onFulfilled" db:"onFulfilled"`
}

type ContractCargo struct {
	TradeSymbol       string `json:"tradeSymbol" db:"tradeSymbol"`
	DestinationSymbol string `json:"destinationSymbol" db:"destinationSymbol"`
	UnitsRequired     int    `json:"unitsRequired" db:"unitsRequired"`
	UnitsFulfilled    int    `json:"unitsFulfilled" db:"unitsFulfilled"`
}
