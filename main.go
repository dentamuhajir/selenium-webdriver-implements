package main

import (
	"fmt"
	"os"

	"github.com/tebeka/selenium"
)

func main() {
	//
}

func example() {
	const (
		seleniumPath    = "script/selenium-server.jar"
		geckoDriverPath = "scripts/firefox"
		port            = 8080
	)
	opts := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
		selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
		selenium.Output(os.Stderr),            // Output debug information to STDERR.
	}
	selenium.SetDebug(true)

	fmt.Println(opts)
}
