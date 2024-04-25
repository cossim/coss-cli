package cmd

import (
	"fmt"
	"github.com/cossim/coss-cli/pkg/consul"
	"github.com/urfave/cli/v2"
	"strings"
)

func initConfig(cCtx *cli.Context) error {
	if cCtx.String("path") == "" {
		return fmt.Errorf("path is empty")
	}
	paths := strings.Split(cCtx.String("path"), ",")
	if len(paths) == 0 {
		return fmt.Errorf("no paths provided")
	}
	nameSpace := cCtx.String("namespace")
	host := cCtx.String("host")
	token := cCtx.String("token")
	token = token + "/"
	client := consul.NewConsulClient(host, nameSpace, "", token)

	for _, path := range paths {
		client.SetPath(path)
		err := client.PutConfig()
		if err != nil {
			return err
		}
	}

	return nil
}
