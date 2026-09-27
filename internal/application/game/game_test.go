package game_test

import (
	"testing"

	"survivor/internal/application/engine"
	gameapp "survivor/internal/application/game"
	"survivor/internal/domain/camera"
	"survivor/internal/domain/enemy"
	"survivor/internal/domain/input"
	"survivor/internal/domain/player"
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
