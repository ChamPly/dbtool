package export

import (
	"fmt"
	"os"
	"path/filepath"
)

type FileExport struct {
}

func NewFileExport() *FileExport {
	return &FileExport{}
}

func (fe *FileExport) Do(sqlContent map[string]string, conf Conf) (err error) {
	path, err := filepath.Abs(conf.FileName)
	if err != nil {
		err = fmt.Errorf("文件的路径不对:fileName:%s", conf.FileName)
		return
	}
	dir := filepath.Dir(path)
	err = os.MkdirAll(dir, 0777)
	if err != nil {
		err = fmt.Errorf("创建文件夹：%s失败:%+v", path, err)
		return
	}

	file, err := os.OpenFile(conf.FileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		err = fmt.Errorf("打开文件:%s 失败：%+v", conf.FileName, err)
		return
	}
	defer file.Close()

	for _, sql := range sqlContent {
		file.WriteString(sql)
	}
	return
}
