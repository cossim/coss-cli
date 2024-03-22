package config

import "fmt"

func GenServiceConfig(httpName string, grpcName string, httpPort string, grpcPort string) string {
	return fmt.Sprintf(`
system:
  environment: "dev" # dev、prod
  ssl: false # 是否启用ssl true的话不会使用port
  gateway_address: ""
  gateway_port: 
  gateway_address_dev: "127.0.0.1"
  gateway_port_dev: 8080

log:
  stdout: true
  level: -1
  file: "logs/app.log"
  format: "console" # console、json

email:
  enable: false
  smtp_server: "smtp.gmail.com"
  port: 25
  username: ""
  password: ""

livekit:
  address: ""
  port: 7880
  url: "wss://"
  api_key: ""
  secret_key: ""
  timeout: "1m"

cache:
  enable: true

mysql:
  address: "mysql"
  port: 3306
  username: "root"
  password: "Hitosea@123.."
  database: "coss"
  opts:
   allowNativePasswords: "true"
   timeout: "800ms"
   readTimeout: "200ms"
   writeTimeout: "800ms"
   parseTime: "true"
   loc: "Local"
   charset: "utf8mb4"

redis:
  proto: "tcp"
  address: "redis"
  port: 6379
  password: "Hitosea@123.."
#  protocol: 3

dtm:
  name: "dtm"
  address: "dtm"
  port: 36790

http:
  name: "%s"
  port: %s

grpc:
  name: "%s"
  port: %s

# 注册本服务
register:
  # 注册中心地址
  address: "consul"
  # 注册中心端口
  port: 8500
  tags: ["%s", "service", "%s service"]

discovers:
  user:
    name: "user_service"
    address: "user"
    port: 10002
    direct: true
  relation:
    name: "relation_service"
    address: "relation"
    port: 10001
    direct: true
  storage:
    name: "storage_service"
    address: "storage"
    port: 10006
    direct: true
  msg:
    name: "msg_service"
    address: "msg"
    port: 10000
    direct: true
  group:
    name: "group_service"
    address: "group"
    port: 10005
    direct: true

message_queue:
  name: "rabbitmq"
  username: "root"
  password: "Hitosea@123.."
  address: "rabbitmq"
  port: 5672

encryption:
  enabled: false
  name: coss-im
  email: max.mustermann@example.com
  passphrase: LongSecret
  rsaBits: 2048

multiple_device_limit:
  enable: false
  max: 1

oss:
  name: "minio"
  address: "minio"
  port: 9000
  accessKey: "root"
  secretKey: "Hitosea@123.."
  ssl: false
  presignedExpires: ""
  dial: "3000ms"
  timeout: "5000ms"
`, httpName, httpPort, grpcName, grpcPort, grpcName, grpcName)
}

func GenConsulServiceConfig(httpName, grpcName, httpPort, grpcPort string) string {
	return fmt.Sprintf(`
system:
  environment: "dev" # dev、prod
  ssl: false # 是否启用ssl true的话不会使用port
  gateway_address: "tuo.gezi.vip"
  gateway_port: 8080
  gateway_address_dev: "127.0.0.1"
  gateway_port_dev: 8080

log:
  stdout: true
  level: -1
  file: "logs/app.log"
  format: "console" # console、json

email:
  enable: false
  smtp_server: "smtp.gmail.com"
  port: 25
  username: ""
  password: ""


cache:
  enable: true

http:
  name: "%s"
  port: %s

grpc:
  name: "%s"
  port: %s

register:
  address: "consul"
  port: 8500
  tags: ["%s", "service", "%s"]

discovers:
  msg:
    name: "msg_service"
    port: 10000
    direct: false
  user:
    name: "user_service"
    port: 10002
    direct: false
  group:
    name: "group_service"
    port: 10005
    direct: false
  relation:
    name: "relation_service"
    port: 10001
    direct: false
  storage:
    name: "storage_service"
    port: 10006
    direct: false

encryption:
  enabled: false
  name: coss-im
  email: max.mustermann@example.com
  passphrase: LongSecret
  rsaBits: 2048

multiple_device_limit:
  enable: false
  max: 1
`, httpName, httpPort, grpcName, grpcPort, grpcName, grpcName)
}
