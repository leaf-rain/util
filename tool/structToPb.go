package tool

import (
	"fmt"
	"reflect"
)

type stu struct {
	Mid      string `json:"mid"`
	Amount   int    `json:"amount"`
	OrderID  string `json:"orderid"`
	TxnID    string `json:"txnid"`
	Status   int    `json:"status"`
	Checksum string `json:"checksum"`
}

func main() {
	var a = stu{}
	var tf = reflect.TypeOf(a)
	fmt.Println(tf.Name(), tf.Kind())
	for i := 0; i < tf.NumField(); i++ {
		ft := tf.Field(i)
		var ty string
		switch ft.Type.String() {
		case "int":
			ty = "int64"
		case "uint":
			ty = "uint64"
		case "float64":
			ty = "double"
		case "float32":
			ty = "float"
		case "[]byte":
			ty = "bytes"
		case "uint8":
			ty = "uint32"
		default:
			ty = ft.Type.String()
		}
		fmt.Println(fmt.Sprintf("%s %s = %d;", ty, ft.Name, i+1))
	}
}
