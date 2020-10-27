package helper

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/EslRain/ftool/model"
)

func GetFiles(folder string, isRecursive bool) ([]model.File, error) {
	var files []model.File
	if isRecursive {
		// 递归子目录版
		err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
			file, err := filepath.Abs(path)
			if err != nil {
				return nil
			}

			if !info.IsDir() {
				files = append(files, model.ConstructFile(file))
			}
			return nil
		})
		return files, err
	} else {
		// 不递归子目录版
		f, err := os.Open(folder)
		defer f.Close()

		if err != nil {
			return files, err
		}

		if fileinfo, err := f.Readdir(-1); err == nil {
			for _, file := range fileinfo {
				if !file.IsDir() {
					folder, err := filepath.Abs(folder)
					if err != nil {
						return files, err
					}
					files = append(files, model.ConstructFile(folder+"\\"+file.Name()))
				}
			}
		} else {
			return files, err
		}
	}

	return files, nil
}

func GetNewFilePath(file model.File, sourceFolder, targetFolder string) (string, error) {
	if file.FullPath == "" {
		return "", fmt.Errorf("Original file is not valid.")
	}
	subFolder := GetSubFolder(file, sourceFolder)
	return targetFolder + subFolder + file.File, nil
}

func GetSubFolder(file model.File, sourceFolder string) string {
	return strings.TrimPrefix(file.Folder, sourceFolder)
}
