package apisix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cossim/coss-cli/config"
	"io/ioutil"
	"net/http"
	"sync"
)

var RouteName = []string{
	"user",
	"group",
	"relation",
	"admin",
	"msg",
	"storage",
	"live",
}

var WsRouteName = []string{
	"ws",
	"livekit",
}

// ApiService 定义单例结构体
type ApiService struct {
	apiKey  string
	baseURL string
	client  *http.Client
}

// ApiClient 定义客户端
type ApiClient struct {
	apiService *ApiService
}

var (
	apiServiceInstance *ApiService
	once               sync.Once
)

// NewApiClient 创建一个新的 ApiClient 实例
func NewApiClient(apiKey, baseURL string) *ApiClient {
	once.Do(func() {
		apiServiceInstance = &ApiService{
			apiKey:  apiKey,
			baseURL: baseURL,
			client:  &http.Client{},
		}
	})
	return &ApiClient{apiService: apiServiceInstance}
}

func getUpstream(node string, direct bool) string {
	if !direct {
		return fmt.Sprintf(`{"service_name": "%s", "type": "roundrobin", "discovery_type": "consul"}`, node)
	}
	return fmt.Sprintf(`{"type": "roundrobin", "nodes": {"%s": 1}}`, node)
}

func (c *ApiClient) GetRoute(uri string, node string, domain string, serviceName string, ws bool, direct bool) string {
	if domain == "" {
		if ws {
			return fmt.Sprintf(`{"uri": "%s", "name": "%s", "enable_websocket": true, "upstream": %s, "plugins": {"api-breaker": {
            "break_response_code": 502,
			"max_breaker_sec": 5,
            "unhealthy": {
                "http_statuses": [500, 503],
                "failures": 3
            },
            "healthy": {
                "http_statuses": [200],
                "successes": 1
            }
        },"limit-count": {"count": 5, "time_window": 10, "rejected_code": 503, "_meta": {"disable": true}}, "cors": {}}}`, uri, serviceName, getUpstream(node, direct))
		}
		return fmt.Sprintf(`{"uri": "%s",  "name": "%s","upstream": %s, "plugins": {"api-breaker": {
            "break_response_code": 502,
			"max_breaker_sec": 5,
            "unhealthy": {
                "http_statuses": [500, 503],
                "failures": 3
            },
            "healthy": {
                "http_statuses": [200],
                "successes": 1
            }
        },"limit-count": {"count": 5, "time_window": 10, "rejected_code": 503, "_meta": {"disable": true}}, "cors": {}}}`, uri, serviceName, getUpstream(node, direct))
	}
	if ws {
		return fmt.Sprintf(`{"uri": "%s",  "name": "%s","host": "%s", "enable_websocket": true, "upstream": %s, "plugins": {"api-breaker": {
            "break_response_code": 502,
			"max_breaker_sec": 5,
            "unhealthy": {
                "http_statuses": [500, 503],
                "failures": 3
            },
            "healthy": {
                "http_statuses": [200],
                "successes": 1
            }
        },"limit-count": {"count": 5, "time_window": 10, "rejected_code": 503, "_meta": {"disable": true}}, "cors": {}}}`, uri, serviceName, domain, getUpstream(node, direct))
	}
	return fmt.Sprintf(`{"uri": "%s", "name": "%s","host": "%s", "upstream": %s, "plugins": {"api-breaker": {
            "break_response_code": 502,
			"max_breaker_sec": 5,
            "unhealthy": {
                "http_statuses": [500, 503],
                "failures": 3
            },
            "healthy": {
                "http_statuses": [200],
                "successes": 1
            }
        },"limit-count": {"count": 5, "time_window": 10, "rejected_code": 503, "_meta": {"disable": true}}, "cors": {}}}`, uri, serviceName, domain, getUpstream(node, direct))
}

func (c *ApiClient) GetLiveKitRoute(domain string) string {
	route := ""
	if domain != "" {
		route = fmt.Sprintf(`{
    "uri": "/*",
    "name": "livekit",
    "host": "%s",
    "enable_websocket": true,
    "upstream": {
        "type": "roundrobin",
        "nodes": [{
			"host": "livekit",
        	"port": 7880,
        	"weight": 1
        }]
    },
    "plugins": {
		"api-breaker": {
            "break_response_code": 502,
			"max_breaker_sec": 5,
            "unhealthy": {
                "http_statuses": [500, 503],
                "failures": 3
            },
            "healthy": {
                "http_statuses": [200],
                "successes": 1
            }
        },
        "limit-count": {
            "count": 5,
            "time_window": 10,
            "rejected_code": 503,
            "_meta": {
                "disable": true
            }
        },
        "cors": {}
    }
}`, domain)
	} else {
		route = fmt.Sprintf(`{
    "uri": "/*",
    "name": "livekit",
    "enable_websocket": true,
    "upstream": {
        "type": "roundrobin",
        "nodes": [{
			"host": "livekit",
        	"port": 7880,
        	"weight": 1
        }]
    },
    "plugins": {
		"api-breaker": {
            "break_response_code": 502,
			"max_breaker_sec": 5,
            "unhealthy": {
                "http_statuses": [500, 503],
                "failures": 3
            },
            "healthy": {
                "http_statuses": [200],
                "successes": 1
            }
        },
        "limit-count": {
            "count": 5,
            "time_window": 10,
            "rejected_code": 503,
            "_meta": {
                "disable": true
            }
        },
        "cors": {}
    }
}`)
	}
	return route
}

// SendRequest 发送请求
func (c *ApiClient) SendRequest(method, endpoint, payload string) ([]byte, error) {
	url := c.apiService.baseURL + endpoint
	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", c.apiService.apiKey)

	resp, err := c.apiService.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// UpdateSSL 更新SSL证书
func (c *ApiClient) UpdateSSL(cert, key []byte, snis []string, num int) error {
	payload := map[string]interface{}{
		"cert": string(cert),
		"key":  string(key),
		"snis": snis,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	_, err = c.SendRequest("PUT", fmt.Sprintf("/apisix/admin/ssls/%d", num), string(jsonPayload))
	return err
}

func (c *ApiClient) GetRoutes(domain string, host, livekitDomain string, direct bool) []string {
	var routes []string
	rhost := ""
	for _, v := range RouteName {
		rhost = v
		uri := fmt.Sprintf("/api/v1/%s*", v)
		if !direct {
			routes = append(routes, c.GetRoute(uri, config.HttpName[v], domain, v, false, direct))
			continue
		}
		if host != "" {
			rhost = host
		}
		routes = append(routes, c.GetRoute(uri, fmt.Sprintf("%s:%s", rhost, config.HttpPort[config.HttpName[v]]), domain, v, false, direct))
	}
	for _, s := range WsRouteName {
		uri := ""
		serviceName := ""
		switch s {
		case "ws":
			serviceName = "push"
			uri = fmt.Sprintf("/api/v1/%s*", serviceName)
		case "livekit":
			route := c.GetLiveKitRoute(livekitDomain)
			routes = append(routes, route)
			continue
		}
		rhost = serviceName
		if !direct {
			routes = append(routes, c.GetRoute(uri, config.HttpName[serviceName], domain, serviceName, true, direct))
			continue
		}
		if host != "" {
			rhost = host
		}
		routes = append(routes, c.GetRoute(uri, fmt.Sprintf("%s:%s", rhost, config.HttpPort[config.HttpName[serviceName]]), domain, serviceName, true, direct))
	}

	return routes
}
