package main

import (
	"fmt"
	"flag"
	"optiopay/api/cmd"
)

func main() {
	cmd_flag := flag.Bool("cmd", false, "Interactive commandline prompt")
	api_flag := flag.Bool("api", false, "Web API - not implemented")
	file_flag := flag.Bool("file", false, "File input - not implemented")

	flag.Parse()

	if *api_flag == true || *file_flag == true || *cmd_flag == false {
		fmt.Println("Only interactive commandline promt is implemented. Please use -cmd flag")
	} else {
		cmd.Init()
	}
}
