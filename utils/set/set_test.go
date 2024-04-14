package set

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	s1 := []uint{
		1, 2, 3, 4,
	}
	s2 := []uint{
		1, 2, 3, 6,
	}
	fmt.Println(Difference(s2, s1))
}
