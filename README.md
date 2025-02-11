# Pokedex

This is a commandline Pokedex, written in Go.

## Setup & Usage

This project uses the standard library and has **no external dependencies**.

### Usage

To build the application:
```bash
go build
```

To run the built application and have the time of your life:
```bash
./pokedex
```

or, you can cheat and do both at the same time with:
```bash
go run .
```

### Installation

If you'd like a **p**ermanent **P**okedex on your system (idk, maybe you *really* like Pokemon?), run the install command inside the source directory
```bash
go install
```

### Available Pokedex Commands
- `help` - Display the help screen and command list
- `exit` - Quit
- `map` - List the next 20 map locations
- `mapb` - List the previous 20 map locations
- `explore <location>` - Display details about the selected location
- `catch <pokemon>` - Attempt to catch a selected Pokemon
- `inspect <pokemon>` - Display the info of the selected (caught) Pokemon
- `pokedex` - List any caught Pokemon

## Project Structure

- `main.go` - entrypoint
- `repl.go` - controls the REPL loop for the pokedex
  - `repl_test.go` - contains tests for these functions
- Commands are listed in their relevant files:
  - Help/Exit in `commands.go`
  - Map commands in `commands_map.go` (map,mapb,explore)
  - Pokemon related commands in `commands_pokemon.go` (catch,inspect,pokedex)
- There are two internal packages:
  - Pokeapi (`internal/pokeapi`) calls [pokeapi.co](https://pokeapi.co/) to fetch data
    - `types.go` - types for the pokeapi calls & Pokemon data
    - `constants.go` - constants for the api
    - `client.go` - api client
    - `locations.go` - functions to make requests to the `location-area` endpoint
    - `pokemon.go` - functions to make requests to the `pokemon` endpoint
  - Pokecache (`internal/pokecache`) caches api requests for performance
    - `types.go` - types for the cache structure
    - `cache_functions.go` - handle cache creation & functionality
    - `cache_test.go` - tests for cache functions

---

This project was built as part of the Build a Pokedex course on [Boot.Dev](https://www.boot.dev/courses/build-pokedex-cli-golang)
