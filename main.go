package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Initialization Starting....")

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Can not get working directory!")
		return
	}

	isEmpty, err := IsEmpty(cwd)
	if err != nil {
		fmt.Println("Somethig wrong")
		return
	}
	if !isEmpty {
		fmt.Println("Current directory is not empty")
		return
	}

	cmd := cwd + "/cmd"
	os.Mkdir(cmd, os.ModePerm)
	os.Mkdir(cwd+"/internal", os.ModePerm)
	os.Mkdir(cwd+"/internal/app", os.ModePerm)
	os.Mkdir(cwd+"/internal/pkg", os.ModePerm)
	os.Mkdir(cwd+"/init", os.ModePerm)

	mainPath := cmd + "/main.go"
	f, err := os.Create(mainPath)
	if err == nil {
		defer f.Close()
	} else {
		println("Create main.go failed")
		return
	}

	f.WriteString(mainTemplate)
	println("Initialization completed")
}

// 判断文件夹是否为空
func IsEmpty(dir string) (bool, error) {
	f, err := os.Open(dir)
	if err != nil {
		return false, err
	}
	defer f.Close()

	f.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err
}
