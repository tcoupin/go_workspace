package crypto

import "testing"

func TestCesar(t *testing.T) {
	receive := Cesar("salut",1)
	if receive != "tbmvu" {
		t.Errorf("result=%q expect=tbmvu",receive)
	}
}