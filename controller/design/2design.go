package design

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)


type Header struct {
	AppId string
	Sign  string
}
type Data struct {
	Header Header
	Body   string
}

type AuthToken struct {
	sign string
}

func CreateAuthToken() *AuthToken {
	return &AuthToken{}
}

func (authToken *AuthToken) Create(appId string, appKey string, body string) string {
	h := md5.New()
	h.Write([]byte(appId + body + appKey))
	authToken.sign = strings.ToUpper(fmt.Sprintf("%x", h.Sum(nil)))
	return authToken.sign
}

func (authToken *AuthToken) Match(token *AuthToken) bool {
	if authToken.sign == token.sign {
		return true
	}
	return false
}

type ApiRequest struct {
	appId string
	sign  string
	data  *Data
}

func CreateApiRequest() *ApiRequest {
	return &ApiRequest{}
}
func (apiRequest *ApiRequest) Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}
func (apiRequest *ApiRequest) Decode(data string) (appId string, sign string, err error) {
	bytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return
	}
	apiRequest.data = &Data{}
	if err := json.Unmarshal(bytes, apiRequest.data); err != nil {
		return "", "", err
	}
	apiRequest.appId = apiRequest.data.Header.AppId
	apiRequest.sign = apiRequest.data.Header.Sign
	return apiRequest.appId, apiRequest.sign, nil
}

func (apiRequest *ApiRequest) GetAppid() string {
	return apiRequest.appId
}

func (apiRequest *ApiRequest) GetSign() string {
	return apiRequest.sign
}

type CredentialStorage interface {
	GetAppkeyByAppid(appId string) string
}

type CredentialStorageConfig struct {
}

func (config *CredentialStorageConfig) GetAppkeyByAppid(appId string) string {
	if appId == "test" {
		return "test"
	}
	return "test"
}

type ApiAuthencator struct {
	credentialStorage CredentialStorage
}

func CreateApiAuthenCator(cs CredentialStorage) *ApiAuthencator {
	return &ApiAuthencator{credentialStorage: cs}
}

func (apiAuthencator *ApiAuthencator) Auth(data string) (bool, error) {
	//1.解析数据
	apiRequest := CreateApiRequest()
	appId, sign, err := apiRequest.Decode(data)
	//fmt.Println(appId, sign, apiRequest.data)
	if err != nil {
		return false, fmt.Errorf("Decode failed")
	}
	//2.获取appId对应的appkey
	appKey := apiAuthencator.credentialStorage.GetAppkeyByAppid(appId)
	//3.重新计算sign
	authToken := CreateAuthToken()
	newSign := authToken.Create(appId, appKey, apiRequest.data.Body)
	if sign == newSign {
		return true, nil
	}
	return false, nil
}
func maindesign() {
	//客户端
	appId := "test"
	appKey := "test"
	sendData := &Data{
		Header: Header{
			AppId: appId,
		},
		Body: "for test",
	}
	authToken := CreateAuthToken()
	sign := authToken.Create(appId, appKey, sendData.Body)
	sendData.Header.Sign = sign
	sendDataMarshal, _ := json.Marshal(sendData)
	sendDataString := CreateApiRequest().Encode(string(sendDataMarshal))
	//fmt.Println(sign, sendData, string(sendDataMarshal), string(sendDataString))

	//服务端
	apiAuthenCator := CreateApiAuthenCator(new(CredentialStorageConfig))
	auth, err := apiAuthenCator.Auth(sendDataString)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if auth == false {
		fmt.Println("auth failed")
		return
	}
	fmt.Println("auth success")
	return
}

