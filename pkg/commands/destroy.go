package commands

import (
	"fmt"
	"github.com/stevedomin/frenzy/pkg"
	"github.com/stevedomin/frenzy/pkg/environment"
	"github.com/DimShadoWWW/terminal/color"
	"github.com/stevedomin/cli"
	"sync"
)

func Destroy(env *environment.Environment) *cli.Command {
	destroyCmd := cli.NewCommand("destroy")
	destroyCmd.HandlerFunc = func(args []string) {
		var wg sync.WaitGroup
		for _, node := range env.Nodes {
			wg.Add(1)
			go func(node *pkg.Node) {
				defer wg.Done()
				if node.Status != "not running" {
					err := node.Provider.Destroy(node.Hostname, node.ID)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					color.Println("[@r" + node.Hostname + color.ResetCode + "] already destroyed\n")
				}

				err := env.DestroyState()
				if err != nil {
					fmt.Println(err)
				}
			}(node)
		}
		wg.Wait()
	}

	return destroyCmd
}
