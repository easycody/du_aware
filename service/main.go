package main

import (
	"encoding/json"
	"fmt"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func main() {
	sc := []constant.ServerConfig{
		{
			IpAddr: "192.168.1.71",
			Port:   8848,
		},
	}

	cc := constant.ClientConfig{
		TimeoutMs:           5000,
		ListenInterval:      10000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}
	client, err := clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		panic(err)
	}
	// registerService(client)
	getService(client)

	// configClient, err := clients.CreateConfigClient(map[string]interface{}{
	// 	"serverConfigs": sc,
	// 	"clientConfig":  cc,
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// config(configClient)

}

func config(client config_client.IConfigClient) {
	client.PublishConfig(vo.ConfigParam{
		DataId:  "DAA-QPS",
		Group:   "FLOW-CONTROL",
		Content: "qps:1000,ts:1602648414852",
	})
}

func getService(client naming_client.INamingClient) {
	service, err := client.GetService(vo.GetServiceParam{ServiceName: "ddv", GroupName: "DDV_TEST"})
	if err != nil {
		fmt.Printf("get service error:%v", err)
	}
	data, err := json.MarshalIndent(service, "", "  ")
	if err != nil {
		fmt.Printf("marshall error:%v", err)
	}
	fmt.Printf("services:\n%s\n", string(data))
	fmt.Println("========================")
	instances, err := client.SelectInstances(vo.SelectInstancesParam{ServiceName: "ddv", HealthyOnly: true})
	if err != nil {
		fmt.Printf("get instances error:%v\n", err)
	}
	data, err = json.MarshalIndent(instances, "", " ")
	if err != nil {
		fmt.Printf("marshall error:%v\n", err)
	}
	fmt.Printf("instances:%s\n", string(data))
}

func registerService(client naming_client.INamingClient) {

	client.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "192.168.1.71",
		Port:        8848,
		ServiceName: "nacos",
		GroupName:   "DDV_TEST",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Metadata:    map[string]string{"idc": "beijing", "debug": "true"},
	})

	client.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "192.168.1.71",
		Port:        6379,
		ServiceName: "nacos",
		GroupName:   "DDV_TEST",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Metadata:    map[string]string{"idc": "beijing", "debug": "true"},
	})

	client.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "192.168.1.71",
		Port:        8801,
		ServiceName: "nacos",
		GroupName:   "DDV_TEST",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Metadata:    map[string]string{"idc": "beijing", "debug": "true"},
	})

	// client.RegisterInstance(vo.RegisterInstanceParam{
	// 	Ip:          "192.168.1.71",
	// 	Port:        9090,
	// 	ServiceName: "test",
	// 	Weight:      10,
	// 	Enable:      true,
	// 	Healthy:     true,
	// 	Metadata:    map[string]string{"idc": "beijing", "debug": "true"},
	// })

	// client.RegisterInstance(vo.RegisterInstanceParam{
	// 	Ip:          "192.168.1.71",
	// 	Port:        3306,
	// 	ServiceName: "mysql",
	// 	Weight:      10,
	// 	Enable:      true,
	// 	Healthy:     true,
	// 	Metadata:    map[string]string{"idc": "beijing", "debug": "true"},
	// })

	// client.DeregisterInstance(vo.DeregisterInstanceParam{
	// 	Ip:          "192.168.1.71",
	// 	Port:        8848,
	// 	ServiceName: "nacos",
	// })
}
