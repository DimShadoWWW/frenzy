package commands

import (
	// "fmt"
	"github.com/stevedomin/frenzy/pkg/environment"
	"github.com/DimShadoWWW/terminal/color"
	"github.com/stevedomin/cli"
)

func Stop(env *environment.Environment) *cli.Command {
	stopCmd := cli.NewCommand("stop")
	stopCmd.HandlerFunc = func(args []string) {
		color.Println("@rnot yet implemented\n")
		// Docker commit container?
		// Docker stop
		// defer env.SaveState()
	}
	return stopCmd
}
