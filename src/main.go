package main

import (
	"cheese.self/go-jvm/src/classpath"
	"fmt"
	"strings"
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
	currentClasspath := classpath.Parse(cmd.XJreOption, cmd.cpOption)

	fmt.Printf("classpath:%s class:%s args:%v\n",
		currentClasspath, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := currentClasspath.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
	}

	fmt.Printf("Class Data: %v\n", classData)
}