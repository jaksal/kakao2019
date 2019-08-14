package kakao2019

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1st(t *testing.T) {
	t.Log("start kakao 1st")

	tc := []struct {
		record []string
		out    []string
	}{
		{[]string{"Enter uid1234 Muzi", "Enter uid4567 Prodo", "Leave uid1234", "Enter uid1234 Prodo", "Change uid4567 Ryan"}, []string{"Prodo님이 들어왔습니다.", "Ryan님이 들어왔습니다.", "Prodo님이 나갔습니다.", "Prodo님이 들어왔습니다."}},
	}

	for i, tt := range tc {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			result := k1st(tt.record)
			assert.Equal(t, tt.out, result)
		})
	}
}
