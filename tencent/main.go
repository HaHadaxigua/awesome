package main

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

func main() {
	provider := common.DefaultCvmRoleProvider()
	credential, err := provider.GetCredential()
	if err != nil {
		logrus.Errorf("get credential failed")
		panic(err)
	}
	logrus.Infof("id: %s", credential.GetSecretId())
	logrus.Infof("key: %s", credential.GetSecretKey())
	logrus.Infof("token: %s", credential.GetToken())

	logrus.Info("get credential succeed")

	client, err := cvm.NewClient(credential, regions.Shanghai, profile.NewClientProfile())
	if err != nil {
		logrus.Errorf("make client failed")
		panic(err)
	}

	request := cvm.NewDescribeInstancesRequest()
	response, err := client.DescribeInstances(request)

	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", response.ToJsonString())
}
