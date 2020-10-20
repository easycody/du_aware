package slb

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/slb"
)

type AccessKey struct {
	RegionId        string `yaml:"regionId" json:"regionId"`
	AccessKeyId     string `yaml:"accessKeyId" json:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret" json:"accessKeySecret"`
}

type Client struct {
	slbClient *slb.Client
	ecsClient *ecs.Client
}

type BackendServer struct {
	ServerId string `json:"ServerId" xml:"ServerId"`
	Weight   int    `json:"Weight" xml:"Weight"`
	ServerIp string `json:"ServerIp" xml:"ServerIp"`
}

func NewClient(access AccessKey) (*Client, error) {
	slibClient, err := slb.NewClientWithAccessKey(access.RegionId, access.AccessKeyId, access.AccessKeySecret)
	if err != nil {
		return nil, err
	}
	ecsClient, err := ecs.NewClientWithAccessKey(access.RegionId, access.AccessKeyId, access.AccessKeySecret)
	if err != nil {
		return nil, err
	}
	client := new(Client)
	client.slbClient = slibClient
	client.ecsClient = ecsClient
	return client, nil
}

func (c *Client) DescribeBackendServers(slbInstance string) ([]BackendServer, error) {
	slbClient := c.slbClient
	request := slb.CreateDescribeLoadBalancerAttributeRequest()
	request.Scheme = "https"
	request.LoadBalancerId = slbInstance
	response, err := slbClient.DescribeLoadBalancerAttribute(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	var backendServers []BackendServer
	servers := response.BackendServers
	for _, server := range servers.BackendServer {
		if ip, err := c.DescribeECS(server.ServerId); err == nil {
			backendServer := BackendServer{
				ServerId: server.ServerId,
				Weight:   server.Weight,
				ServerIp: ip,
			}
			backendServers = append(backendServers, backendServer)
		}
	}
	return backendServers, nil
}

func (c *Client) DescribeECS(ecsInstance string) (string, error) {
	client := c.ecsClient
	request := ecs.CreateDescribeInstanceAttributeRequest()
	request.Scheme = "https"
	request.InstanceId = ecsInstance
	response, err := client.DescribeInstanceAttribute(request)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	privateAddress := response.VpcAttributes.PrivateIpAddress.IpAddress
	if len(privateAddress) > 0 {
		return privateAddress[0], nil
	}
	return "", fmt.Errorf("not found private ip address for ecs:%s", ecsInstance)
}
