package cmd

import (
	"fmt"
	"github.com/cossim/coss-cli/config"
	"github.com/cossim/coss-cli/pkg/consul"
	"github.com/cossim/coss-cli/pkg/pgp"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
	"path/filepath"
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
		},
		{
			Name:  "gen",
			Usage: "gen coss config",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "direct",
					Value: true,
					Usage: "true or false",
				},
				&cli.StringFlag{
					Name:  "path",
					Value: "./",
					Usage: "config path",
				},
			},
			Action: func(cCtx *cli.Context) error {
				direct := cCtx.Bool("direct")
				outputDir := cCtx.String("path")
				cacheDir := "./config"
				if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
					err := os.Mkdir(cacheDir, 0755) // 创建文件夹并设置权限
					if err != nil {
						return err
					}
				}

				cacheDir = "./config/interface"
				if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
					err := os.Mkdir(cacheDir, 0755) // 创建文件夹并设置权限
					if err != nil {
						return err
					}
				}

				cacheDir = "./config/service"
				if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
					err := os.Mkdir(cacheDir, 0755) // 创建文件夹并设置权限
					if err != nil {
						return err
					}
				}

				cacheDir = "./config/common"
				if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
					err := os.Mkdir(cacheDir, 0755) // 创建文件夹并设置权限
					if err != nil {
						return err
					}
				}

				if direct {
					// 生成接口配置文件
					for _, name := range config.InterfaceList {
						configStr := config.GenInterfaceConfig(name, config.InetrfaceConfig[name], direct)
						filePath := filepath.Join(outputDir+"/config/interface/", fmt.Sprintf("%s.yaml", config.InetrfaceName[name]))
						err := ioutil.WriteFile(filePath, []byte(configStr), 0644)
						if err != nil {
							fmt.Printf("写入文件 %s 失败：%v\n", filePath, err)
						} else {
							fmt.Printf("生成 %s 成功\n", filePath)
						}
					}

					// 生成服务配置文件
					for _, name := range config.ServiceList {
						configStr := config.GenServiceConfig(name, config.ServiceConfig[name])
						filePath := filepath.Join(outputDir+"/config/service/", fmt.Sprintf("%s.yaml", config.ServiceName[name]))
						err := ioutil.WriteFile(filePath, []byte(configStr), 0644)
						if err != nil {
							fmt.Printf("写入文件 %s 失败：%v\n", filePath, err)
						} else {
							fmt.Printf("生成 %s 成功\n", filePath)
						}
					}
				} else {
					fmt.Println("生成 consul 配置文件")
					// 生成接口配置文件
					for _, name := range config.InterfaceList {
						configStr := config.GenConsulInterfaceConfig(name, config.InetrfaceConfig[name], direct)
						filePath := filepath.Join(outputDir+"/config/interface/", fmt.Sprintf("%s.yaml", config.InetrfaceName[name]))
						err := ioutil.WriteFile(filePath, []byte(configStr), 0644)
						if err != nil {
							fmt.Printf("写入文件 %s 失败：%v\n", filePath, err)
						} else {
							fmt.Printf("生成 %s 成功\n", filePath)
						}
					}

					// 生成服务配置文件
					for _, name := range config.ServiceList {
						configStr := config.GenConsulServiceConfig(name, config.ServiceConfig[name])
						filePath := filepath.Join(outputDir+"/config/service/", fmt.Sprintf("%s.yaml", config.ServiceName[name]))
						err := ioutil.WriteFile(filePath, []byte(configStr), 0644)
						if err != nil {
							fmt.Printf("写入文件 %s 失败：%v\n", filePath, err)
						} else {
							fmt.Printf("生成 %s 成功\n", filePath)
						}
					}

					//生成公共配置
					for _, name := range config.ConsulCommonList {
						configStr := config.GenConsulCommonConfig(name)
						filePath := filepath.Join(outputDir+"/config/common/", fmt.Sprintf("%s.yaml", name))
						err := ioutil.WriteFile(filePath, []byte(configStr), 0644)
						if err != nil {
							fmt.Printf("写入文件 %s 失败：%v\n", filePath, err)
						} else {
							fmt.Printf("生成 %s 成功\n", filePath)
						}
					}
				}

				// 生成公共配置文件
				for _, name := range config.CommonClist {
					configStr := config.GenCommonConfig(name)
					if name == "consul" {
						filePath := filepath.Join(outputDir+"/config/common/", fmt.Sprintf("%s.json", name))
						err := ioutil.WriteFile(filePath, []byte(configStr), 0644)
						if err != nil {
							fmt.Printf("写入文件 %s 失败：%v\n", filePath, err)
						} else {
							fmt.Printf("生成 %s 成功\n", filePath)
						}
					} else {
						filePath := filepath.Join(outputDir+"/config/common/", fmt.Sprintf("%s.yaml", name))
						err := ioutil.WriteFile(filePath, []byte(configStr), 0644)
						if err != nil {
							fmt.Printf("写入文件 %s 失败：%v\n", filePath, err)
						} else {
							fmt.Printf("生成 %s 成功\n", filePath)
						}
					}

				}

				//生成docker-compose
				configStr := ""
				if direct {
					configStr = config.GenDockerCompose(false)
				} else {
					configStr = config.GenDockerCompose(true)
				}
				filePath := filepath.Join(outputDir, "docker-compose.yaml")
				err := ioutil.WriteFile(filePath, []byte(configStr), 0644)
				if err != nil {
					fmt.Printf("写入文件 %s 失败：%v\n", filePath, err)
					return err
				} else {
					fmt.Printf("生成 %s 成功\n", filePath)
				}

				//生成pgp公私钥
				err = pgp.GenerateKeyPair()
				if err != nil {
					return err
				}

				return nil
			},
		},
	},
}
