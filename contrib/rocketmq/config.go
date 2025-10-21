package rocketmq

type Config struct {
	Endpoint      string `json:"endpoint" validate:"required"`
	NameSpace     string `json:"nameSpace"`
	ConsumerGroup string `json:"consumerGroup"`
	AccessKey     string `json:"accessKey" validate:"required"`
	AccessSecret  string `json:"accessSecret" validate:"required"`
	SecurityToken string `json:"securityToken"`
}
