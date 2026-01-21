package rocketmq

type Config struct {
	Endpoint      string `json:"endpoint"`
	NameSpace     string `json:"nameSpace"`
	ConsumerGroup string `json:"consumerGroup"`
	Topic         string `json:"topic"`
	AccessKey     string `json:"accessKey"`
	SecretKey     string `json:"secretKey"`
	SecurityToken string `json:"securityToken"`
}
