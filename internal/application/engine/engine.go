package engine

import (
	"survivor/internal/application/resource"
	"survivor/internal/domain/collision"
	"survivor/internal/domain/graph"
	"survivor/internal/domain/input"
	"survivor/internal/domain/world"
	"survivor/internal/ports"
)

type Engine struct {
	world     *world.World
	input     *input.Manager
	resource  *resource.Manager
	collision *collision.Manager
}

// New wires the engine together. inputProvider is the outgoing interface
// implementation (e.g. infrastructure/ebiten.InputProvider) supplied by the
// composition root, keeping the engine free of any hardware dependency.
func New(world *world.World, rm *resource.Manager, inputProvider ports.InputProvider) *Engine {
	return &Engine{
		world:     world,
		input:     input.NewManager(inputProvider),
		resource:  rm,
		collision: collision.NewManager(),
	}
}
func (e *Engine) World() *world.World                  { return e.world }
func (e *Engine) InputManager() *input.Manager         { return e.input }
func (e *Engine) ResourceManager() *resource.Manager   { return e.resource }
func (e *Engine) CollisionManager() *collision.Manager { return e.collision }

func (e *Engine) UpdateWorld() {
	e.world.Update()
}

func (e *Engine) UpdateCamera() {
	e.world.Camera().Update()
}

func (e *Engine) ProcessCollisions() {
	e.collision.ProcessCollisions()
}

func (e *Engine) Draw(target graph.Image) {
	e.world.Draw(target)
}
