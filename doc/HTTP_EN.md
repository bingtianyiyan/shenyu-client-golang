
## The Http type Register

**1.Fist make sure The ShenYuAdmin is Started, and ShenYuAdmin service active port is 9095.**
```go
Or you will see this error :

2022-05-05 15:24:28 [WARN] [github.com/apache/shenyu-client-golang/example/http_client/main.go:53] MetaDataRegister has error: The errCode is ->:503, The errMsg is  ->:Please check ShenYu admin service status

caused by:
Post "http://127.0.0.1:9095/shenyu-client/register-metadata": dial tcp 127.0.0.1:9095: connect: connection refused
2022-05-05 15:24:28 [INFO] [github.com/apache/shenyu-client-golang/example/http_client/main.go:55] finish register metadata ,the result is-> false
2022-05-05 15:24:28 [WARN] [github.com/apache/shenyu-client-golang/example/http_client/main.go:68] UrlRegister has error: The errCode is ->:503, The errMsg is  ->:Please check ShenYu admin service status

caused by:
Post "http://127.0.0.1:9095/shenyu-client/register-uri": dial tcp 127.0.0.1:9095: connect: connection refused
2022-05-05 15:24:28 [INFO] [github.com/apache/shenyu-client-golang/example/http_client/main.go:70] finish UrlRegister ,the result is-> false

```

**2.Step 1 Get shenyu_http_client. (Register service need this)**

```go
//init ShenYuHttpClient
var serverList = []string{
"http://127.0.0.1:9095",
}
acp := &http_client.HttpClientParams{
ServerList: serverList,
UserName: "admin",  //require user provide
Password: "123456", //require user provide
}


sdkClient := shenyu_sdk_client.GetFactoryClient(constants.RPCTYPE_HTTP)
client, createResult, err := sdkClient.NewClient(acp)
hcc := client.(*http_client.ShenYuHttpClient)
```


**3.Step 2 Register MetaData to ShenYu GateWay. (Need step 1 token to invoke)**
```go
//MetaDataRegister(Need Step 1 toekn adminToken.AdminTokenData)
metaData := &model.MetaDataRegister{
		AppName: "testMetaDataRegister", //require user provide
		Path:    "/your/path",           //require user provide
		Enabled: true,                   //require user provide
		Host:    "127.0.0.1",            //require user provide
		Port:    "8080",                 //require user provide
	}
    result, err := hcc.PersistInterface(metaData)
    if err != nil {
    logger.Warnf("MetaDataRegister has error:", err)
    }
    logger.Infof("finish register metadata ,the result is->", result)
	
When Register success , you will see this :  
finish register metadata ,the result is-> true
```

**4.Step 3  Url  Register  to ShenYu GateWay. **
```go
//URIRegister
//init urlRegister
	urlRegister := &model.URIRegister{
		Protocol:    "testMetaDataRegister", //require user provide
		AppName:     "testURLRegister",      //require user provide
		ContextPath: "contextPath",          //require user provide
		RPCType:     constants.RPCTYPE_HTTP, //require user provide
		Host:        "127.0.0.1",            //require user provide
		Port:        "8080",                 //require user provide
	}
    result, err = hcc.PersistURI(urlRegister)
    if err != nil {
    logger.Warnf("UrlRegister has error:", err)
    }
    logger.Infof("finish UrlRegister ,the result is->", result)

```

## Entire Success log
```go
2022-10-31 16:43:56 [INFO] [github.com/apache/shenyu-client-golang/clients/admin_client/shenyu_admin_client.go:51] Get ShenYu Http response, body is ->{Code:200 Message:login dashboard user success AdminTokenData:{ID:1 UserName:admin Role:1 Enabled:true DateCreated:2022-05-26 02:02:52 DateUpdated:2022-05-26 02:02:52 Token:eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyTmFtZSI6ImFkbWluIiwiZXhwIjoxNjY3MjkyMTAxfQ.x3kIz7xB-AuSuCWUHqpbDrhRA_pi-tj9lco7XUgNgGU}}
2022-10-31 16:43:56 [INFO] [github.com/apache/shenyu-client-golang/example/http_client/main.go:40] Create customer http client success!
2022-10-31 16:43:57 [INFO] [github.com/apache/shenyu-client-golang/example/http_client/main.go:55] finish register metadata ,the result is-> true
2022-10-31 16:43:57 [INFO] [github.com/apache/shenyu-client-golang/example/http_client/main.go:70] finish UrlRegister ,the result is-> true

```
