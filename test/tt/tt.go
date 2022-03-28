package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	m.Store(1, 1)
	fmt.Println(m.Load(1))
	m = sync.Map{}
	fmt.Println(m.Load(1))
}
