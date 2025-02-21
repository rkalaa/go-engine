mygame/
├── cmd/                # Entry points (command-line apps or main packages)
│   └── mygame/         # Main app binary name (e.g., "mygame")
│       └── main.go     # Main function and game initialization
├── internal/           # Private code (e.g., game logic, structs)
│   ├── game/           # Game-specific logic
│   │   └── game.go     # Game struct and Ebiten methods
│   └── objects/        # Object definitions
│       └── object.go   # Object struct and related methods
├── pkg/                # Reusable, public code (optional, if you share libs)
│   └── <future-lib>/   # (Empty for now)
├── go.mod              # Module definition
├── go.sum              # Dependency checksums
└── README.md           # Project docs