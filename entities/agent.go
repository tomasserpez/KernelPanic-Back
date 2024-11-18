// entities/agent.go
package entities

type AgentInfoResponse struct {
	Data Agent `json:"data"`
}

type Agents []Agent

type Agent struct {
	ID              int    `json:"-" db:"id"` // Primary Key in DB (Optional for API)
	AccountId       string `json:"accountId" db:"accountId"`
	Symbol          string `json:"symbol" db:"symbol"`
	Headquarters    string `json:"headquarters" db:"headquarters"`
	Credits         int    `json:"credits" db:"credits"`
	StartingFaction string `json:"startingFaction" db:"startingFaction"`
	ShipCount       int    `json:"shipCount" db:"shipCount"`
	Token           string `json:"token" db:"token"`             // Not part of the API response, used for internal storage
	FirebaseUid     string `json:"firebaseUid" db:"firebaseUid"` // same as above

}

type RegisterResponse struct {
	Data RegisterData `json:"data"`
}

type RegisterData struct {
	Token    string   `json:"token"`
	Agent    Agent    `json:"agent"`
	Contract Contract `json:"contract"`
}
