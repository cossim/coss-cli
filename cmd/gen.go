package cmd

import (
	"fmt"
	"github.com/cossim/coss-cli/config"
	"github.com/cossim/coss-cli/pkg/pgp"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

func gen(cCtx *cli.Context) error {
	direct := cCtx.Bool("direct")
	outputDir := cCtx.String("path")
	domain := cCtx.String("domain")
	enableSsl := cCtx.Bool("ssl")
	cacheDir := "./config"
	if outputDir != "" {
		cacheDir = outputDir + "/config"
	}

	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		err := os.Mkdir(cacheDir, 0755) // 创建文件夹并设置权限
		if err != nil {
			return err
		}
	}

	serviceDir := cacheDir + "/service"
	if _, err := os.Stat(serviceDir); os.IsNotExist(err) {
		err := os.Mkdir(serviceDir, 0755) // 创建文件夹并设置权限
		if err != nil {
			return err
		}
	}

	commonDir := cacheDir + "/common"
	if _, err := os.Stat(commonDir); os.IsNotExist(err) {
		err := os.Mkdir(commonDir, 0755) // 创建文件夹并设置权限
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

			configStr := config.GenServiceConfig(httpname, grpcname, httpport, grpcport, enableSsl, domain)
			filePath := filepath.Join(serviceDir, fmt.Sprintf("%s.yaml", name))
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

			configStr := config.GenConsulServiceConfig(httpname, grpcname, httpport, grpcport, enableSsl, domain)
			filePath := filepath.Join(serviceDir, fmt.Sprintf("%s.yaml", name))
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
			filePath := filepath.Join(commonDir, fmt.Sprintf("%s.yaml", name))
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
			filePath := filepath.Join(commonDir, fmt.Sprintf("%s.json", name))
			err := ioutil.WriteFile(filePath, []byte(configStr), 0644)
			if err != nil {
				fmt.Printf("写入文件 %s 失败：%v\n", filePath, err)
			} else {
				fmt.Printf("生成 %s 成功\n", filePath)
			}

		} else {
			filePath := filepath.Join(commonDir, fmt.Sprintf("%s.yaml", name))
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
}
