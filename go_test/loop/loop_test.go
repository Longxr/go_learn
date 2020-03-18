package loop

import "testing"

func TestWhileLoop(t *testing.T) {

	for n := 0; n < 5; n++ {
		t.Log(n)
	}
}
