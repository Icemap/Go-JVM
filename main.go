package main

import (
	"fmt"
)

func main() {
	cmd := ParseCmd()
	if cmd.versionFlag {
		fmt.Println("Golang JVM By Cheese v0.0.1")
	} else if cmd.helpFlag || cmd.class == ""  {
		PrintUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd)  {
	fmt.Printf("classpath:%s class:%s args:%v\n",
		cmd.cpOption, cmd.class, cmd.args)
}