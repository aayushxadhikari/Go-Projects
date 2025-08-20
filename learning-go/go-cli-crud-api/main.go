package main

import (
	"flag"
	"fmt"
	"os"
)

func main(){
	mode := flag.String("mode", "server","Mode: server or cli")
	flag.Parse()

	if *mode == "sever"{
		startServer()
	}else if *mode == "cli" {
		runCLI(os.Args[2:])
	}else{
		fmt.Println("Unknown mode. Use --mode=server or --mode=cli")
	}
}

