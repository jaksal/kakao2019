package kakao2019

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test3rd(t *testing.T) {
	t.Log("start kakao 3rd")

	tc := []struct {
		relation [][]string
		out      int
	}{
		{[][]string{
			[]string{"100", "ryan", "music", "2"},
			[]string{"200", "apeach", "math", "2"},
			[]string{"300", "tube", "computer", "3"},
			[]string{"400", "con", "computer", "4"},
			[]string{"500", "muzi", "music", "3"},
			[]string{"600", "apeach", "music", "2"},
		}, 2},
	}

	for i, tt := range tc {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			result := k3rd(tt.relation)
			assert.Equal(t, tt.out, result)
		})
	}
}
