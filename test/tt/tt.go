package main

import "fmt"

type b uint
type a struct {
	A b `json:"a"`
}

func main() {
	alist := []a{{A: 1}, {A: 2}, {A: 3}}
	blist := make([]*a, len(alist))
	for i := range alist {
		//value := alist[i]
		blist[i] = &alist[i]
	}
	fmt.Println(blist[0], blist[1])
}
