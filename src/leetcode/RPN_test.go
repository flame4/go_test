package leetcode

import(
	"testing"
	"strings"
)
func TestRPN_1(t *testing.T) {
	if(strings.Compare("352*+2+5/63/2*+3+", rpn("((3+5*2)+2)/5+6/3*2+3")) != 0) {
		t.Fatal("Error!")
	}
}
