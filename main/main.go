package main

import (
	"fmt"
	"os"
	"github.com/scream870102/argpas"
)

//this is for test
func main() {
	args := os.Args
	res, err := argpas.ParseArgToBool(args[1], "Test")
	if err != nil {
		fmt.Println(err.Error())
	}else{
		fmt.Println(res)
	}

}
