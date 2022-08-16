package main

import (
	"encoding/json"
	"github.com/apache/shenyu-client-golang/clients/etcd_client"
	"github.com/apache/shenyu-client-golang/common/constants"
	"github.com/apache/shenyu-client-golang/common/shenyu_sdk_client"
	"github.com/apache/shenyu-client-golang/model"
	"github.com/wonderivan/logger"
	"time"
)

func main(){
	ccp := &etcd_client.EtcdClientParam{
		EtcdServers: []string{"http://127.0.0.1:2379"}, //require user provide
		TTL:    50,
		TimeOut: 1000,
	}

	sdkClient := shenyu_sdk_client.GetFactoryClient(constants.ETCD_CLIENT)
	client, createResult, err := sdkClient.NewClient(ccp)
	if !createResult && err != nil {
		logger.Fatal("Create ShenYuEtcdClient error : %+V", err)
	}

	etcd := client.(*etcd_client.ShenYuEtcdClient)
	defer etcd.Close()

	//init MetaDataRegister
	metaData1 := &model.MetaDataRegister{
		AppName: "testMetaDataRegister1", //require user provide
		Path:    "your/path1",            //require user provide
		Enabled: true,                    //require user provide
		Host:    "127.0.0.1",             //require user provide
		Port:    "8080",                  //require user provide
	}

	metaData2 := &model.MetaDataRegister{
		AppName: "testMetaDataRegister2", //require user provide
		Path:    "your/path2",            //require user provide
		Enabled: true,                    //require user provide
		Host:    "127.0.0.1",             //require user provide
		Port:    "8181",                  //require user provide
	}


	//register multiple metaData
	registerResult1, err := etcd.RegisterServiceInstance(metaData1)
	if !registerResult1 && err != nil {
		logger.Fatal("Register etcd Instance error : %+V", err)
	}

	registerResult2, err := etcd.RegisterServiceInstance(metaData2)
	if !registerResult2 && err != nil {
		logger.Fatal("Register etcd Instance error : %+V", err)
	}


	time.Sleep(time.Second)

	instanceDetail, err := etcd.GetServiceInstanceInfo(metaData1)
	nodes1, ok := instanceDetail.([]*model.MetaDataRegister)
	if !ok {
		logger.Fatal("get etcd client metaData error %+v:", err)
	}

	//range nodes
	for index, node := range nodes1 {
		nodeJson, err := json.Marshal(node)
		if err == nil {
			logger.Info("GetNodesInfo ,success Index", index, string(nodeJson))
		}
	}

	instanceDetail2, err := etcd.GetServiceInstanceInfo(metaData2)
	nodes2, ok := instanceDetail2.([]*model.MetaDataRegister)
	if !ok {
		logger.Fatal("get etcd client metaData error %+v:", err)
	}

	//range nodes
	for index, node := range nodes2 {
		nodeJson, err := json.Marshal(node)
		if err == nil {
			logger.Info("GetNodesInfo ,success Index", index, string(nodeJson))
		}
	}

	logger.Info("> DeregisterServiceInstance start")
	deRegisterResult1, err := etcd.DeregisterServiceInstance(metaData1)
	if err != nil {
		panic(err)
	}

	deRegisterResult2, err := etcd.DeregisterServiceInstance(metaData2)
	if err != nil {
		panic(err)
	}

	if deRegisterResult1 && deRegisterResult2 {
		logger.Info("DeregisterServiceInstance success !")
	}
	//DeregisterServiceInstance end

	//do your logic
}