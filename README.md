# 💾 x-object-storage 对象存储
<font size="2" color=#006666>西安豆芽科技有限公司 **制**</font>


> 目前支持 ``AMS(亚马逊云存储)``

> 拉取

```shell
go get -u github.com/xgd16/x-object-storage
```
```shell
go mod tidy
```

> 使用演示

```go
package main

import (
    "os"
    "fmt"
    "github.com/xgd16/x-object-storage/disk"
    "github.com/xgd16/x-object-storage/drive"
)

func main() {
    file, err := os.Open("./file/PROJECT_README.pdf")
    if err != nil {
        panic("读取文件错误" + err.Error())
    }
    defer func() { _ = file.Close() }()

    diskObj, err := disk.New(&drive.AmsDrive{
        Region:    "***",
        SecretId:  "***",
        SecretKey: "***",
        Bucket:    "***",
    })
    if err != nil {
        panic("初始化 对象失败" + err.Error())
    }

    //fmt.Println(diskObj.PutObject(file, "test/PROJECT_README.pdf"))
    fileList, err := diskObj.GetPathList()
    if err != nil {
        panic("获取列表失败")
    }
    for _, item := range fileList {
        fmt.Println(item)
        fmt.Println(diskObj.GetObjectUrl(item.Path))
    }
}
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