package ebiten

import (
	"survivor/internal/ports"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// keyMapping translates the domain's hardware-agnostic ports.Key into a
// concrete ebiten.Key. It is the single place where the domain's Key type
// meets the ebiten library.
var keyMapping = map[ports.Key]ebiten.Key{
	ports.KeyW: ebiten.KeyW,
	ports.KeyA: ebiten.KeyA,
	ports.KeyS: ebiten.KeyS,
	ports.KeyD: ebiten.KeyD,

	ports.KeyUp:    ebiten.KeyArrowUp,
	ports.KeyDown:  ebiten.KeyArrowDown,
	ports.KeyLeft:  ebiten.KeyArrowLeft,
	ports.KeyRight: ebiten.KeyArrowRight,

	ports.KeySpace:  ebiten.KeySpace,
	ports.KeyEscape: ebiten.KeyEscape,
	ports.KeyEnter:  ebiten.KeyEnter,
}

// InputProvider implements ports.InputProvider using the ebiten library.
// It is the outgoing interface's concrete adapter and is injected into
// domain/input.NewManager from the composition root (e.g. cmd/*/main.go).
type InputProvider struct{}

func NewInputProvider() *InputProvider {
	return &InputProvider{}
}

func (p *InputProvider) IsKeyPressed(key ports.Key) bool {
	ebKey, ok := keyMapping[key]
	if !ok {
		return false
	}

	return ebiten.IsKeyPressed(ebKey)
}

func (p *InputProvider) IsKeyJustPressed(key ports.Key) bool {
	ebKey, ok := keyMapping[key]
	if !ok {
		return false
	}

	return inpututil.IsKeyJustPressed(ebKey)
}

func (p *InputProvider) IsKeyJustReleased(key ports.Key) bool {
	ebKey, ok := keyMapping[key]
	if !ok {
		return false
	}

	return inpututil.IsKeyJustReleased(ebKey)
}
