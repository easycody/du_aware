package slb

import (
	"testing"

	"github.com/nacos-group/nacos-sdk-go/util"
)

func TestDescribeBackendServers(t *testing.T) {
	accessKey := AccessKey{
		RegionId:        "cn-beijing",
		AccessKeyId:     "LTAI4GByPfZrsqkpUbkmAH69",
		AccessKeySecret: "iMpqARegSGHeprLlj2dYvPJFBTT7vh",
	}
	client, err := NewClient(accessKey)
	if err != nil {
		t.Logf("new client error:%v", err)
	}
	slbInstanceId := "lb-2zegrhxfptym9ot1ve6zo"
	bkServers, err := client.DescribeBackendServers(slbInstanceId)
	if err != nil {
		t.Logf("get slb backend server error:%v", err)
	}
	t.Logf("backendServers:%s", util.ToJsonString(bkServers))
}

func TestDescribeECS(t *testing.T) {
	accessKey := AccessKey{
		RegionId:        "cn-beijing",
		AccessKeyId:     "LTAI4GByPfZrsqkpUbkmAH69",
		AccessKeySecret: "iMpqARegSGHeprLlj2dYvPJFBTT7vh",
	}
	client, err := NewClient(accessKey)
	if err != nil {
		t.Logf("new client error:%v", err)
	}
	ecsInstanceId := "i-2ze4qda06sy5o93nh1xe"
	client.DescribeECS(ecsInstanceId)
}
