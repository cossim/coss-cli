package cmd

import (
	"fmt"
	"github.com/cossim/coss-cli/pkg/consul"
	"github.com/urfave/cli/v2"
)

var App = &cli.App{
	Name:  "coss-cli",
	Usage: "coss-cli is a command line tool for coss",
	Commands: []*cli.Command{
		{
			Name:  "config",
			Usage: "init consul config",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "path",
					Value: "",
					Usage: "config path",
				},
				&cli.StringFlag{
					Name:    "namespace",
					Aliases: []string{"n"},
					Value:   "default",
					Usage:   "config namespace",
				},
				&cli.StringFlag{
					Name:  "host",
					Value: "127.0.0.1:8500",
					Usage: "consul host",
				},
				&cli.StringFlag{
					Name:  "token",
					Value: "",
					Usage: "consul acl token",
				},
				&cli.BoolFlag{
					Name:  "ssl",
					Value: false,
					Usage: "consul enable ssl",
				},
			},
			Action: func(cCtx *cli.Context) error {
				if cCtx.String("path") == "" {
					return fmt.Errorf("path is empty")
				}
				path := cCtx.String("path")
				nameSpace := cCtx.String("namespace")
				host := cCtx.String("host")
				token := cCtx.String("token")
				ssl := cCtx.Bool("ssl")
				client := consul.NewConsulClient(host, nameSpace, path, token, ssl)
				err := client.PutConfig()
				if err != nil {
					return err
				}
				return nil
			},
			//Subcommands: []*cli.Command{
			//	{
			//		Name:  "get",
			//		Usage: "get config",
			//		Action: func(cCtx *cli.Context) error {
			//			ssl := cCtx.Bool("ssl")
			//			host := cCtx.String("host")
			//			fmt.Println("ssl: ", ssl)
			//			fmt.Println("host: ", host)
			//			client := consul.NewConsulClient(host, "", "", "", ssl)
			//			token, err := client.GetToken()
			//			if err != nil {
			//				return err
			//			}
			//			fmt.Println(token)
			//			return nil
			//		},
			//	},
			//},
		},
	},
}
