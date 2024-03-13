package file

import (
	"os"
	"path/filepath"
	"strings"
)

type ConfigInfo struct {
	Name string
	Path string
}

// 检查文件扩展名是否为yaml或yml
func isYAMLFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".yaml" || ext == ".yml"
}

func FindYamlFiles(dir string) ([]ConfigInfo, error) {
	var files []ConfigInfo
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if isYAMLFile(path) {

			files = append(files, ConfigInfo{
				Name: info.Name(),
				Path: path,
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
