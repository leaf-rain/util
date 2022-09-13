// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // JSON

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// JSONListener is a complete listener for a parse tree produced by JSONParser.
type JSONListener interface {
	antlr.ParseTreeListener

	// EnterJson is called when entering the json production.
	EnterJson(c *JsonContext)

	// EnterObj is called when entering the obj production.
	EnterObj(c *ObjContext)

	// EnterPair is called when entering the pair production.
	EnterPair(c *PairContext)

	// EnterArr is called when entering the arr production.
	EnterArr(c *ArrContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// ExitJson is called when exiting the json production.
	ExitJson(c *JsonContext)

	// ExitObj is called when exiting the obj production.
	ExitObj(c *ObjContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)

	// ExitArr is called when exiting the arr production.
	ExitArr(c *ArrContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)
}
