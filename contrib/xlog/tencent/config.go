package tencent

type Config struct {
	EndPoint  string `json:"endpoint" dc:"日志端点"`
	SecretId  string `json:"secretId" dc:"AccessKeyID"`
	SecretKey string `json:"secretKey" dc:"AccessKeySecret"`
	TopicID   string `json:"topicID" dc:"日志主题ID"`
}
