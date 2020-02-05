package main

import "my-echo.com/src/app/infrastructure"

func main() {
	// Start server
	infrastructure.Router.Logger.Fatal(infrastructure.Router.Start(":1323"))
}
