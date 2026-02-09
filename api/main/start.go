package main

import (
	"github.com/aldesgroup/goald"

	// --- generic features

	// to add translations handling; requires
	_ "github.com/aldesgroup/goald/_include/i18n"

	// --- specific features

	// the package that should contain, among other things, the config definition
	_ "github.com/aldesgroup/devotion-project-template/shared"

	// all the applicative packages must be imported here
	_ "github.com/aldesgroup/devotion-project-template/_include/reference"
)

func main() {
	// Starting a new server
	goald.NewServer().Start()
}
