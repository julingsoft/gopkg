package xoss

import (
	"context"
	"io"
	"time"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"github.com/gogf/gf/v2/frame/g"
)

type OSS struct {
	ctx    context.Context
	cfg    Config
	client *oss.Client
}

// GetInstance 返回 OSS 的单例实例
func GetInstance(ctx context.Context, config Config) *OSS {
	provider := credentials.NewStaticCredentialsProvider(config.AccessKeyID, config.AccessKeySecret)
	cfg := oss.LoadDefaultConfig().
		WithCredentialsProvider(provider).
		WithRegion(config.RegionName)

	return &OSS{
		ctx:    ctx,
		cfg:    config,
		client: oss.NewClient(cfg),
	}
}

func (o *OSS) Client() *oss.Client {
	return o.client
}

func (o *OSS) PutObject(objectName string, body io.Reader) (*oss.PutObjectResult, error) {
	putRequest := &oss.PutObjectRequest{
		Bucket:       oss.Ptr(o.cfg.BucketName), // 存储空间名称
		Key:          oss.Ptr(objectName),       // 对象名称
		Body:         body,                      // 对象内容
		StorageClass: oss.StorageClassStandard,  // 指定对象的存储类型为标准存储
		Acl:          oss.ObjectACLPrivate,      // 指定对象的访问权限为私有访问
	}

	return o.Client().PutObject(o.ctx, putRequest)
}

func (o *OSS) PutObjectFromFile(objectName string, filePath string) (*oss.PutObjectResult, error) {
	putRequest := &oss.PutObjectRequest{
		Bucket:       oss.Ptr(o.cfg.BucketName), // 存储空间名称
		Key:          oss.Ptr(objectName),       // 对象名称
		StorageClass: oss.StorageClassStandard,  // 指定对象的存储类型为标准存储
		Acl:          oss.ObjectACLPrivate,      // 指定对象的访问权限为私有访问
	}

	return o.Client().PutObjectFromFile(o.ctx, putRequest, filePath)
}

func (o *OSS) GetObject(objectName string) (*oss.GetObjectResult, error) {
	getRequest := &oss.GetObjectRequest{
		Bucket: oss.Ptr(o.cfg.BucketName), // 存储空间名称
		Key:    oss.Ptr(objectName),       // 对象名称
	}

	return o.Client().GetObject(o.ctx, getRequest)
}

func (o *OSS) MustGetObject(objectName string) string {
	result, err := o.GetObject(objectName)
	if err != nil {
		g.Log().Error(o.ctx, err, "[oss] GetObject", objectName)
		return ""
	}
	defer result.Body.Close()

	data, err := io.ReadAll(result.Body)
	if err != nil {
		g.Log().Error(o.ctx, err, "[oss] ReadAll", objectName)
		return ""
	}

	return string(data)
}

func (o *OSS) GetPresign(objectName string, expires time.Duration) (*oss.PresignResult, error) {
	getRequest := &oss.GetObjectRequest{
		Bucket: oss.Ptr(o.cfg.BucketName), // 存储空间名称
		Key:    oss.Ptr(objectName),       // 对象名称
	}

	return o.Client().Presign(o.ctx, getRequest, oss.PresignExpires(expires))
}

func (o *OSS) GetSignURL(objectName string, expires time.Duration) (string, error) {
	presignResult, err := o.GetPresign(objectName, expires)
	if err != nil {
		return "", err
	}
	return presignResult.URL, nil
}
