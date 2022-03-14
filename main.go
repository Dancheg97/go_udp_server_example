package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	v, e := os.LookupEnv("EXPIRATION")
	fmt.Println(e, v)
	time.Sleep(time.Second)
}
