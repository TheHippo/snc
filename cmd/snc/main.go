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

func init() {
	flag.VarP(&listen, "listen", "l", "port to listen")
	flag.Parse()
}

func main() {
	fmt.Println(listen)
	// if listen.set == true {
	// 	// open the server and listen
	// 	fmt.Println("Starting server")
	// } else {
	// 	fmt.Println("Starting client")
	// 	// connect to server
	// }

}
