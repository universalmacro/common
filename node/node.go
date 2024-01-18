package node

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/universalmacro/common/config"
	"github.com/universalmacro/common/singleton"
)

func newNodeConfigClient(apiUrl, nodeId, secretKey string) *NodeConfigClient {
	return &NodeConfigClient{
		ApiUrl:    apiUrl,
		NodeId:    nodeId,
		SecretKey: secretKey,
	}
}

type NodeConfigClient struct {
	ApiUrl    string
	NodeId    string
	SecretKey string
}

type NodeConfig struct {
	NodeID    uint
	SecretKey string        `gorm:"type:varchar(64)"`
	Api       *ApiConfig    `gorm:"type:json"`
	Server    *ServerConfig `gorm:"type:json"`
	Database  *DBConfig     `gorm:"type:json"`
	Redis     *RedisConfig  `gorm:"type:json"`
}
type ApiConfig struct {
	MerchantUrl string `json:"merchantUrl" gorm:"type:varchar(256)"`
}
type ServerConfig struct {
	Port      string `json:"port" gorm:"type:varchar(64)"`
	JwtSecret string `json:"jwtSecret" gorm:"type:varchar(64)"`
}
type DBConfig struct {
	Host     string `json:"host" gorm:"type:varchar(64)"`
	Port     string `json:"port" gorm:"type:varchar(64)"`
	Username string `json:"username" gorm:"type:varchar(64)"`
	Password string `json:"password" gorm:"type:varchar(64)"`
	Database string `json:"database" gorm:"type:varchar(64)"`
}
type RedisConfig struct {
	Host     string `json:"host" gorm:"type:varchar(64)"`
	Port     string `json:"port" gorm:"type:varchar(64)"`
	Password string `json:"password" gorm:"type:varchar(64)"`
}

func (n *NodeConfigClient) GetConfig() *NodeConfig {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	req, err := http.NewRequest("GET", n.ApiUrl+"/nodes/"+n.NodeId+"/config", nil)
	req.Header.Add("ApiKey", n.SecretKey)
	if err != nil {
		fmt.Printf("error %s", err)
		return nil
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic("id or secret key is wrong")
	}
	body, err := ioutil.ReadAll(resp.Body)
	var config NodeConfig
	json.Unmarshal(body, &config)
	return &config
}

var nodeConfigClient = singleton.NewSingleton[NodeConfigClient](func() *NodeConfigClient {
	return newNodeConfigClient(config.GetString("core.apiUrl"), config.GetString("node.id"), config.GetString("node.secretKey"))
}, singleton.Lazy)

func GetNodeConfigClient() *NodeConfigClient {
	return nodeConfigClient.Get()
}
