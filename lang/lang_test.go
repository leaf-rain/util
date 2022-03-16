package lang

import (
	"fmt"
	"golang.org/x/text/language"
	"math/rand"
	"testing"
)

func TestInit(t *testing.T) {
	Init("./lang", language.English)
	fmt.Println(atI18n.mustGetMessage("hello_world"))
	ids := GetIds()
	id := ids[rand.Intn(len(ids))]
	fmt.Println(MustGetMessage(id, "en"))
}
