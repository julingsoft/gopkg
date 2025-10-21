package xoss

import (
	"context"
	"io"
	"sync"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"github.com/gogf/gf/v2/frame/g"
)

type OSS struct {
	Ctx    context.Context
	Config Config
	Client *oss.Client
}

var (
	instance *OSS
	once     sync.Once
)

// GetInstance 返回 OSS 的单例实例
func GetInstance(ctx context.Context, config Config) *OSS {
	once.Do(func() {
		provider := credentials.NewStaticCredentialsProvider(config.AccessKeyID, config.AccessKeySecret)
		cfg := oss.LoadDefaultConfig().
			WithCredentialsProvider(provider).
			WithRegion(config.RegionName)

		instance = &OSS{
			Ctx:    ctx,
			Config: config,
			Client: oss.NewClient(cfg),
		}
	})

	return instance
}

func (o *OSS) PutObject(objectName string, body io.Reader) (*oss.PutObjectResult, error) {
	putRequest := &oss.PutObjectRequest{
		Bucket:       oss.Ptr(o.Config.BucketName), // 存储空间名称
		Key:          oss.Ptr(objectName),          // 对象名称
		Body:         body,                         // 对象内容
		StorageClass: oss.StorageClassStandard,     // 指定对象的存储类型为标准存储
		Acl:          oss.ObjectACLPrivate,         // 指定对象的访问权限为私有访问
	}

	return o.Client.PutObject(o.Ctx, putRequest)
}

func (o *OSS) PutObjectFromFile(objectName string, filePath string) (*oss.PutObjectResult, error) {
	putRequest := &oss.PutObjectRequest{
		Bucket:       oss.Ptr(o.Config.BucketName), // 存储空间名称
		Key:          oss.Ptr(objectName),          // 对象名称
		StorageClass: oss.StorageClassStandard,     // 指定对象的存储类型为标准存储
		Acl:          oss.ObjectACLPrivate,         // 指定对象的访问权限为私有访问
	}

	return o.Client.PutObjectFromFile(o.Ctx, putRequest, filePath)
}

func (o *OSS) GetObject(ctx context.Context, objectName string) (*oss.GetObjectResult, error) {
	getRequest := &oss.GetObjectRequest{
		Bucket: oss.Ptr(o.Config.BucketName), // 存储空间名称
		Key:    oss.Ptr(objectName),          // 对象名称
	}

	return o.Client.GetObject(ctx, getRequest)
}

func (o *OSS) MustGetObject(objectName string) string {
	result, err := o.GetObject(o.Ctx, objectName)
	if err != nil {
		g.Log().Error(o.Ctx, err, "[oss] GetObject", objectName)
		return ""
	}
	defer result.Body.Close()

	data, err := io.ReadAll(result.Body)
	if err != nil {
		g.Log().Error(o.Ctx, err, "[oss] ReadAll", objectName)
		return ""
	}

	return string(data)
}
