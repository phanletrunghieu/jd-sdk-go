package request

import (
	"encoding/json"
	"fmt"
	"github.com/jarvis4901/jd-sdk-go/utils"
	"time"
)

const ServerUrl = "https://router.jd.com/api"

type JdClient struct {
	accessToken string
	appKey      string
	appSecret   string
}
type Config struct {
	Version     string `url:"v"`
	Method      string `url:"method"`
	AccessToken string `url:"access_token"`
	AppKey      string `url:"app_key"`
	SignMethod  string `url:"sign_method"`
	Format      string `url:"format"`
	Timestamp   string `url:"timestamp"`
	Sign        string `url:"sign"`
	ParamJson   string `url:"param_json"`
}

func NewClient(appKey string, appSecret string, accessToken string) *JdClient {
	return &JdClient{appKey: appKey, appSecret: appSecret, accessToken: accessToken}
}

func (c *JdClient) Execute(methodName string, reqJson map[string]interface{}) ([]byte, error) {
	paramJsonBytes, err := json.Marshal(&reqJson)
	if err != nil {
		return nil, err
	}
	paramJsonString := string(paramJsonBytes)
	fmt.Println("Bussiness params:", paramJsonString)

	// sign
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signParams := map[string]string{
		"app_key":     c.appKey,
		"format":      "json",
		"method":      methodName, //京东接口名
		"param_json":  paramJsonString,
		"sign_method": "md5",
		"timestamp":   timestamp,
		"v":           "1.0",
	}
	signValue := utils.Sign(signParams, c.appSecret)
	params := Config{
		Version:     "1.0",
		Method:      methodName,
		AccessToken: c.accessToken,
		AppKey:      c.appKey,
		SignMethod:  "md5",
		Format:      "json",
		Timestamp:   timestamp,
		Sign:        signValue,
		ParamJson:   paramJsonString,
	}
	respBytes, err := utils.HttpGet(ServerUrl, params)
	if err != nil {
		return nil, err
	}

	return []byte(respBytes), nil
}
