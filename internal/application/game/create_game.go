package game

import (
	"log"

	"survivor/internal/application/engine"
	"survivor/internal/application/resource"
	"survivor/internal/application/settings"
	"survivor/internal/domain/camera"
	"survivor/internal/domain/world"
	"survivor/internal/ports"
)

// Use case for creating a new game instance.
// This function sets up the game world, including the player,
// enemy, camera, HUD, and game over widget.
// Here is the place where the dependencies are wired together.

func CreateGame(textures *resource.Manager,
	inputProvider ports.InputProvider,
	screenWidth, screenHeight uint,
) *Game {
	camera := camera.New(screenWidth, screenHeight)
	world := world.New(camera)
	engine := engine.New(world, textures, inputProvider)

	RegisterDefaultBindings(engine)

	if _, err := CreateTileMap(settings.FloorName, engine, camera, 0); err != nil {
		log.Fatalf("failed to create tilemap: %v", err)
	}

	p := CreatePlayer(settings.PlayerName, engine, 1)
	p.SetPosition(float64(screenWidth)/2, float64(screenHeight)/2)
	engine.World().AddNode(p, 1)
	camera.SetFollow(p)

	e := CreateEnemy(settings.EnemyName, engine, 1)
	e.SetPosition(float64(screenWidth)/2+250, float64(screenHeight)/2)
	e.SetSpeed(settings.EnemySpeed)
	e.SetTarget(p)
	engine.World().AddNode(e, 1)

	gameOverWidget := CreateGameOverWidget(float64(screenWidth)/2, float64(screenHeight)/2)
	engine.HUD().AddWidget(gameOverWidget)

	g := NewGame(engine, p, e, gameOverWidget)
	WirePlayerCallbacks(p, g)

	return g
}
