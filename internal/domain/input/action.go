package input

// Action is a semantic, game-specific input intent (e.g. "move up").
// Gameplay code should depend on Action, never on a concrete Key or hardware button,
// so that rebinding controls never touches domain or application code.
type Action int

const (
	ActionMoveUp Action = iota
	ActionMoveDown
	ActionMoveLeft
	ActionMoveRight

	ActionConfirm
	ActionCancel
)
