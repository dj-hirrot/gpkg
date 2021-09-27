package main

import (
	"fmt"
	"three/mypkg"
	"three/mypkg/underpkg"
)

func main() {
	fmt.Println(mypkg.Rectangle(4, 5))
	mypkg.Intro()

	underpkg.Hello()
}
