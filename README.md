# 💾 x-object-storage 对象存储
> 目前支持 ``AMS(亚马逊云存储)``

>拉取

```shell
go get -u github.com/xgd16/x-object-storage
```
```shell
go mod tidy
```

> 添加兼容项目 (需要实现以下接口函数)

```go
package main

import (
    "context"
    "os"
)

type ObjectInfoList []*ObjectInfo

type ObjectInfo struct {
    Path     string `json:"path"`
    Size     int64  `json:"size"`
    UnixTime int64  `json:"unixTime"`
    Ext      string `json:"ext"`
}

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
```