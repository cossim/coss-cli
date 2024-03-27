package config

import "fmt"

func GenCommonConfig(name CommonClientType) string {
	switch name {
	case CONSUL:
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
	case HIPUSH:
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
	case APISIX:
		return fmt.Sprintf(`apisix:
  node_listen: 9080              # APISIX listening port
  enable_ipv6: false

  enable_control: true
  control:
    ip: "0.0.0.0"
    port: 9092

deployment:
  admin:
    allow_admin: # https://nginx.org/en/docs/http/ngx_http_access_module.html#allow
      - 0.0.0.0/0
  etcd:
    host:                           # it's possible to define multiple etcd hosts addresses of the same etcd cluster.
      - http://etcd:2379          # multiple etcd address
    prefix: "/apisix"               # apisix configurations prefix
    timeout: 30                     # 30 seconds

discovery:
  consul:
    servers:
      - http://consul:8500
`)
	case APISIX_DASHBOARD:
		return fmt.Sprintf(`
conf:
  listen:
    host: 0.0.0.0
    port: 9000
  etcd:
    endpoints:
      - etcd:2379
    log:
      error_log:
        level: warn
        file_path: logs/error.log
authentication:
    secret: secret            
    expire_time: 3600    
    users:
    - username: admin  
      password: admin
    - username: user
      password: user
`)
	case LiveKit:
		return fmt.Sprintf(`
port: 7880
rtc:
  udp_port: 7882
  tcp_port: 7881
  port_range_start: 50000
  port_range_end: 60000
  # use_external_ip should be set to true for most cloud environments where
  # the host has a public IP address, but is not exposed to the process.
  # LiveKit will attempt to use STUN to discover the true IP, and advertise
  # that IP with its clients
  use_external_ip: true
  enable_loopback_candidate: false
keys:
  APIbsEc4M9ceob3: 
logging:
  json: false
  level: info
room:
 auto_create: true

turn:
 enabled: true
 tls_port: 443
 domain: 
 cert_file: 
 key_file: 
`)
	}
	return ""
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
		return fmt.Sprintf(`
oss:
  name: "minio"
  address: "minio"
  port: 9000
  accessKey: "root"
  secretKey: "Hitosea@123.."
  ssl: false
  #    presignedExpires: ""
`)
	}
	return ""
}
