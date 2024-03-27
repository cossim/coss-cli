package config

type CommonClientType string

const (
	CONSUL           CommonClientType = "consul"
	HIPUSH                            = "hipush"
	APISIX                            = "apisix"
	APISIX_DASHBOARD                  = "apisix-dashboard"
	LiveKit                           = "livekit"
)

var CommonClist = []CommonClientType{
	CONSUL,
	HIPUSH,
	APISIX,
	APISIX_DASHBOARD,
	LiveKit,
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

var ServiceList = []string{
	"relation",
	"group",
	"user",
	"msg",
	"storage",
	"live",
	"admin",
}

var HttpPort = map[string]string{
	"admin_bff":    "8087",
	"relation_bff": "8082",
	"group_bff":    "8084",
	"live_bff":     "8086",
	"user_bff":     "8083",
	"msg_bff":      "8081",
	"storage_bff":  "8085",
}

var HttpName = map[string]string{
	"admin":    "admin_bff",
	"relation": "relation_bff",
	"group":    "group_bff",
	"live":     "live_bff",
	"user":     "user_bff",
	"msg":      "msg_bff",
	"storage":  "storage_bff",
}

var GrpcName = map[string]string{
	"relation": "relation_service",
	"group":    "group_service",
	"user":     "user_service",
	"msg":      "msg_service",
	"storage":  "storage_service",
}

var GrpcPort = map[string]string{
	"relation_service": "10001",
	"group_service":    "10005",
	"user_service":     "10002",
	"msg_service":      "10000",
	"storage_service":  "10006",
}
