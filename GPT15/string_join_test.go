package GPT15

import (
	"testing"
)

func TestHttpServer(t *testing.T) {
	res := StringJoin([]string{"a", "b", "c"}, ",")
	if res != "a,b,c" {
		t.Errorf("want a,b,c, got %s", res)
	}
}
