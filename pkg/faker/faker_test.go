package faker

import (
	"testing"
	"unicode/utf8"
)

func Test_rand_int_spec(t *testing.T) {
	result := RandIntSpec()
	
	if utf8.RuneCountInString(result) != 6 {
		t.Fatal("not enough size")
	}
	
	t.Logf("%v\n", result)
}

func Test_show(t *testing.T) {
	t.Log(ID())       // RsQW-oOXT-LOqO-WHxJ
	t.Log(Username()) // TXFgLIM
	t.Log(Email())    // tmlgeytga@example.com
}

