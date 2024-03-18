package main

// we use go get https://github.com/username/packagename
import (
	"fmt"

	tinytime "github.com/wagslane/go-tinytime"
)

func main() {
	tt := tinytime.Parse()
	fmt.Println(tt)
}
