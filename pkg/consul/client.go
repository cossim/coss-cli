package consul

import (
	"bytes"
	"coss-cli/utils/file"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type ConsulClient struct {
	host      string
	path      string
	namespace string
	token     string
	ssl       bool
}

var opt = map[string]string{
	"key": "/kv",
	"acl": "/acl",
}

func NewConsulClient(host string, namespace string, path string, token string, ssl bool) *ConsulClient {
	host = host + "/v1"
	if ssl {
		host = "https://" + host
	} else {
		host = "http://" + host
	}
	return &ConsulClient{host: host, namespace: namespace, path: path, token: token, ssl: ssl}
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

//type GetTokenResponse struct {
//	ID          string `json:"ID"`
//	AccessorID  string `json:"AccessorID"`
//	SecretID    string `json:"SecretID"`
//	Description string `json:"Description"`
//	Policies    []struct {
//		ID   string `json:"ID"`
//		Name string `json:"Name"`
//	} `json:"Policies"`
//	Local       bool      `json:"Local"`
//	CreateTime  time.Time `json:"CreateTime"`
//	Hash        string    `json:"Hash"`
//	CreateIndex int       `json:"CreateIndex"`
//	ModifyIndex int       `json:"ModifyIndex"`
//}

//func (c *ConsulClient) GetToken() (string, error) {
//	data := `
//	{
//	  "Description": "Agent token for 'node1'",
//	  "Policies": [
//		{
//		  "ID": "165d4317-e379-f732-ce70-86278c4558f7"
//		},
//		{
//		  "Name": "node-read"
//		}
//	  ],
//	  "TemplatedPolicies": [
//		{
//		  "TemplateName": "builtin/service",
//		  "TemplateVariables": {
//			"Name": "api"
//		  }
//		}
//	  ],
//	  "Local": false
//	}`
//
//	// 构建PUT请求
//	url := c.host + opt["acl"] + "/token"
//	fmt.Println(url)
//	req, err := http.NewRequest("PUT", url, bytes.NewBuffer([]byte(data)))
//	if err != nil {
//		fmt.Println("Error creating request:", err)
//		return "", err
//	}
//
//	// 发送PUT请求
//	client := http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		fmt.Println("Error sending request:", err)
//		return "", err
//	}
//	defer resp.Body.Close()
//
//	// 处理响应
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Println("Error reading response:", err)
//		return "", err
//	}
//	fmt.Println(string(body))
//
//	//body反序列化成对象
//	var resp1 *GetTokenResponse
//	err = json.Unmarshal(body, resp)
//	if err != nil {
//		fmt.Println("解析报错")
//		return "", err
//	}
//
//	fmt.Println(resp1.SecretID)
//	c.token = resp1.SecretID
//	return resp1.SecretID, nil
//}
