package settings

// Texture keys for various game assets.
const (
	PlayerTexture = "player"
)

// Names for various game entities.
const (
	PlayerName = "player"
)

// Settings for the player character.
const (
	PlayerInitialHP    = 100.0
	PlayerInitialXP    = 0
	PlayerInitialLevel = 1
	PlayerSpeed        = 3.0 // pixels moved per Update call
)

// Resource keys for the floor tilemap. FloorTileset is the name prefix
// under which each tile of the spritesheet is stored in the resource
// Manager (e.g. "floor_0", "floor_1", ...); FloorMapPattern is the name
// under which the parsed floor.map tile grid is stored.
const (
	FloorTileset    = "floor"
	FloorMapPattern = "floor_map"
)

// Settings for the floor tilemap. The spritesheet stacks tiles vertically
// with 1px spacing between them (see infrastructure/assets.Spritesheet).
const (
	TileWidth         = 16
	TileHeight        = 16
	TileSpacing       = 1
	FloorTileCount    = 3
	TileMapLoadMargin = 1 // extra chunks kept loaded around the viewport
)
