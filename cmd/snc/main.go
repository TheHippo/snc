package main

import (
	"fmt"
	"os"

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
var bind string
var showVersion bool

func init() {
	flag.BoolVarP(&showVersion, "version", "v", false, "display snc version")
	flag.StringVarP(&bind, "bind", "b", "0.0.0.0", "ip adress to bind")
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
	if listen.Parsed == true {
		// open the server and listen
		fmt.Println("Starting server")
		server := snc.NewServer(bind, int(listen.Value))
		err := server.Listen()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	} else {
		fmt.Println("Starting client")
		host, port, err := snc.ParseArgs(os.Args)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		client := snc.NewClient(host, int(port))
		err = client.Dial()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		// connect to server
	}

}
