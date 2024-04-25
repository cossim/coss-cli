package cmd

import (
	"fmt"
	"github.com/cossim/coss-cli/config"
	"github.com/cossim/coss-cli/pkg/apisix"
	"github.com/cossim/coss-cli/pkg/consul"
	"github.com/cossim/coss-cli/pkg/pgp"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const VERSION = "v1.0.0"

var App = &cli.App{
	Name:    "coss-cli",
	Usage:   "coss-cli is a command line tool for coss",
	Version: VERSION,
	Commands: []*cli.Command{
		{
			Name:  "config",
			Usage: "init consul config",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "path",
					Value:   "",
					Aliases: []string{"p"},
					Usage:   "config path (multiple paths separated by commas)",
				},
				&cli.StringFlag{
					Name:    "namespace",
					Aliases: []string{"n"},
					Value:   "default",
					Usage:   "config namespace",
				},
				&cli.StringFlag{
					Name:  "host",
					Value: "http://127.0.0.1:8500",
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
			},
		},
		{
			Name:  "gen",
			Usage: "gen coss config",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:    "direct",
					Value:   true,
					Usage:   "true or false: --direct=false or -d=false",
					Aliases: []string{"d"},
				},
				&cli.StringFlag{
					Name:  "path",
					Value: "./",
					Usage: "config path",
				},
				&cli.StringFlag{
					Name:  "env",
					Value: "dev",
					Usage: "dev or prod",
				},
				&cli.StringFlag{
					Name:  "domain",
					Value: "tuo.gezi.vip",
					Usage: "your domain name",
				},
				&cli.BoolFlag{
					Name:  "ssl",
					Value: false,
					Usage: "consul enable ssl",
				},
			},
			Action: func(cCtx *cli.Context) error {
				direct := cCtx.Bool("direct")
				outputDir := cCtx.String("path")
				env := cCtx.String("env")
				domain := cCtx.String("domain")
				enableSsl := cCtx.Bool("ssl")
				cacheDir := "./config"

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
					// 生成服务配置文件
					for _, name := range config.ServiceList {

						httpname := config.HttpName[name]
						grpcname := config.GrpcName[name]
						httpport := config.HttpPort[httpname]
						grpcport := config.GrpcPort[grpcname]

						configStr := config.GenServiceConfig(httpname, grpcname, httpport, grpcport, env, enableSsl, domain)
						filePath := filepath.Join(outputDir+"/config/service/", fmt.Sprintf("%s.yaml", name))
						err := ioutil.WriteFile(filePath, []byte(configStr), 0644)
						if err != nil {
							fmt.Printf("写入文件 %s 失败：%v\n", filePath, err)
						} else {
							fmt.Printf("生成 %s 成功\n", filePath)
						}
					}

				} else {
					fmt.Println("生成 consul 配置文件")

					// 生成服务配置文件
					for _, name := range config.ServiceList {

						httpname := config.HttpName[name]
						grpcname := config.GrpcName[name]
						httpport := config.HttpPort[httpname]
						grpcport := config.GrpcPort[grpcname]

						configStr := config.GenConsulServiceConfig(httpname, grpcname, httpport, grpcport, env, enableSsl, domain)
						filePath := filepath.Join(outputDir+"/config/service/", fmt.Sprintf("%s.yaml", name))
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
		{
			Name:  "route",
			Usage: "init gateway route",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:    "direct",
					Value:   true,
					Usage:   "true or false: --direct=false or -d=false",
					Aliases: []string{"d"},
				},
				&cli.StringFlag{
					Name:  "key",
					Value: "edd1c9f034335f136f87ad84b625c8f1",
					Usage: "apisix api key ",
				},
				&cli.StringFlag{
					Name:  "host",
					Value: "http://127.0.0.1:9180",
					Usage: "apisix host",
				},
				&cli.StringFlag{
					Name:  "domain",
					Value: "",
					Usage: "route domain name",
				},
				&cli.StringFlag{
					Name:  "livekit",
					Value: "",
					Usage: "livekit domain name",
				},
				&cli.StringFlag{
					Name:  "route-host",
					Value: "",
					Usage: "route host",
				},
			},
			Action: func(context *cli.Context) error {
				apiKey := context.String("key")
				host := context.String("host")
				direct := context.Bool("direct")
				domain := context.String("domain")
				livekitDomain := context.String("livekit")
				routeHost := context.String("route-host")

				baseURL := host + "/apisix/admin/routes/"

				client := apisix.NewApiClient(apiKey, baseURL)

				route := client.GetRoutes(domain, routeHost, livekitDomain, direct)

				for i, route := range route {
					resp, err := client.SendRequest("PUT", fmt.Sprintf("%d", i+1), route)
					if err != nil {
						fmt.Printf("Error sending request for route %d: %v\n", i+1, err)
						continue
					}
					fmt.Printf("Route %d created successfully: %s\n", i+1, resp)
				}
				return nil
			},
		},
		{
			Name:  "ssl",
			Usage: "init consul config",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "cert",
					Value: "./",
					Usage: "cert file path",
				},
				&cli.StringFlag{
					Name:  "private_key",
					Value: "./",
					Usage: "key file path",
				},
				&cli.StringFlag{
					Name:  "domain",
					Value: "",
					Usage: "your domain name",
				},
				&cli.IntFlag{
					Name:  "num",
					Value: 1,
					Usage: "ssl key num",
				},
				&cli.StringFlag{
					Name:  "key",
					Value: "edd1c9f034335f136f87ad84b625c8f1",
					Usage: "apisix api key ",
				},
				&cli.StringFlag{
					Name:  "host",
					Value: "http://127.0.0.1:9180",
					Usage: "apisix host",
				},
			},
			Action: func(context *cli.Context) error {
				certPath := context.String("cert")
				keyPath := context.String("private_key")
				domain := context.String("domain")
				apiKey := context.String("key")
				host := context.String("host")
				num := context.Int("num")

				cert, err := ioutil.ReadFile(certPath)
				if err != nil {
					return fmt.Errorf("failed to read certificate file: %v", err)
				}

				key, err := ioutil.ReadFile(keyPath)
				if err != nil {
					return fmt.Errorf("failed to read key file: %v", err)
				}

				if domain == "" {
					return fmt.Errorf("domain name is required")
				}
				client := apisix.NewApiClient(apiKey, host)
				if err := client.UpdateSSL(cert, key, []string{domain}, num); err != nil {
					fmt.Printf("Error updating SSL: %v\n", err)
				} else {
					fmt.Printf("SSL updated successfully\n")
				}
				return nil
			},
		},
	},
}
