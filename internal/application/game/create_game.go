package game

import (
	"fmt"

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
) (*Game, error) {
	camera := camera.New(screenWidth, screenHeight)
	world := world.New(camera)
	engine := engine.New(world, textures, inputProvider)

	RegisterDefaultBindings(engine)

	tm, err := CreateTileMap(settings.FloorName, engine.ResourceManager(), camera, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to create tilemap: %v", err)
	}

	engine.SetTileMap(tm)
	engine.World().AddNode(tm, 0)

	p, err := CreatePlayer(settings.PlayerName, engine, 1)
	if err != nil {
		return nil, fmt.Errorf("failed to create player: %v", err)
	}

	p.SetPosition(float64(screenWidth)/2, float64(screenHeight)/2)
	engine.World().AddNode(p, 1)
	camera.SetFollow(p)

	e, err := CreateEnemy(settings.EnemyName, engine, 1)
	if err != nil {
		return nil, fmt.Errorf("failed to create enemy: %v", err)
	}

	e.SetPosition(float64(screenWidth)/2+250, float64(screenHeight)/2)
	e.SetSpeed(settings.EnemySpeed)
	e.SetTarget(p)
	engine.World().AddNode(e, 1)

	gameOverWidget := CreateGameOverWidget(float64(screenWidth)/2, float64(screenHeight)/2)

	g := NewGame(engine, p, e, gameOverWidget)
	g.HUD().AddWidget(gameOverWidget)
	WirePlayerCallbacks(p, g)

	return g, nil
}
