package service

import (
	"dfc/db"
	"dfc/types"
	"fmt"
)

type Player struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Deck string `json:"deck"`
}

type PlayerServices struct {
	Player      Player
	PlayerStore db.Store
}

func NewPlayerServices(t Player, tStore db.Store) *PlayerServices {
	return &PlayerServices{
		Player:      t,
		PlayerStore: tStore,
	}
}

func (ps *PlayerServices) ListPlayersForLeague() ([]Player, error) {
	query := "SELECT id, name, deck FROM players ORDER BY created_at DESC"

	rows, err := ps.PlayerStore.Db.Query(query)
	if err != nil {
		return []Player{}, err
	}
	// We close the resource
	defer rows.Close()

	players := []Player{}
	for rows.Next() {
		var player Player
		if err := rows.Scan(&player.ID, &player.Name, &player.Deck); err != nil {
			return nil, err
		}
		players = append(players, player)
	}

	return players, nil
}

func (ps *PlayerServices) CreatePlayer(params types.PlayerParams) (Player, error) {
	query := fmt.Sprintf("INSERT INTO players (name, deck) VALUES ('%s', '%s') RETURNING id", params.Name, params.Deck)

	var playerID uint
	err := ps.PlayerStore.Db.QueryRow(query).Scan(&playerID)
	if err != nil {
		return Player{}, err
	}

	createdPlayer := Player{
		ID:   playerID,
		Name: params.Name,
		Deck: params.Deck,
	}

	return createdPlayer, nil
}

func (ps *PlayerServices) GetPlayer(playerID string) (Player, error) {
	query := fmt.Sprintf("SELECT id, name, deck FROM players WHERE id = '%s' LIMIT 1", playerID)

	var player Player
	err := ps.PlayerStore.Db.QueryRow(query).Scan(&player)

	if err != nil {
		return Player{}, err
	}

	return player, nil
}
