package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

// FSData 存储到文件系统中的数据
type FSData struct {
	IPAddr string `json:"ip_addr"`
}

// 数据存储的路径
const fdPath = "/mnt/mydevices/fsdata.json"

// LoadFSData 导入数据
func LoadFSData() (*FSData, error) {
	// 先创建目录
	err := os.MkdirAll(filepath.Dir(fdPath), 0755)
	if err != nil {
		return nil, err
	}

	// 判断数据文件是否存在，不存在，就直接实例化一个新的结构体对象
	_, err = os.Stat(fdPath)
	if err != nil {
		return new(FSData), nil
	}

	// 存在，就读取数据再反序列为对象
	bs, err := ioutil.ReadFile(fdPath)
	if err != nil {
		return nil, err
	}

	var fsData FSData
	err = json.Unmarshal(bs, &fsData)
	if err != nil {
		return nil, err
	}

	return &fsData, nil
}

// Save 保存内容到文件系统
func (f *FSData) Save() error {
	bs, _ := json.Marshal(f)
	return ioutil.WriteFile(fdPath, bs, 0644)
}
