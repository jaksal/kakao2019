package kakao2019

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test2nd(t *testing.T) {
	t.Log("start kakao 2nd")

	tc := []struct {
		n      int
		stages []int
		out    []int
	}{
		{5, []int{2, 1, 2, 6, 2, 4, 3, 3}, []int{3, 4, 2, 1, 5}},
		{4, []int{4, 4, 4, 4, 4}, []int{4, 1, 2, 3}},
	}

	for i, tt := range tc {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			result := k2nd(tt.n, tt.stages)
			assert.Equal(t, tt.out, result)
		})
	}
}
