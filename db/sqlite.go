// db/sqlite.go
package db

import (
	"KernelPanic-Back/entities"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strings"
)

// La DB está envuelta en una conección SQLite3
type DB struct {
	Conn *sql.DB
}

// NewDB inicializa una nueva conección a la base de datos
func NewDB(filepath string) *DB {
	conn, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatalf("Falló al conectarse a la DB: %v", err)
	}
	query := `
	CREATE TABLE IF NOT EXISTS agents (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    accountId TEXT,
	    symbol TEXT,
	    headquarters TEXT,
	    credits INTEGER,
	    token TEXT,
	    firebaseUid TEXT
  );`
	if _, err := conn.Exec(query); err != nil {
		log.Fatalf("Fallo al crear la tabla agente: %v", err)
	}
	return &DB{Conn: conn}
}

// SaveAgent guarda el agente en la base de datos SQLite
func (db *DB) SaveAgent(agent *entities.Agent) error {
	query := `INSERT INTO agents (accountId, symbol, headquarters, credits, token, firebaseUid) VALUES (?,?,?,?,?,?)`
	_, err := db.Conn.Exec(query, agent.AccountId, agent.Symbol, agent.Headquarters, agent.Credits, agent.Token, agent.FirebaseUid)
	return err
}

// GetAgentByName devuelve el agente por su nombre desde la base de datos
func (db *DB) GetAgentByName(name string) (*entities.Agent, error) {
	name = strings.ToUpper(name)
	query := `SELECT id, accountId, symbol, headquarters, credits, token, firebaseUid FROM agents WHERE symbol = ?`
	row := db.Conn.QueryRow(query, name)
	var agent entities.Agent

	err := row.Scan(&agent.ID, &agent.AccountId, &agent.Symbol, &agent.Headquarters, &agent.Credits, &agent.Token, &agent.FirebaseUid)
	if err != nil {
		return nil, err
	}
	if agent.ID == 0 {
		fmt.Errorf("AGENTE NO ENCONTRADO EN DB")
	}

	return &agent, nil
}

// GetAgentsAndTokens Listar agentes y sus tokens
func (db *DB) GetAgentsAndTokens() (*entities.Agents, error) {
	query := `SELECT symbol, token FROM agents`
	rows, err := db.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var agents entities.Agents
	var agent entities.Agent
	for rows.Next() {
		err := rows.Scan(&agent.Symbol, &agent.Token)
		if err != nil {
			return nil, err
		}
		agents = append(agents, agent)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &agents, nil
}

func (db *DB) GetAgentsAndTokensForUser(uid string) (*entities.Agents, error) {
	query := `SELECT symbol, token FROM agents WHERE firebaseUid = ?`
	rows, err := db.Conn.Query(query, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var agents entities.Agents
	var agent entities.Agent
	for rows.Next() {
		err := rows.Scan(&agent.Symbol, &agent.Token)
		if err != nil {
			return nil, err
		}
		agents = append(agents, agent)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &agents, nil
}
