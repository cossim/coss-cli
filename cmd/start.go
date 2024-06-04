package cmd

import (
	"github.com/urfave/cli/v2"
	"io"
	"os"
	"os/exec"
)

func start(cCtx *cli.Context) error {
	err := gen(cCtx)
	if err != nil {
		return err
	}
	//启动docker-compose
	cmd := exec.Command("docker-compose", "-f", "docker-compose.yaml", "up", "-d")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	go io.Copy(os.Stdout, stdout)
	go io.Copy(os.Stderr, stderr)

	err = cmd.Wait()
	if err != nil {
		return err
	}

	if !cCtx.Bool("direct") {
		err := initConfig(cCtx)
		if err != nil {
			return err
		}
	}

	return nil
}
