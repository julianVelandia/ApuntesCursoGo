package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(3, 2)
	if sum != 5 {
		t.Error("La suma no es correcta")
		t.Fail()
	} else {
		t.Log("Suma exitosa")
	}
}

func TestSubtract(t *testing.T) {
	subtraction := Subtract(3, 2)
	if subtraction != 1 {
		t.Error("La resta no es correcta")
		t.Fail()
	} else {
		t.Log("Resta exitosa")
	}
}
