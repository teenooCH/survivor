package engine

import (
	"survivor/internal/entities/collision"
	"survivor/internal/entities/graph"
	"survivor/internal/entities/input"
	"survivor/internal/entities/resource"
	"survivor/internal/entities/world"
)

type Engine struct {
	world     *world.World
	input     *input.Manager
	resource  *resource.Manager
	collision *collision.Manager
}

func New(surface graph.Image, rm *resource.Manager) *Engine {
	return &Engine{
		world:     world.New(surface),
		input:     input.NewManager(),
		resource:  rm,
		collision: collision.NewManager(),
	}
}
func (e *Engine) World() *world.World                  { return e.world }
func (e *Engine) InputManager() *input.Manager         { return e.input }
func (e *Engine) ResourceManager() *resource.Manager   { return e.resource }
func (e *Engine) CollisionManager() *collision.Manager { return e.collision }

func (e *Engine) Update() {
	e.world.Update()
	// Add more update logic for input, collision, etc. if needed
}

func (e *Engine) Draw(target graph.Image) {
	e.world.Draw(target)
	// Add more draw logic for input, collision, etc. if needed
}
