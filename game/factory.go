package game

import (
	"github.com/andrebq/space-invaders/ces"
	"github.com/andrebq/space-invaders/ces/sfx"
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

	goodGameFrames    = findResource("good_game.png")
	goodGameAnimation = findResource("good_game.json")

	bombSfx      = findResource("bomb_explosion.wav")
	playerGunSfx = findResource("explosion-1.wav")

	youWinSfx   = findResource("special01.wav")
	goodGameSfx = findResource("special02.wav")
)

// CreatePlayerGun creates a new gun element fired by the player
func CreatePlayerGun(w *ces.World, p *Player) (*Gun, error) {
	gun, err := NewGun(gunFrames, gunAnimation, p.Key())
	if err != nil {
		return nil, err
	}

	gun.MoveTo(sdl.Point{
		X: p.Pos.X,
		Y: p.Pos.Y + 40,
	})
	w.AddEntity(gun)

	CreatePlayerGunSfx(w)

	return gun, nil
}

// CreateEnemyGun creates a new gun element fired by the enemy
func CreateEnemyGun(w *ces.World, p *Enemy) (*Gun, error) {
	gun, err := NewGun(gunFrames, gunAnimation, p.Key())
	if err != nil {
		return nil, err
	}

	gun.MoveTo(sdl.Point{
		X: p.Pos.X,
		Y: p.Pos.Y - 5,
	})
	gun.SetDirection(-1)
	gun.SetSpeed(200)
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
			Y: centeredRect.Y + 50 + (lane * 40),
			X: centeredRect.X - (int32(i) * 40),
		})
		if err != nil {
			return err
		}
		e.SetDirection(1)
	}

	for i := 0; i < 4; i++ {
		e, err := CreateEnemy(w, sdl.Point{
			Y: centeredRect.Y + 50 + (lane * 40),
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

	CreateExplosionSfx(w)

	return explosion, nil
}

// CreateYouWin adds the YouWin banner
func CreateYouWin(w *ces.World) (*EndAnimation, error) {
	youwin, err := NewGameEnd(youWinFrames, youWinAnimation)
	if err != nil {
		return nil, err
	}
	w.AddEntity(youwin)
	centeredRect := GetWorld(w).GetCentered(youwin.RectAt(sdl.Point{}))
	youwin.MoveTo(sdl.Point{
		X: centeredRect.X,
		Y: centeredRect.Y,
	})
	CreateYouWinSfx(w)

	return youwin, nil
}

// CreateGoodGame adds the YouWin banner
func CreateGoodGame(w *ces.World) (*EndAnimation, error) {
	youwin, err := NewGameEnd(goodGameFrames, goodGameAnimation)
	if err != nil {
		return nil, err
	}
	w.AddEntity(youwin)
	centeredRect := GetWorld(w).GetCentered(youwin.RectAt(sdl.Point{}))
	youwin.MoveTo(sdl.Point{
		X: centeredRect.X,
		Y: centeredRect.Y,
	})
	CreateYouLoseSfx(w)

	return youwin, nil
}

// CreateExplosionSfx creates a new sound effect for explosions
func CreateExplosionSfx(w *ces.World) (*sfx.Effect, error) {
	effect, err := sfx.NewEffect(bombSfx)
	if err != nil {
		return nil, err
	}
	w.AddEntity(effect)
	return effect, nil
}

// CreatePlayerGunSfx creates a new sound effect for player gun fire
func CreatePlayerGunSfx(w *ces.World) (*sfx.Effect, error) {
	effect, err := sfx.NewEffect(playerGunSfx)
	if err != nil {
		return nil, err
	}
	w.AddEntity(effect)
	return effect, nil
}

func CreateYouWinSfx(w *ces.World) (*sfx.Effect, error) {
	effect, err := sfx.NewEffect(youWinSfx)
	if err != nil {
		return nil, err
	}
	w.AddEntity(effect)
	return effect, nil
}

func CreateYouLoseSfx(w *ces.World) (*sfx.Effect, error) {
	effect, err := sfx.NewEffect(goodGameSfx)
	if err != nil {
		return nil, err
	}
	w.AddEntity(effect)
	return effect, nil
}
