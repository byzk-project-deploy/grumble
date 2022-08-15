package cmd

import (
	"fmt"
	"github.com/byzk-project-deploy/grumble"
)

func init() {
	App.AddCommand(&grumble.Command{
		Name: "confirm",
		Help: "shellTools测试",
		Run: func(c *grumble.Context) error {
			fmt.Println(c.ShellTools.Confirm("test"))
			return nil
		},
	})
}
