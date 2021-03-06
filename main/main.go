package main

import (
	"fmt"
	"github.com/scream870102/argpas"
	"os"
)

type arguments struct {
	StringArg  string
	BoolArg    bool
	IntArg     int
	FloatArg   float64
	TriggerArg argpas.Trigger
}

func main() {
	interfaceTest()
	stringTest()
	argTest()
}

func interfaceTest() {
	var momoko arguments
	args := os.Args[1:]
	if err := argpas.ParseToInterface(&momoko, args); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(momoko)
}

func stringTest(){
	var yuuka string
	arg := os.Args[1]
	if result,err:= argpas.ParseToStr(arg,"StringArg");err!=nil{
		fmt.Println(err.Error())
	}else{
		yuuka = result
	}
	fmt.Println(yuuka)
}

func argTest(){
	arg:= os.Args[1]
	if result,err:= argpas.Parse(arg,"haruka");err!=nil{
		fmt.Println(err.Error())
	}else{
		fmt.Println(result)
	}
}
