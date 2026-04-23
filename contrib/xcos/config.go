package xcos

type Config struct {
	SecretId   string `json:"secretId"`
	SecretKey  string `json:"secretKey"`
	Bucket     string `json:"bucket"`
	AppId      string `json:"appId"`
	Region     string `json:"region"`
	DriverPath string `json:"driverPath"`
}
