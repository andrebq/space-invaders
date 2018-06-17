package game

import (
	"github.com/andrebq/space-invaders/ces"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	playerShip      = findResource("player_ship.png")
	playerAnimation = findResource("player_ship.json")

	enemyShip      = findResource("enemy1.png")
	enemyAnimation = findResource("enemy1.json")

	gunFrames    = findResource("gun.png")
	gunAnimation = findResource("gun.json")

	explosionFrames    = findResource("explosion.png")
	explosionAnimation = findResource("explosion.json")

	youWinFrames    = findResource("you_win.png")
	youWinAnimation = findResource("you_win.json")
)

// CreatePlayerGun creates a new gun element fired by the player
func CreatePlayerGun(w *ces.World, p *Player) (*Gun, error) {
	gun, err := NewGun(gunFrames, gunAnimation)
	if err != nil {
		return nil, err
	}

	gun.MoveTo(sdl.Point{
		X: p.Pos.X,
		Y: p.Pos.Y + 40,
	})
	w.AddEntity(gun)
	return gun, nil
}

// CreatePlayer returns a new player
func CreatePlayer(w *ces.World) (*Player, error) {
	player, err := NewPlayer(playerShip, playerAnimation)
	if err != nil {
		return nil, err
	}
	w.AddEntity(player)

	centeredRect := GetWorld(w).GetCentered(player.RectAt(player.Pos))
	player.MoveTo(sdl.Point{
		X: centeredRect.X,
		Y: 50,
	})

	return player, nil
}

// CreateEnemy returns a new enemy
func CreateEnemy(w *ces.World, pos sdl.Point) (*Enemy, error) {
	enemy, err := NewEnemy(enemyShip, enemyAnimation)
	if err != nil {
		return nil, err
	}
	w.AddEntity(enemy)
	enemy.MoveTo(pos)

	return enemy, nil
}

// CreateEnemyLane creates a full lane of enemies
func CreateEnemyLane(w *ces.World, lane int32) error {
	templateEnemy, err := NewEnemy(enemyShip, enemyAnimation)
	if err != nil {
		return err
	}

	centeredRect := GetWorld(w).GetCentered(templateEnemy.RectAt(sdl.Point{}))

	for i := 0; i < 4; i++ {
		e, err := CreateEnemy(w, sdl.Point{
			Y: centeredRect.Y + (lane * 40) - 20,
			X: centeredRect.X - (int32(i) * 40),
		})
		if err != nil {
			return err
		}
		e.SetDirection(1)
	}

	for i := 0; i < 4; i++ {
		e, err := CreateEnemy(w, sdl.Point{
			Y: centeredRect.Y + (lane * 40) - 20,
			X: centeredRect.X + (int32(i) * 40),
		})
		if err != nil {
			return err
		}
		e.SetDirection(1)
	}
	return nil
}

// CreateExplosion adds a new explosion at the given point
func CreateExplosion(w *ces.World, pos sdl.Point) (*Explosion, error) {
	explosion, err := NewExplosion(explosionFrames, explosionAnimation)
	if err != nil {
		return nil, err
	}
	w.AddEntity(explosion)
	explosion.MoveTo(pos)

	return explosion, nil
}

// CreateYouWin adds the YouWin banner
func CreateYouWin(w *ces.World) (*YouWin, error) {
	youwin, err := NewYouWin(youWinFrames, youWinAnimation)
	if err != nil {
		return nil, err
	}
	w.AddEntity(youwin)
	centeredRect := GetWorld(w).GetCentered(youwin.RectAt(sdl.Point{}))
	youwin.MoveTo(sdl.Point{
		X: centeredRect.X,
		Y: centeredRect.Y,
	})

	return youwin, nil
}
