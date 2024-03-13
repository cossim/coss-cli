package config

import "fmt"

var InterfaceList = []string{
	"admin_bff",
	"relation_bff",
	"group_bff",
	"user_bff",
	"msg_bff",
	"storage_bff",
	"live_bff",
	"gateway",
}

var CommonClist = []string{
	"consul",
	"hipush",
}

var ServiceList = []string{
	"relation_service",
	"group_service",
	"user_service",
	"msg_service",
	"storage_service",
}

var InetrfaceConfig = map[string]string{
	"admin_bff":    "8087",
	"gateway":      "8080",
	"relation_bff": "8082",
	"group_bff":    "8084",
	"live_bff":     "8086",
	"user_bff":     "8083",
	"msg_bff":      "8081",
	"storage_bff":  "8085",
}

var InetrfaceName = map[string]string{
	"admin_bff":    "admin",
	"gateway":      "gateway",
	"relation_bff": "relation",
	"group_bff":    "group",
	"live_bff":     "live",
	"user_bff":     "user",
	"msg_bff":      "msg",
	"storage_bff":  "storage",
}

var ServiceName = map[string]string{
	"relation_service": "relation",
	"group_service":    "group",
	"user_service":     "user",
	"msg_service":      "msg",
	"storage_service":  "storage",
}

var ServiceConfig = map[string]string{
	"relation_service": "10001",
	"group_service":    "10005",
	"user_service":     "10002",
	"msg_service":      "10000",
	"storage_service":  "10006",
}

func GenInterfaceConfig(name string, port string, direct bool) string {
	return fmt.Sprintf(`
system:
  environment: "prod" # dev、prod
  ssl: false # 是否启用ssl true的话不会使用port
  avatar_file_path: "/.cache/"
  avatar_file_path_dev:
  gateway_address: "43.229.28.107"
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
  dsn: "root:Hitosea@123..@tcp(mysql:33066)/coss?allowNativePasswords=true&timeout=800ms&readTimeout=200ms&writeTimeout=800ms&parseTime=true&loc=Local&charset=utf8mb4"

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
  address: "0.0.0.0"
  port: %s

grpc:
  address: "0.0.0.0"
  port:

# 注册本服务
register:
  # 服务注册名称
  name: "%s"
  # 注册中心地址
  address: "consul"
  # 注册中心端口
  port: 8500
  tags: ["%s", "bff", "%s interface"]
  # 启用服务发现 默认为true
  #discover: true
  # 启用服务注册 默认为true
  #register: true

discovers:
  user:
    name: "user_service"
    address: "user_service"
    port: 10002
    # 不使用服务发现，使用addr直接连接
    # 默认为false
    direct: %t
  relation:
    name: "relation_service"
    address: "relation_service"
    port: 10001
    direct: %t
  storage:
    name: "storage_service"
    address: "storage_service"
    port: 10006
    direct: %t
  gateway:
    name: "gateway"
    address: "gateway"
    port: 8080
    direct: %t
  msg:
    name: "msg_service"
    address: "msg_service"
    port: 10000
    direct: %t
  group:
    name: "group_service"
    address: "group_service"
    port: 10005
    direct: %t

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
  minio:
    name: "minio"
    address: "minio"
    port: 9000
    accessKey: "root"
    secretKey: "Hitosea@123.."
    ssl: false
    presignedExpires: ""
    dial: "3000ms"
    timeout: "5000ms"

`, port, name, name, name, direct, direct, direct, direct, direct, direct)
}

func GenCommonConfig(name string) string {
	switch name {
	case "consul":
		return fmt.Sprintf(`
{
  "acl": {
    "enabled": false,
    "default_policy": "deny",
    "enable_token_persistence": true,
    "tokens": {
      "agent": "",
      "default": "",
      "master": ""
    }
  }
}`)
	case "hipush":
		return fmt.Sprintf(`
ios:
  - enabled: true
    # 应用程序的 Bundle ID
    # ios capacitor.config文件中的appId 例如com.hitosea.apptest
    appid: ""
    # APNs 密钥文件路径
    key_path: ""
    # 密钥类型（例如：pem）
    key_type: pem
    # 密钥文件的密码（如果有）
    password: ""
    # 是否为生产环境
    production: false
    # 最大并发推送数
    max_concurrent_pushes: 100
    # 最大重试次数
    max_retry: 5
    # 密钥 ID
    key_id: ""
    # 开发团队 ID
    team_id: ""
  - enabled: true
    appid: ""
    key_path: key.pem
    key_type: pem
    password: ""
    production: false
    max_concurrent_pushes: 100
    max_retry: 0
    key_id: ""
    team_id: ""
huawei:
  - enabled: true
    appid: ""
    appsecret: ""
    max_retry: 0
`)
	}
	return ""
}

func GenServiceConfig(name string, port string) string {
	return fmt.Sprintf(`
log:
  stdout: true
  level: -1
  file: "logs/app.log"
  format: "console" # console、json

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

grpc:
  address: "%s"
  port: %s

# 注册本服务
register:
  # 服务注册名称
  name: "%s"
  # 注册中心地址
  address: "consul"
  # 注册中心端口
  port: 8500
  tags: ["%s", "service", "%s service"]
  # 启用服务发现 默认为true
  #discover: true
  # 启用服务注册 默认为true
  #register: true
`, name, port, name, name, name)
}

func GenDockerCompose(consul bool) string {
	if consul {
		return `
version: '3.9'
services:
  mysql:
    container_name: cossim_mysql
    image: mysql:5.7
    volumes:
      - ./data/var/lib/mysql:/var/lib/mysql
    command: [
      '--character-set-server=utf8mb4',
    ]
    expose:
      - "3306"
    ports:
      - "33066:3306"
    environment:
      MYSQL_ROOT_PASSWORD: "Hitosea@123.."
      MYSQL_DATABASE: coss
      MYSQL_USER: coss
      MYSQL_PASSWORD: "Hitosea@123.."
      MYSQL_TCP_PORT: '3306'
      MYSQL_ROOT_HOST: '%'
#      MARIADB_AUTO_UPGRADE: 'true'
#      MARIADB_DISABLE_UPGRADE_BACKUP: 'true'
    healthcheck:
      test: mysqladmin ping -h mysql -P 3306 -p$$MYSQL_ROOT_PASSWORD
      interval: 5s
      timeout: 10s
      retries: 10
      start_period: 30s
  minio:
    image: hub.hitosea.com/cossim/minio
    container_name: cossim_minio
    ports:
      - "9000:9000"
      - "9001:9001"
    expose:
      - "9000"
      - "9001"
    environment:
      MINIO_ROOT_USER: root
      MINIO_ROOT_PASSWORD: Hitosea@123..
    volumes:
    - ./data/var/lib/minio:/data
    command: server /data --console-address ":9001"
  rabbitmq:
    image: "rabbitmq:management"
    container_name: cossim_rabbitmq
    hostname: rabbitmq3-management-master
    logging:
      driver: json-file
      options:
        max-size: "100m"
        max-file: "1"
    volumes:
      - "./data/var/lib/rabbitmq:/var/lib/rabbitmq"
    ports:
      - "5672:5672"
      - "15672:15672"
    environment:
      RABBITMQ_DEFAULT_USER: root
      RABBITMQ_DEFAULT_PASS: Hitosea@123..
  redis:
    image: redis:latest
    container_name: cossim_redis
    ports:
      - 6379:6379
    command: redis-server --requirepass Hitosea@123..
  dtm:
    container_name: cossim_dtm
    image: hub.hitosea.com/cossim/dtm
    ports:
      - '36789:36789'
      - '36790:36790'
  consul:
    image: hub.hitosea.com/cossim/consul:latest
    container_name: cossim_consul
    volumes:
      - ./data/var/lib/consul:/consul/data
      - ./config/common/consul.json:/etc/consul/consul-config.json
    command: consul agent -server -bootstrap-expect=1 -client=0.0.0.0 -ui -data-dir=/consul/data -config-dir=/etc/consul
    ports:
      - '8300:8300'
      - '8301:8301'
      - '8301:8301/udp'
      - '8500:8500'
      - '8600:8600'
      - '8600:8600/udp'
  hipush:
    container_name: cossim_push
    image: hub.hitosea.com/cossim/hipush
    volumes:
      - ./config/common/hipush.yaml:/config/config.yaml
      # 如果启用了ios推送，需要将推送证书挂载到容器中
      #- /Users/macos-15/Desktop/key/AuthKey_UU2D9Z4ANF.p8:/config/key/AuthKey_UU2D9Z4ANF.p8
    command:
      - "-config"
      - "/config/config.yaml"
    ports:
      - "7070:7070"
      - "7071:7071"
    restart: on-failure
  gateway:
    container_name: cossim_gateway
    image: hub.hitosea.com/cossim/gateway-interface
#    volumes:
#      - ./config/interface/gateway.yaml:/config/config.yaml
    command:
#      - "--config"
#      - "/config/config.yaml"
      - "--discover"
      - "--remote-config"
      - "--config-center-addr"
      - "consul:8500"
    ports:
      - "8080:8080"
    environment:
      CONSUL_HTTP_TOKEN:
    restart: on-failure
  msg_bff:
    container_name: cossim_msg_bff
    image: hub.hitosea.com/cossim/msg-interface
    command:
#      - "--config"
#      - "/config/config.yaml"
      - "--discover"
      - "--register"
      - "--remote-config"
      - "--config-center-addr"
      - "consul:8500"
    volumes:
#      - ./config/interface/msg.yaml:/config/config.yaml
      - ./config/pgp:/.cache
    depends_on:
      - dtm
      - redis
      - rabbitmq
      - group_service
      - relation_service
      - user_service
      - msg_service
    environment:
      CONSUL_HTTP_TOKEN:
    restart: on-failure
  admin_bff:
    container_name: cossim_admin_bff
    image: hub.hitosea.com/cossim/admin-admin
    command:
#      - "--config"
#      - "/config/config.yaml"
      - "--discover"
      - "--register"
      - "--remote-config"
      - "--config-center-addr"
      - "consul:8500"
    volumes:
#      - ./config/interface/admin.yaml:/config/config.yaml
      - ./config/pgp:/.cache
    depends_on:
      - dtm
      - redis
      - rabbitmq
      - group_service
      - relation_service
      - user_service
      - msg_service
    environment:
      CONSUL_HTTP_TOKEN:
    restart: on-failure
  relation_bff:
    container_name: cossim_relation_bff
    image: hub.hitosea.com/cossim/relation-interface
    command:
#      - "--config"
#      - "/config/config.yaml"
      - "--discover"
      - "--register"
      - "--remote-config"
      - "--config-center-addr"
      - "consul:8500"
    volumes:
#      - ./config/interface/relation.yaml:/config/config.yaml
      - ./config/pgp:/.cache
    depends_on:
      mysql:
        condition: service_healthy
    environment:
      CONSUL_HTTP_TOKEN:
    restart: on-failure
  user_bff:
    container_name: cossim_user_bff
    image: hub.hitosea.com/cossim/user-interface
    command:
#    - "--config"
#    - "/config/config.yaml"
    - "--discover"
    - "--register"
    - "--remote-config"
    - "--config-center-addr"
    - "consul:8500"
    volumes:
#      - ./config/interface/user.yaml:/config/config.yaml
      - ./config/pgp:/.cache
    depends_on:
      - dtm
      - redis
      - relation_service
      - user_service
    environment:
      CONSUL_HTTP_TOKEN:
    restart: on-failure
  group_bff:
    container_name: cossim_group_bff
    image: hub.hitosea.com/cossim/group-interface
    command:
#      - "--config"
#      - "/config/config.yaml"
      - "--discover"
      - "--register"
      - "--remote-config"
      - "--config-center-addr"
      - "consul:8500"
    volumes:
#      - ./config/interface/group.yaml:/config/config.yaml
      - ./config/pgp:/.cache
    depends_on:
      - dtm
      - redis
      - rabbitmq
      - group_service
      - relation_service
      - user_service
      - msg_service
    environment:
      CONSUL_HTTP_TOKEN:
    restart: on-failure
  storage_bff:
    container_name: cossim_storage_bff
    image: hub.hitosea.com/cossim/storage-interface
    command:
      - "--discover"
      - "--register"
      - "--remote-config"
      - "--config-center-addr"
      - "consul:8500"
    volumes:
      - ./config/pgp:/.cache
    depends_on:
      - consul
      - minio
      - redis
      - storage_service
    environment:
      CONSUL_HTTP_TOKEN:
    restart: on-failure
  live_bff:
    container_name: cossim_live_bff
    image: hub.hitosea.com/cossim/live-interface
    command:
      - "--discover"
      - "--register"
      - "--remote-config"
      - "--config-center-addr"
      - "consul:8500"
    volumes:
      - ./config/pgp:/.cache
    depends_on:
      - consul
      - redis
      - user_service
      - relation_service
    environment:
      CONSUL_HTTP_TOKEN:
    restart: on-failure
  msg_service:
    container_name: cossim_msg_service
    image: hub.hitosea.com/cossim/msg-service
    command:
#      - "--config"
#      - "/config/config.yaml"
      - "--discover"
      - "--register"
      - "--remote-config"
      - "--config-center-addr"
      - "consul:8500"
#    volumes:
#      - ./config/service/msg.yaml:/config/config.yaml
    depends_on:
      mysql:
        condition: service_healthy
    environment:
      CONSUL_HTTP_TOKEN:
    restart: on-failure
  relation_service:
    container_name: cossim_relation_service
    image: hub.hitosea.com/cossim/relation-service
    command:
#      - "--config"
#      - "/config/config.yaml"
      - "--discover"
      - "--register"
      - "--remote-config"
      - "--config-center-addr"
      - "consul:8500"
#    volumes:
#      - ./config/service/relation.yaml:/config/config.yaml
    depends_on:
      mysql:
        condition: service_healthy
    environment:
      CONSUL_HTTP_TOKEN:
    restart: on-failure
  user_service:
    container_name: cossim_user_service
    image: hub.hitosea.com/cossim/user-service
    command:
#      - "--config"
#      - "/config/config.yaml"
      - "--discover"
      - "--register"
      - "--remote-config"
      - "--config-center-addr"
      - "consul:8500"
#    volumes:
#      - ./config/service/user.yaml:/config/config.yaml
    depends_on:
      mysql:
        condition: service_healthy
    environment:
      CONSUL_HTTP_TOKEN:
    restart: on-failure
  group_service:
    container_name: cossim_group_service
    image: hub.hitosea.com/cossim/group-service
    command:
#      - "--config"
#      - "/config/config.yaml"
      - "--discover"
      - "--register"
      - "--remote-config"
      - "--config-center-addr"
      - "consul:8500"
#    volumes:
#      - ./config/service/group.yaml:/config/config.yaml
    depends_on:
      mysql:
        condition: service_healthy
    environment:
      CONSUL_HTTP_TOKEN:
    restart: on-failure
  storage_service:
    container_name: cossim_storage_service
    image: hub.hitosea.com/cossim/storage-service
    command:
#      - "--config"
#      - "/config/config.yaml"
      - "--discover"
      - "--register"
      - "--remote-config"
      - "--config-center-addr"
      - "consul:8500"
#    volumes:
#      - ./config/service/storage.yaml:/config/config.yaml
    depends_on:
      mysql:
        condition: service_healthy
    environment:
      CONSUL_HTTP_TOKEN:
    restart: on-failure
`
	} else {
		return `
version: '3.9'
services:
  mysql:
    container_name: cossim_mysql
    image: mysql:5.7
    volumes:
      - ./data/var/lib/mysql:/var/lib/mysql
    command: [
      '--character-set-server=utf8mb4',
    ]
    expose:
      - "3306"
    ports:
      - "33066:3306"
    environment:
      MYSQL_ROOT_PASSWORD: "Hitosea@123.."
      MYSQL_DATABASE: coss
      MYSQL_USER: coss
      MYSQL_PASSWORD: "Hitosea@123.."
      MYSQL_TCP_PORT: '3306'
      MYSQL_ROOT_HOST: '%'
#      MARIADB_AUTO_UPGRADE: 'true'
#      MARIADB_DISABLE_UPGRADE_BACKUP: 'true'
    healthcheck:
      test: mysqladmin ping -h mysql -P 3306 -p$$MYSQL_ROOT_PASSWORD
      interval: 5s
      timeout: 10s
      retries: 10
      start_period: 30s
  minio:
    image: hub.hitosea.com/cossim/minio
    container_name: cossim_minio
    ports:
      - "9000:9000"
      - "9001:9001"
    expose:
      - "9000"
      - "9001"
    environment:
      MINIO_ROOT_USER: root
      MINIO_ROOT_PASSWORD: Hitosea@123..

    volumes:
    - ./data/var/lib/minio:/data
    command: server /data --console-address ":9001"
  rabbitmq:
    image: "rabbitmq:management"
    container_name: cossim_rabbitmq
    hostname: rabbitmq3-management-master
    logging:
      driver: json-file
      options:
        max-size: "100m"
        max-file: "1"
    volumes:
      - "./data/var/lib/rabbitmq:/var/lib/rabbitmq"
    ports:
      - "5672:5672"
      - "15672:15672"
    environment:
      RABBITMQ_DEFAULT_USER: root
      RABBITMQ_DEFAULT_PASS: Hitosea@123..
  redis:
    image: redis:latest
    container_name: cossim_redis
    ports:
      - 6379:6379
    command: redis-server --requirepass Hitosea@123..
  dtm:
    container_name: cossim_dtm
    image: hub.hitosea.com/cossim/dtm
    ports:
      - '36789:36789'
      - '36790:36790'
  consul:
    image: hub.hitosea.com/cossim/consul:latest
    container_name: cossim_consul
    volumes:
      - ./data/var/lib/consul:/consul/data
      - ./config/common/consul.json:/etc/consul/consul-config.json
    command: consul agent -server -bootstrap-expect=1 -client=0.0.0.0 -ui -data-dir=/consul/data -config-dir=/etc/consul
    ports:
      - '8300:8300'
      - '8301:8301'
      - '8301:8301/udp'
      - '8500:8500'
      - '8600:8600'
      - '8600:8600/udp'
  hipush:
    container_name: cossim_push
    image: hub.hitosea.com/cossim/hipush
    volumes:
      - ./config/common/hipush.yaml:/config/config.yaml
      # 如果启用了ios推送，需要将推送证书挂载到容器中
      #- /Users/macos-15/Desktop/key/AuthKey_UU2D9Z4ANF.p8:/config/key/AuthKey_UU2D9Z4ANF.p8
    command:
      - "-config"
      - "/config/config.yaml"
    ports:
      - "7070:7070"
      - "7071:7071"
    restart: on-failure
  gateway:
    container_name: cossim_gateway
    image: hub.hitosea.com/cossim/gateway-interface
    volumes:
      - ./config/interface/gateway.yaml:/config/config.yaml
    command:
      - "--config"
      - "/config/config.yaml"
      - "--discover"
    ports:
      - "8080:8080"
    environment:
      CONSUL_HTTP_TOKEN:
    restart: on-failure
  msg_bff:
    container_name: cossim_msg_bff
    image: hub.hitosea.com/cossim/msg-interface
    command:
      - "--config"
      - "/config/config.yaml"
      - "--discover"
    volumes:
      - ./config/interface/msg.yaml:/config/config.yaml
      - ./config/pgp:/.cache
    depends_on:
      - dtm
      - redis
      - rabbitmq
      - group_service
      - relation_service
      - user_service
      - msg_service
    environment:
      CONSUL_HTTP_TOKEN:
    restart: on-failure
  admin_bff:
    container_name: cossim_admin_bff
    image: hub.hitosea.com/cossim/admin-admin
    command:
      - "--config"
      - "/config/config.yaml"
      - "--discover"
    volumes:
      - ./config/interface/admin.yaml:/config/config.yaml
      - ./config/pgp:/.cache
    depends_on:
      - dtm
      - redis
      - rabbitmq
      - group_service
      - relation_service
      - user_service
      - msg_service
    environment:
      CONSUL_HTTP_TOKEN:
    restart: on-failure
  relation_bff:
    container_name: cossim_relation_bff
    image: hub.hitosea.com/cossim/relation-interface
    command:
      - "--config"
      - "/config/config.yaml"
      - "--discover"
    volumes:
      - ./config/interface/relation.yaml:/config/config.yaml
      - ./config/pgp:/.cache
    depends_on:
      mysql:
        condition: service_healthy
    environment:
      CONSUL_HTTP_TOKEN:
    restart: on-failure
  user_bff:
    container_name: cossim_user_bff
    image: hub.hitosea.com/cossim/user-interface
    command:
      - "--config"
      - "/config/config.yaml"
      - "--discover"
#      - "--register"
#      - "--remote-config"
#      - "--config-center-addr"
#      - "consul:8500"
    volumes:
      - ./config/interface/user.yaml:/config/config.yaml
      - ./config/pgp:/.cache
    depends_on:
      - dtm
      - redis
      - relation_service
      - user_service
    environment:
      CONSUL_HTTP_TOKEN:
    restart: on-failure
  group_bff:
    container_name: cossim_group_bff
    image: hub.hitosea.com/cossim/group-interface
    command:
      - "--config"
      - "/config/config.yaml"
      - "--discover"
    volumes:
      - ./config/interface/group.yaml:/config/config.yaml
      - ./config/pgp:/.cache
    depends_on:
      - minio
      - dtm
      - redis
      - rabbitmq
      - group_service
      - relation_service
      - user_service
      - msg_service
    environment:
      CONSUL_HTTP_TOKEN:
    restart: on-failure
  storage_bff:
    container_name: cossim_storage_bff
    image: hub.hitosea.com/cossim/storage-interface
    command:
      - "--config"
      - "/config/config.yaml"
      - "--discover"
    volumes:
      - ./config/pgp:/.cache
      - ./config/interface/storage.yaml:/config/config.yaml
    depends_on:
      - consul
      - minio
      - redis
      - storage_service
    environment:
      CONSUL_HTTP_TOKEN:
    restart: on-failure
  live_bff:
    container_name: cossim_live_bff
    image: hub.hitosea.com/cossim/live-interface
    command:
      - "--config"
      - "/config/config.yaml"
      - "--discover"
    volumes:
      - ./config/pgp:/.cache
      - ./config/interface/live.yaml:/config/config.yaml
    depends_on:
      - consul
      - redis
      - user_service
      - relation_service
    environment:
      CONSUL_HTTP_TOKEN:
    restart: on-failure
  msg_service:
    container_name: cossim_msg_service
    image: hub.hitosea.com/cossim/msg-service
    command:
      - "--config"
      - "/config/config.yaml"
      - "--discover"
    volumes:
      - ./config/service/msg.yaml:/config/config.yaml
    depends_on:
      mysql:
        condition: service_healthy
    environment:
      CONSUL_HTTP_TOKEN:
    restart: on-failure
  relation_service:
    container_name: cossim_relation_service
    image: hub.hitosea.com/cossim/relation-service
    command:
      - "--config"
      - "/config/config.yaml"
      - "--discover"
    volumes:
      - ./config/service/relation.yaml:/config/config.yaml
    depends_on:
      mysql:
        condition: service_healthy
    environment:
      CONSUL_HTTP_TOKEN:
    restart: on-failure
  user_service:
    container_name: cossim_user_service
    image: hub.hitosea.com/cossim/user-service
    command:
      - "--config"
      - "/config/config.yaml"
      - "--discover"
    volumes:
      - ./config/service/user.yaml:/config/config.yaml
    depends_on:
      mysql:
        condition: service_healthy
    environment:
      CONSUL_HTTP_TOKEN:
    restart: on-failure
  group_service:
    container_name: cossim_group_service
    image: hub.hitosea.com/cossim/group-service
    command:
      - "--config"
      - "/config/config.yaml"
      - "--discover"
    volumes:
      - ./config/service/group.yaml:/config/config.yaml
    depends_on:
      mysql:
        condition: service_healthy
    environment:
      CONSUL_HTTP_TOKEN:
    restart: on-failure
  storage_service:
    container_name: cossim_storage_service
    image: hub.hitosea.com/cossim/storage-service
    command:
      - "--config"
      - "/config/config.yaml"
      - "--discover"
    volumes:
      - ./config/service/storage.yaml:/config/config.yaml
    depends_on:
      mysql:
        condition: service_healthy
    environment:
      CONSUL_HTTP_TOKEN:
    restart: on-failure
`
	}
}

type ConsulCommon string

const (
	Dtm          ConsulCommon = "dtm"
	MessageQueue ConsulCommon = "message_queue"
	Mysql        ConsulCommon = "mysql"
	Oss          ConsulCommon = "oss"
	Redis        ConsulCommon = "redis"
)

var ConsulCommonList = []ConsulCommon{
	Dtm, MessageQueue, Mysql, Oss, Redis,
}

func GenConsulCommonConfig(name ConsulCommon) string {
	switch name {
	case Dtm:
		return `
dtm:
  name: "dtm"
  address: "dtm"
  port: 36790
`
	case Redis:
		return `
redis:
  proto: "tcp"
  address: "redis"
  port: 6379
  password: "Hitosea@123.."
#  protocol: 3
`
	case MessageQueue:
		return `
message_queue:
  name: "rabbitmq"
  username: "root"
  password: "Hitosea@123.."
  address: "rabbitmq"
  port: 5672
`
	case Mysql:
		return `
mysql:
  address: "mysql"
  port: 3306
  username: "root"
  password: "Hitosea@123.."
  database: "coss"
  opts:
    allowNativePasswords: "true"
    timeout: "1000ms"
    readTimeout: "500ms"
    writeTimeout: "1000ms"
    parseTime: "true"
    loc: "Local"
    charset: "utf8mb4"  
`
	case Oss:
		return `
oss:
  minio:
    name: "minio"
    address: "minio"
    port: 9000
    accessKey: "root"
    secretKey: "Hitosea@123.."
    ssl: false
#    presignedExpires: ""
`
	}
	return ""
}

func GenConsulInterfaceConfig(name string, port string, direct bool) string {
	if name != "gateway" {
		return fmt.Sprintf(`
system:
  environment: "prod" # dev、prod
  ssl: false # 是否启用ssl true的话不会使用port
  avatar_file_path: "/.cache/"
  avatar_file_path_dev:
  gateway_address: "43.229.28.107"
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

livekit:
  address: http://43.229.28.107
  port: 7880
  url: wss://coss.gezi.vip
  api_key: APIbsEc4M9ceob3
  secret_key: Op5frnZoFRUlG0lnCUhlh12I1XfdrB90ZEji07fXQZbB
  timeout: 2m

cache:
  enable: true

http:
  address: "%s"
  port: %s

grpc:
  address: "0.0.0.0"
  port:

# 注册本服务
register:
  # 服务注册名称
  name: "%s"
  # 注册中心地址
  address: "consul"
  # 注册中心端口
  port: 8500
  tags: ["%s", "bff", "%s interface"]
  # 启用服务发现 默认为true
  #discover: true
  # 启用服务注册 默认为true
  #register: true

discovers:
  user:
    name: "user_service"
    address: "user_service"
    port: 10002
    # 不使用服务发现，使用addr直接连接
    # 默认为false
    direct: %t
  relation:
    name: "relation_service"
    address: "relation_service"
    port: 10001
    direct: %t
  storage:
    name: "storage_service"
    address: "storage_service"
    port: 10006
    direct: %t
  gateway:
    name: "gateway"
    address: "gateway"
    port: 8080
    direct: %t
  msg:
    name: "msg_service"
    address: "msg_service"
    port: 10000
    direct: %t
  group:
    name: "group_service"
    address: "group_service"
    port: 10005
    direct: %t


encryption:
  enabled: false
  name: coss-im
  email: max.mustermann@example.com
  passphrase: LongSecret
  rsaBits: 2048

multiple_device_limit:
  enable: false
  max: 1
`, name, port, name, name, name, direct, direct, direct, direct, direct, direct)
	} else {
		return fmt.Sprintf(`
system:
  environment: "prod" # dev、prod
  ssl: false # 是否启用ssl true的话不会使用port
  avatar_file_path: "/.cache/"
  avatar_file_path_dev:
  gateway_address: "43.229.28.107"
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

livekit:
  address: http://43.229.28.107
  port: 7880
  url: wss://coss.gezi.vip
  api_key: APIbsEc4M9ceob3
  secret_key: Op5frnZoFRUlG0lnCUhlh12I1XfdrB90ZEji07fXQZbB
  timeout: 2m

cache:
  enable: true

http:
  address: "%s"
  port: %s

grpc:
  address: "0.0.0.0"
  port:

# 注册本服务
register:
  # 服务注册名称
  name: "%s"
  # 注册中心地址
  address: "consul"
  # 注册中心端口
  port: 8500
  tags: ["%s", "bff", "%s interface"]
  # 启用服务发现 默认为true
  #discover: true
  # 启用服务注册 默认为true
  #register: true

discovers:
  user:
    name: "user_service"
    address: "user_service"
    port: 10002
    # 不使用服务发现，使用addr直接连接
    # 默认为false
    direct: %t
  relation:
    name: "relation_service"
    address: "relation_service"
    port: 10001
    direct: %t
  storage:
    name: "storage_service"
    address: "storage_service"
    port: 10006
    direct: %t
  msg:
    name: "msg_service"
    address: "msg_service"
    port: 10000
    direct: %t
  group:
    name: "group_service"
    address: "group_service"
    port: 10005
    direct: %t

encryption:
  enabled: false
  name: coss-im
  email: max.mustermann@example.com
  passphrase: LongSecret
  rsaBits: 2048

multiple_device_limit:
  enable: false
  max: 1
`, name, port, name, name, name, direct, direct, direct, direct, direct)
	}
}

func GenConsulServiceConfig(name string, port string) string {
	return fmt.Sprintf(`
log:
  stdout: true
  level: -1
  file: "logs/app.log"
  format: "console" # console、json

grpc:
  address: "%s"
  port: %s

# 注册本服务
register:
  # 服务注册名称
  name: "%s"
  # 注册中心地址
  address: "consul"
  # 注册中心端口
  port: 8500
  tags: ["%s", "service", "%s service"]
  # 启用服务发现 默认为true
  #discover: true
  # 启用服务注册 默认为true
  #register: true
`, name, port, name, name, name)
}
