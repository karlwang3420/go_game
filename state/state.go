package state

import (
	"go_game/item"
	"go_game/location"
)

type State struct {
	Loc        location.Location
	Health     int
	Equipments []item.Item
}
