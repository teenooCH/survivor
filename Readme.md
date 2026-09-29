# Survivor - a 2D Game written in Go

The book [Beginning Game Programming with Go](https://leanpub.com/gameprogramminggolang) is a nice introduction into 2D Game Programming using Go. It explains the concepts of a Scene Graph, Input and Resource Management, dealing with different types of enemies, weapons, projectiles and more sophisticated concepts in a game.
There is a [GitHub repository](https://github.com/LuigiVanacore/beginning-game-programming-go-ebitengine/tree/main) for the code shown in the book.

I struggled with following the code and decided to write a new version myself. I follow roughly the chapters of the book.

## Chapters

Completed chapters are marked with Git tags. There are branches for the chapters, so it is possible to compare them to the original gopher-survivor code.

Last finished chapter: **06**

## Prerequisites

|Component |Version |
|:----|:-----|
|Go   |1.27  |
|Ebitengine|2.10| 

## Building and running

To run the game with the ebitengine:
```
go run ./cmd/survivor-ebiten
```

to build it
```
go build -o survivor ./cmd/survivor-ebiten
```

## Structure

The project follows [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html). Dependencies point inward toward the
domain, while framework-specific code stays in the infrastructure layer.

```text
.
├── cmd/
│   └── survivor-ebiten/       Application entry point and composition root
├── internal/
│   ├── domain/                 Core game concepts and rules
│   │   ├── camera/             Camera positioning and viewport behavior
│   │   ├── collision/          Collision shapes, masks, and management
│   │   ├── enemy/              Enemy entities
│   │   ├── input/              Input actions and controls
│   │   ├── node/               Scene graph nodes
│   │   ├── player/             Player entity and behavior
│   │   ├── sprite/             Sprite state and rendering data
│   │   ├── tile/               Tile maps, tilesets, and map generation
│   │   ├── ui/                 Game UI widgets and text
│   │   ├── vector/              Two-dimensional vector operations
│   │   └── world/              World state and callbacks
│   │
│   ├── application/            Use cases and game orchestration
│   │   ├── engine/              Main game engine coordination
│   │   ├── game/                Game creation and gameplay setup
│   │   ├── resource/            Loading and managing game resources
│   │   └── settings/            Application settings
│   │
│   ├── infrastructure/         Framework and external-library adapters
│   │   ├── assets/              Embedded maps, sprites, and manifests
│   │   └── ebiten/              Ebiten implementations of rendering and input
│   │
│   ├── ports/                  Interfaces between application and infrastructure
│   └── pkg/                    Reusable helpers independent of game rules
│       ├── stack/              Stack data structure
│       └── timer/              Timer utilities
│
└── go.mod                      Go module and dependency definitions
```
The application is assembled in cmd/survivor-ebiten. It connects application
services to infrastructure implementations through the interfaces in internal/ports.
The domain package contains the core game model and should remain independent of
Ebiten and other external frameworks.