package utils

import (
	"os"
	"path/filepath"
)

const gorchFilePat = "*.gorch"

// GlobGorchFiles 扫描出目录里的所有jgf文件
func GlobGorchFiles(dir string) ([]string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var allfile []string
	for _, v := range files {
		path := v.Name()
		if v.IsDir() {
			a, err := GlobGorchFiles(path)
			if err != nil {
				return nil, err
			}
			allfile = append(allfile, a...)
		} else {
			matched, _ := filepath.Match(gorchFilePat, v.Name())
			if matched {
				allfile = append(allfile, path)
			}
		}
	}

	return allfile, nil
}
