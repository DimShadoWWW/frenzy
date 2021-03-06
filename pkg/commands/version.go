package commands

import (
	"fmt"
	"github.com/stevedomin/frenzy/pkg"
	"github.com/stevedomin/cli"
)

func Version() *cli.Command {
	versionCmd := cli.NewCommand("version")
	versionCmd.HandlerFunc = func(args []string) {
		fmt.Printf("frenzy v%s\n", pkg.Version)
	}
	return versionCmd
}
