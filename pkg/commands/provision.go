package commands

import (
	// "fmt"
	"github.com/stevedomin/frenzy/pkg"
	"github.com/stevedomin/frenzy/pkg/environment"
	"github.com/DimShadoWWW/terminal/color"
	"github.com/stevedomin/cli"
	"sync"
)

func Provision(env *environment.Environment) *cli.Command {
	provisionCmd := cli.NewCommand("provision")
	provisionCmd.HandlerFunc = func(args []string) {
		var wg sync.WaitGroup
		for _, node := range env.Nodes {
			wg.Add(1)
			go func(node *pkg.Node) {
				defer wg.Done()
				if node.Status != "running" {
					color.Println("[@r" + node.Hostname + color.ResetCode + "] skip provisioning since node is not running\n")
					return
				}
				for _, provisioner := range node.Provisioners {
					provisioner.Run(node)
				}
			}(node)
		}
		wg.Wait()
	}
	return provisionCmd
}
