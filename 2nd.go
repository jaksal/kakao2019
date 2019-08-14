package kakao2019

import (
	"fmt"
	"sort"
)

type stage struct {
	no    int
	fail  int
	total int
}

func (s *stage) calc() float64 {
	if s.total == 0 {
		return 0
	}
	return float64(s.fail) / float64(s.total)
}

func k2nd(n int, stages []int) []int {
	ts := make([]*stage, n+1)
	for i := 0; i < n+1; i++ {
		ts[i] = &stage{no: i + 1}
	}

	for _, s := range stages {
		for i := 1; i <= s; i++ {
			ts[i-1].total++
		}
		ts[s-1].fail++
	}
	ts = ts[:n]

	sort.Slice(ts, func(i, j int) bool {
		if ts[i].calc() > ts[j].calc() {
			return true
		} else if ts[i].calc() < ts[j].calc() {
			return false
		}

		return ts[i].no < ts[j].no
	})

	var result []int
	for i := 0; i < len(ts); i++ {
		fmt.Printf("no=%d fail=%d total=%d\n", ts[i].no, ts[i].fail, ts[i].total)
		result = append(result, ts[i].no)
	}

	return result
}
