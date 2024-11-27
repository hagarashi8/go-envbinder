package envbinder

import (
	"fmt"
	"testing"
)

type F struct{}

func TestHelloName(t *testing.T) {
	type T struct {
		L struct {
			Name string `env:"PWD"`
		}
	}
	var l T
	Bind(&l)
	fmt.Println(l)
}
