package node

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/universalmacro/common/config"
	"github.com/universalmacro/common/singleton"
	"github.com/universalmacro/core/controllers/models"
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

func (n *NodeConfigClient) GetConfig() models.NodeConfig {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	req, err := http.NewRequest("GET", n.ApiUrl+"/nodes/"+n.NodeId+"/config", nil)
	req.Header.Add("ApiKey", n.SecretKey)
	if err != nil {
		fmt.Printf("error %s", err)
		return models.NodeConfig{}
	}
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return models.NodeConfig{}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
	return models.NodeConfig{}
}

var nodeConfigClient = singleton.NewSingleton[NodeConfigClient](func() *NodeConfigClient {
	return newNodeConfigClient(config.GetString("core.apiUrl"), config.GetString("node.id"), config.GetString("node.secretKey"))
}, singleton.Lazy)

func GetNodeConfigClient() *NodeConfigClient {
	return nodeConfigClient.Get()
}
