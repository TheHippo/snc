package main

import (
	"fmt"

	"github.com/TheHippo/snc"
	flag "github.com/ogier/pflag"
)

// injected when building
var (
	version   = "undefined"
	date      = "undefined"
	goVersion = "undefined"
)

var listen snc.OptionalIntValue
var showVersion bool

func init() {
	flag.BoolVarP(&showVersion, "version", "v", false, "display snc version")
	flag.VarP(&listen, "listen", "l", "port to listen")
	flag.Parse()
}

func main() {
	if showVersion == true {
		fmt.Printf("Version: %s\n", version)
		fmt.Printf("Build date: %s\n", date)
		fmt.Printf("Go version: %s\n", goVersion)
		return
	}
	fmt.Println(listen)
	// if listen.set == true {
	// 	// open the server and listen
	// 	fmt.Println("Starting server")
	// } else {
	// 	fmt.Println("Starting client")
	// 	// connect to server
	// }

}
