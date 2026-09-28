package game_test

import (
	"testing"

	"survivor/internal/application/engine"
	gameapp "survivor/internal/application/game"
	"survivor/internal/domain/camera"
	"survivor/internal/domain/collision"
	"survivor/internal/domain/enemy"
	"survivor/internal/domain/graph"
	"survivor/internal/domain/input"
	"survivor/internal/domain/player"
	"survivor/internal/domain/tile"
	"survivor/internal/domain/ui"
	"survivor/internal/domain/world"
	"survivor/internal/ports"
)

type inputProvider struct {
	pressed map[input.Key]bool
}

func (p inputProvider) IsKeyPressed(key input.Key) bool {
	return p.pressed[key]
}

func (inputProvider) IsKeyJustPressed(input.Key) bool                  { return false }
func (inputProvider) IsKeyJustReleased(input.Key) bool                 { return false }
func (inputProvider) IsMouseButtonPressed(input.MouseButton) bool      { return false }
func (inputProvider) IsMouseButtonJustPressed(input.MouseButton) bool  { return false }
func (inputProvider) IsMouseButtonJustReleased(input.MouseButton) bool { return false }
func (inputProvider) IsGamepadButtonPressed(input.GamepadID, input.GamepadButton) bool {
	return false
}

func (inputProvider) IsGamepadButtonJustPressed(input.GamepadID, input.GamepadButton) bool {
	return false
}

func (inputProvider) IsGamepadButtonJustReleased(input.GamepadID, input.GamepadButton) bool {
	return false
}
func (inputProvider) ConnectedGamepadIDs() []input.GamepadID { return nil }

var _ ports.InputProvider = inputProvider{}

type fakeImage struct{}

func (fakeImage) DrawImage(graph.Image, graph.DrawOpt) {}
func (fakeImage) Dimensions() (width, height float64)  { return 0, 0 }
func (fakeImage) SetDimensions(width, height float64)  {}

type recordingGenerator struct {
	generated []tile.ChunkCoord
}

func (g *recordingGenerator) ChunkSize() (width, height int) {
	return 1, 1
}

func (g *recordingGenerator) Generate(coord tile.ChunkCoord) *tile.Chunk {
	g.generated = append(g.generated, coord)
	return &tile.Chunk{Coord: coord, Width: 1, Height: 1, Tiles: []int{0}}
}

func TestGameUpdateMovesPlayerOnce(t *testing.T) {
	provider := inputProvider{
		pressed: map[input.Key]bool{input.KeyD: true},
	}
	inputManager := input.NewManager(provider)
	inputManager.BindKey(input.ActionMoveRight, input.KeyD)

	player := player.New("player", inputManager, 1)
	player.SetPosition(0, 0)

	enemy := enemy.New("enemy")
	enemy.SetPosition(100, 0)

	world := world.New(camera.New(640, 480))
	world.AddNode(player, 0)
	world.AddNode(enemy, 0)

	engine := engine.New(world, nil, provider)
	game := gameapp.NewGame(engine, player, enemy, ui.NewText("Game Over", 0, 0))

	game.Update()

	got := player.GetPosition().X()
	if got != 1 {
		t.Fatalf("player moved to x=%v, want x=1", got)
	}
}

func TestGameUpdateProcessesCollisionAfterMovement(t *testing.T) {
	provider := inputProvider{pressed: map[input.Key]bool{input.KeyD: true}}
	inputManager := input.NewManager(provider)
	inputManager.BindKey(input.ActionMoveRight, input.KeyD)

	player := player.New("player", inputManager, 1)
	player.SetPosition(0, 0)

	enemy := enemy.New("enemy")
	enemy.SetPosition(2.5, 0)

	world := world.New(camera.New(640, 480))
	world.AddNode(player, 0)
	world.AddNode(enemy, 0)

	engine := engine.New(world, nil, provider)
	playerCollider := collision.NewCollider(
		"player-collider",
		collision.NewMask(collision.LayerPlayer, collision.LayerEnemy),
		collision.NewCircle(1),
	)
	enemyCollider := collision.NewCollider(
		"enemy-collider",
		collision.NewMask(collision.LayerEnemy, collision.LayerPlayer),
		collision.NewCircle(1),
	)

	player.AddChild(playerCollider)
	enemy.AddChild(enemyCollider)
	engine.CollisionManager().AddCollider(playerCollider)
	engine.CollisionManager().AddCollider(enemyCollider)

	gameOverWidget := ui.NewText("Game Over", 0, 0)
	game := gameapp.NewGame(engine, player, enemy, gameOverWidget)

	playerCollider.SetCollisionHandler(func(*collision.Collider) {
		game.SetGameOver()
	})

	game.Update()

	if !gameOverWidget.IsVisible() {
		t.Fatal("collision was not processed after player movement")
	}
}

func TestGameUpdateUpdatesCameraBeforeStreaming(t *testing.T) {
	provider := inputProvider{pressed: map[input.Key]bool{input.KeyD: true}}
	inputManager := input.NewManager(provider)
	inputManager.BindKey(input.ActionMoveRight, input.KeyD)

	player := player.New("player", inputManager, 6)
	player.SetPosition(9, 0)

	enemy := enemy.New("enemy")

	camera := camera.New(10, 10)
	camera.SetFollow(player)
	world := world.New(camera)
	world.AddNode(player, 0)
	world.AddNode(enemy, 0)

	engine := engine.New(world, nil, provider)
	generator := &recordingGenerator{}
	tileMap := tile.NewMap(
		"floor", tile.NewTileSet([]graph.Image{fakeImage{}}), generator,
		camera, 10, 10, 0, 0,
	)
	engine.SetTileMap(tileMap)

	game := gameapp.NewGame(engine, player, enemy, ui.NewText("Game Over", 0, 0))
	game.Update()

	if len(generator.generated) != 4 {
		t.Fatalf("generated chunks = %d, want 4", len(generator.generated))
	}

	for _, coord := range generator.generated {
		if coord.X < 1 {
			t.Fatalf("streaming used stale camera position and generated chunk %+v", coord)
		}
	}
}

func TestGameDrawDoesNotUpdateGameState(t *testing.T) {
	provider := inputProvider{}
	player := player.New("player", input.NewManager(provider), 1)
	player.SetPosition(20, 30)

	enemy := enemy.New("enemy")

	camera := camera.New(100, 100)
	camera.SetPosition(4, 5)
	world := world.New(camera)
	world.AddNode(player, 0)
	world.AddNode(enemy, 0)

	engine := engine.New(world, nil, provider)
	game := gameapp.NewGame(engine, player, enemy, ui.NewText("Game Over", 0, 0))

	game.Draw(fakeImage{})

	if got := player.GetPosition(); got.X() != 20 || got.Y() != 30 {
		t.Fatalf("player position after Draw() = %v, want (20, 30)", got)
	}

	if got := camera.GetPosition(); got.X() != 4 || got.Y() != 5 {
		t.Fatalf("camera position after Draw() = %v, want (4, 5)", got)
	}
}

var (
	_ graph.Image         = fakeImage{}
	_ tile.ChunkGenerator = (*recordingGenerator)(nil)
)
