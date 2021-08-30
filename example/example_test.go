package example

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	r := add(2, 5)
	if r != 7 {
		t.Fatalf("add(2, 5) error, expect:%d, actual:%d", 6, r)
	}
	t.Logf("test succ..")
}

func TestSub(t *testing.T) {
	r := sub(1, 2)
	if r != -1 {
		t.Error("haha, ... fail")
	} else {
		fmt.Print("hah, happy...")
		t.Logf("you are right..")
	}
}
