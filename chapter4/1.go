package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("三秒待つ")
	<-time.After(time.Duration(3 * time.Second))
	fmt.Println("三秒待った")
}
