package initialize

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/axliupore/axapi/axapi-backend/global"
)

func InitAliyunOSS() {
	client, err := oss.New(global.Config.AliyunOSS.Endpoint, global.Config.AliyunOSS.AccessKeyId, global.Config.AliyunOSS.AccessKeySecret)
	if err != nil {
		panic(err)
	}
	buckets, err := client.ListBuckets()
	if err != nil {
		panic(err)
	}
	for _, bucket := range buckets.Buckets {
		if global.Config.AliyunOSS.BucketName == bucket.Name {
			return
		}
	}
	// 不存在这个 bucket 就创建
	if err := client.CreateBucket(global.Config.AliyunOSS.BucketName); err != nil {
		panic(err)
	}
}
