package main

import (
	"fmt"
	"os"

	"github.com/EslRain/ftool/cmd"
)

func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
