package algorithm

import (
	"fmt"
	"testing"
)

func Test_RemoveExtraSpace(t *testing.T) {
	//r := ReverseWords("__bcc___d_ef__")
	//r := reverseWords("__bcc___d_ef__")
	//r := ReverseWords("the_sky_is_blue")
	//fmt.Println(r)
	s := []byte("__bcc___d_ef__")
	removeExtraSpaceBak(&s)
	fmt.Println(string(s))
}
