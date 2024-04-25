# coss-cli

coss-cli 是一个用于快速部署cossim的命令行工具。

## 安装

确保你的Go环境已经安装并设置好，然后执行以下命令安装 coss-cli：

```bash
go get github.com/cossim/coss-cli
go install github.com/cossim/coss-cli
```

## 快速使用
```bash
coss-cli gen
docker-compose up -d
coss-cli route
```

## 使用注册中心
```bash
coss-cli gen --direct=false
docker-compose up -d
coss-cli config --path ./config/common --namespace common
coss-cli config --path ./config/service --namespace service
coss-cli route --direct=false --route-host=192.168.100.236
```
>如果要指定consul地址，可以使用`--host`参数,默认为`http://127.0.0.1:8500`
> 
> 更多详细参数请查看帮助`coss-cli config --help`
```
coss-cli config --path ./config/common --namespace common --host=http://127.0.0.1:8500
```

## Help
```
NAME:
   coss-cli - coss-cli is a command line tool for coss

USAGE:
   coss-cli [global options] command [command options] 

COMMANDS:
   config   init consul config 
   gen      gen coss config 
   route    init gateway route 
   ssl      init consul config 
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help

```
> config 用来初始化consul中的配置文件
> 
> gen 用来生成coss的配置文件
> 
> route 用来初始化网关的路由
> 
> ssl 将ssl证书上传到网关
