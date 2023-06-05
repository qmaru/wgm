//go:build ui

package main

import (
	"embed"

	"wgm/ui/backend"
)

//go:embed all:ui/frontend/dist
var assets embed.FS

func main() {
	backend.Run(assets)
}
