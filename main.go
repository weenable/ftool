package main

import (
	"file-manage/cmd"
	"fmt"
	"os"
)

func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
