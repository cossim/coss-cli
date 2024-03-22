package apisix

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"sync"
)

var ConsulRoutes = []string{
	`{"uri": "/api/v1/user/*", "upstream": {"service_name": "user_bff", "type": "roundrobin", "discovery_type": "consul"}, "plugins": {"limit-count": {"count": 5, "time_window": 10, "rejected_code": 503, "_meta": {"disable": true}}, "cors": {}}}`,
	`{"uri": "/api/v1/group/*", "upstream": {"service_name": "group_bff", "type": "roundrobin", "discovery_type": "consul"}, "plugins": {"limit-count": {"count": 5, "time_window": 10, "rejected_code": 503, "_meta": {"disable": true}}, "cors": {}}}`,
	`{"uri": "/api/v1/relation/*", "upstream": {"service_name": "relation_bff", "type": "roundrobin", "discovery_type": "consul"}, "plugins": {"limit-count": {"count": 5, "time_window": 10, "rejected_code": 503, "_meta": {"disable": true}}, "cors": {}}}`,
	`{"uri": "/api/v1/admin/*", "upstream": {"service_name": "admin_bff", "type": "roundrobin", "discovery_type": "consul"}, "plugins": {"limit-count": {"count": 5, "time_window": 10, "rejected_code": 503, "_meta": {"disable": true}}, "cors": {}}}`,
	`{"uri": "/api/v1/msg/*", "upstream": {"service_name": "msg_bff", "type": "roundrobin", "discovery_type": "consul"}, "plugins": {"limit-count": {"count": 5, "time_window": 10, "rejected_code": 503, "_meta": {"disable": true}}, "cors": {}}}`,
	`{"uri": "/api/v1/storage/*", "upstream": {"service_name": "storage_bff", "type": "roundrobin", "discovery_type": "consul"}, "plugins": {"limit-count": {"count": 5, "time_window": 10, "rejected_code": 503, "_meta": {"disable": true}}, "cors": {}}}`,
	`{"uri": "/api/v1/msg/ws", "name": "ws", "enable_websocket": true, "upstream": {"service_name": "msg_bff", "type": "roundrobin", "discovery_type": "consul"}, "plugins": {"limit-count": {"count": 5, "time_window": 10, "rejected_code": 503, "_meta": {"disable": true}}, "cors": {}}}`,
}

var Routes = []string{
	`{"uri": "/api/v1/user/*", "upstream": {"type": "roundrobin", "nodes": {"user:8083": 1}}, "plugins": {"limit-count": {"count": 5, "time_window": 10, "rejected_code": 503, "_meta": {"disable": true}}, "cors": {}}}`,
	`{"uri": "/api/v1/group/*", "upstream": {"type": "roundrobin", "nodes": {"group:8084": 1}}, "plugins": {"limit-count": {"count": 5, "time_window": 10, "rejected_code": 503, "_meta": {"disable": true}}, "cors": {}}}`,
	`{"uri": "/api/v1/relation/*", "upstream": {"type": "roundrobin", "nodes": {"relation:8082": 1}}, "plugins": {"limit-count": {"count": 5, "time_window": 10, "rejected_code": 503, "_meta": {"disable": true}}, "cors": {}}}`,
	`{"uri": "/api/v1/admin/*", "upstream": {"type": "roundrobin", "nodes": {"admin:8087": 1}}, "plugins": {"limit-count": {"count": 5, "time_window": 10, "rejected_code": 503, "_meta": {"disable": true}}, "cors": {}}}`,
	`{"uri": "/api/v1/msg/*", "upstream": {"type": "roundrobin", "nodes": {"msg:8081": 1}}, "plugins": {"limit-count": {"count": 5, "time_window": 10, "rejected_code": 503, "_meta": {"disable": true}}, "cors": {}}}`,
	`{"uri": "/api/v1/storage/*", "upstream": {"type": "roundrobin", "nodes": {"storage:8085": 1}}, "plugins": {"limit-count": {"count": 5, "time_window": 10, "rejected_code": 503, "_meta": {"disable": true}}, "cors": {}}}`,
	`{"uri": "/api/v1/msg/ws", "name": "ws", "enable_websocket": true, "upstream": {"type": "roundrobin", "nodes": {"msg:8081": 1}}, "plugins": {"limit-count": {"count": 5, "time_window": 10, "rejected_code": 503, "_meta": {"disable": true}}, "cors": {}}}`,
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
