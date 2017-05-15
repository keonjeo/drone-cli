package user

import (
	"fmt"

	"github.com/drone/drone-go/drone"
	"github.com/urfave/cli"

	"github.com/drone/drone-cli/drone/internal"
)

var userAddCmd = cli.Command{
	Name:   "add",
	Usage:  "adds a user",
	Action: userAdd,
}

func userAdd(c *cli.Context) error {
	login := c.Args().First()

	client, err := internal.NewClient(c)
	if err != nil {
		return err
	}

	user, err := client.UserPost(&drone.User{Login: login})
	if err != nil {
		return err
	}
	fmt.Printf("Successfully added user %s\n", user.Login)
	return nil
}
