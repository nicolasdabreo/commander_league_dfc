package services

import (
	"dfc/db"
	"fmt"
)

type Player struct {
	ID   uint `json:"id"`
	Name string
	Deck string

	TotalPoints int
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

func (ps *PlayerServices) ListPlayersForSeason() ([]Player, error) {
	query := fmt.Sprintf("SELECT id, name, deck FROM players ORDER BY created_at DESC")

	rows, err := ps.PlayerStore.Db.Query(query)
	if err != nil {
		return []Player{}, err
	}
	// We close the resource
	defer rows.Close()

	players := []Player{}
	for rows.Next() {
		rows.Scan(&ps.Player.ID)

		players = append(players, ps.Player)
	}

	return players, nil
}
