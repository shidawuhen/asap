package main

import (
	_ "asap/docs"
	"fmt"
	"time"

	f "asap/framework"
)

func ping(c *f.Context) {
	fmt.Println("Response successful！", time.Now().Format("2006-01-02 15:04:05"))
	c.String("%s", "ping")
}

func main() {
	e := f.New()
	e.AddRoute("GET", "/ping", ping)
	e.Run("127.0.0.1:9090")
}
