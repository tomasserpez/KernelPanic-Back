// entities/agent.go
package entities

type AgentInfoResponse struct {
	Data Agent `json:"data"`
}

type Agent struct {
	ID           int    `json:"-" db:"id"` // Primary Key in DB (Optional for API)
	Name         string `json:"name" db:"name"`
	Symbol       string `json:"symbol" db:"symbol"`
	Headquarters string `json:"headquarters" db:"headquarters"`
	Credits      int    `json:"credits" db:"credits"`
	Token        string `json:"-" db:"token"` // Not part of the API response, used for internal storage

}

type RegisterResponse struct {
	Data RegisterData `json:"data"`
}

type RegisterData struct {
	Token string `json:"token"`
	Agent Agent  `json:"agent"`
}
