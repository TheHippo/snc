package main

import (
	"fmt"

	"github.com/TheHippo/snc/utils"
	flag "github.com/ogier/pflag"
)

// injected when building
var version = "undefined"
var date = "undefined"
var goVersion = "undefined"

var listen utils.OptionalIntValue

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
