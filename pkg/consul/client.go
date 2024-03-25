package consul

import (
	"bytes"
	"fmt"
	"github.com/cossim/coss-cli/utils/file"
	"io/ioutil"
	"net/http"
	"strings"
)

type ConsulClient struct {
	host      string
	path      string
	namespace string
	token     string
}

var opt = map[string]string{
	"key": "/kv",
	"acl": "/acl",
}

func NewConsulClient(host string, namespace string, path string, token string) *ConsulClient {
	host = host + "/v1"
	return &ConsulClient{host: host, namespace: namespace, path: path, token: token}
}

func (c *ConsulClient) PutConfig() error {
	files, err := file.FindYamlFiles(c.path)
	if err != nil {
		return err
	}
	for _, i2 := range files {
		key := i2.Name[:strings.LastIndex(i2.Name, ".")]
		fileContent, err := ioutil.ReadFile(i2.Path)
		if err != nil {
			return err
		}

		// 构建PUT请求
		url := c.host + opt["key"] + "/" + c.namespace + "/" + key
		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(fileContent))
		if err != nil {
			return err
		}

		// 添加X-Consul-Token请求头
		req.Header.Set("X-Consul-Token", c.token)

		// 发送PUT请求
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// 处理响应
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		if resp.Status == "200 OK" {
			if string(body) == "true" {
				fmt.Printf("Success put %s to %s\n", i2.Path, c.namespace)
			}
		} else {
			return fmt.Errorf("Error sending request: %s", string(body))
		}
	}
	return nil
}

func (c *ConsulClient) SetPath(path string) {
	c.path = path
}
