package hello

import "testing"

func TestHello(t *testing.T) {
	if v := Hello(); v != "hola, mundo!" {
		t.Errorf("Expect 'hola, mundo!', but got '%s'", v)
	}
}
