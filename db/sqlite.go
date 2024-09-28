// db/sqlite.go
package db

import (
	"KernelPanic-Back/entities"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
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
	    name TEXT,
	    symbol TEXT,
	    headquarters TEXT,
	    credits INTEGER,
	    token TEXT
  );`
	if _, err := conn.Exec(query); err != nil {
		log.Fatalf("Fallo al crear la tabla agente: %v", err)
	}
	return &DB{Conn: conn}
}

// SaveAgent guarda el agente en la base de datos SQLite
func (db *DB) SaveAgent(agent *entities.Agent) error {
	query := `INSERT INTO agents (name, symbol, headquarters, credits, token) VALUES (?,?,?,?,?)`
	_, err := db.Conn.Exec(query, agent.Symbol, agent.Symbol, agent.Headquarters, agent.Credits, agent.Token)
	return err
}

// GetAgentByName devuelve el agente por su nombre desde la base de datos
func (db *DB) GetAgentByName(name string) (*entities.Agent, error) {
	query := `SELECT id, name, symbol, headquarters, credits, token FROM agents WHERE name = ?`
	row := db.Conn.QueryRow(query, name)
	var agent entities.Agent

	err := row.Scan(&agent.ID, &agent.Name, &agent.Symbol, &agent.Headquarters, &agent.Credits, &agent.Token)
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
	query := `SELECT name, token FROM agents`
	rows, err := db.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var agents entities.Agents
	var agent entities.Agent
	for rows.Next() {
		err := rows.Scan(&agent.Name, &agent.Token)
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
