package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Put your database host")
	var user *string = flag.String("user", "root", "Put your database user")
	var password *string = flag.String("password", "root", "Put your database passsword")
	var port *int = flag.Int("port", 8080, "Put your port")
	flag.Parse()

	fmt.Println("Host :", *host)
	fmt.Println("User :", *user)
	fmt.Println("Password :", *password)
	fmt.Println("Port :", *port)
}
