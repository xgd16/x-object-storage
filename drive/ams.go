package drive

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/xgd16/x-object-storage/types"
	"os"
	"path/filepath"
	"strings"
)

type AmsDrive struct {
	Region    string // 区域
	SecretId  string
	SecretKey string
	Bucket    string
	client    *s3.Client
	ctx       context.Context
}

func (t *AmsDrive) Init(ctx context.Context) (types.ObjectStorage, error) {
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(t.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			t.SecretId,
			t.SecretKey,
			"", //此参数为token，若自己的凭据无token则设定为空
		)),
	)
	if err != nil {
		return nil, err
	}
	t.ctx = ctx
	t.client = s3.NewFromConfig(cfg)
	return t, nil
}

func (t *AmsDrive) PutObject(file *os.File, filePath string) error {
	_, err := t.client.PutObject(t.ctx, &s3.PutObjectInput{
		Bucket: aws.String(t.Bucket),
		Key:    aws.String(filePath),
		Body:   file,
	})
	return err
}

func (t *AmsDrive) GetPathList() (types.ObjectInfoList, error) {
	resp, err := t.client.ListObjectsV2(t.ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(t.Bucket),
	})
	if err != nil {
		return nil, err
	}
	objList := make(types.ObjectInfoList, 0)
	for _, obj := range resp.Contents {
		objList = append(objList, &types.ObjectInfo{
			Path:     *obj.Key,
			Size:     obj.Size,
			UnixTime: obj.LastModified.Unix(),
			Ext:      strings.TrimPrefix(filepath.Ext(*obj.Key), "."),
		})
	}
	return objList, nil
}

func (t *AmsDrive) DelObject(path string) error {
	_, err := t.client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: &t.Bucket,
		Key:    &path,
	})
	return err
}

func (t *AmsDrive) GetObjectUrl(path string, opt ...any) (string, error) {
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", t.Bucket, t.Region, path), nil
}
