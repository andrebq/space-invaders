package game

import (
	"github.com/andrebq/space-invaders/ces"
)

// Stage1 creates the entities for the first stage of the game
func (w *World) Stage1(cesw *ces.World) error {
	_, err := CreatePlayer(cesw)
	if err != nil {
		return err
	}
	for i := 1; i <= 3; i++ {
		err = CreateEnemyLane(cesw, int32(i))
		if err != nil {
			return err
		}
	}

	return nil
}
