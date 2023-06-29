package types

import (
	"context"
	"os"
)

type ObjectStorage interface {
	// Init 初始化函数
	Init(ctx context.Context) (ObjectStorage, error)
	// PutObject 上传对象文件
	PutObject(file *os.File, filePath string) error
	// GetPathList 获取路径列表
	GetPathList() (ObjectInfoList, error)
	// DelObject 删除对象
	DelObject(path string) error
	// GetObjectUrl 获取对象url地址
	GetObjectUrl(path string, opt ...any) (string, error)
}

type ObjectInfoList []*ObjectInfo

type ObjectInfo struct {
	Path     string `json:"path"`
	Size     int64  `json:"size"`
	UnixTime int64  `json:"unixTime"`
	Ext      string `json:"ext"`
}
