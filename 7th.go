package kakao2019

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1st(t *testing.T) {
	t.Log("start kakao 1st")

	tc := []struct {
			n          int
			arr1, arr2 []int
			out        []string
	}{
			// http://tech.kakao.com/2017/09/27/kakao-blind-recruitment-round-1/
			// 1번문제 답에서 연속된 공백이 하나로 나오는데.
			// 네가 잘못푼건지. 답이 잘못된건지 모르겠음.
			// ["#####","# # #", "### #", "# ##", "#####"]
			// ["######", "### #", "## ##", " #### ", " #####", "### # "]
			//
			{5, []int{9, 20, 28, 18, 11}, []int{30, 1, 21, 17, 28}, []string{"#####", "# # #", "### #"
, "#  ##", "#####"}},
			{6, []int{46, 33, 33, 22, 31, 50}, []int{27, 56, 19, 14, 14, 10}, []string{"######", "### 
#", "##  ##", " #### ", " #####", "### # "}},
	}

	for i, tt := range tc {
			t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
					result := k1st(tt.n, tt.arr1, tt.arr2)
					assert.Equal(t, tt.out, result)
			})
	}
}