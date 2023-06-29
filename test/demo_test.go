package test

import (
	"fmt"
	"github.com/xgd16/x-object-storage/disk"
	"github.com/xgd16/x-object-storage/drive"
	"os"
	"testing"
)

func TestDisk(t *testing.T) {
	//file, err := os.Open("./file/9eae00cd14d9aa75b84acf5210a532d6.jpg")
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
	//g.DumpWithType(diskObj.DelObject("test/aa.jpg"))
}
