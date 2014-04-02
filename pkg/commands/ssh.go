package commands

import (
	"fmt"
	"github.com/stevedomin/frenzy/pkg"
	"github.com/stevedomin/frenzy/pkg/environment"
	"github.com/DimShadoWWW/terminal/color"
	"github.com/stevedomin/cli"
	"os"
	"os/exec"
)

func SSH(env *environment.Environment) *cli.Command {
	sshCmd := cli.NewCommand("ssh")
	sshCmd.HandlerFunc = func(args []string) {
		if len(args) == 0 {
			color.Println("@rYou need to specify the node you want to SSH into")
			fmt.Println("Example:")
			fmt.Println("	$ frenzy ssh node01")
			fmt.Println("")
			return
		}

		hostname := args[0]
		var node *pkg.Node

		for _, n := range env.Nodes {
			if n.Hostname == hostname {
				node = n
			}
		}

		if node.Status != "running" {
			color.Println("@rError: " + color.ResetCode + "[@r" + node.Hostname + color.ResetCode + "] Can't SSH. The node is not running.")
			return
		}

		cmd := exec.Command(
			"ssh",
			"-o", "StrictHostKeyChecking=no",
			"-o", "UserKnownHostsFile=/dev/null",
			"-i", "keys/frenzy_insecure_key",
			fmt.Sprintf("frenzy@%s", node.Host),
			"-p", node.Port,
		)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}
	return sshCmd
}
