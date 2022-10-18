//go:generate msgp
package main

type ArgumentBody struct {
	Ctx  interface{} `msg:"ctx"`
	Id   string      `msg:"id"`
	Body []byte      `msg:"body"`
	Any  interface{} `msg:"any"`
}
