package main

import (
	_ "asap/docs"
	"fmt"

	f "asap/framework"
)

func ping(c *f.Context) {
	fmt.Println(1)
	c.String("%s", "ping")
}

func main() {
	e := f.New()
	e.AddRoute("GET", "/ping", ping)
	e.Run("127.0.0.1:9090")
}
