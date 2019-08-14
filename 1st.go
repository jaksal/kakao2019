package kakao2019

import (
	"strings"
)

type log struct {
	ltype string
	uid   string
	nick  string
}

func parselog(str string) *log {
	tmp := strings.Split(str, " ")

	l := &log{
		ltype: tmp[0],
		uid:   tmp[1],
	}

	if len(tmp) == 3 {
		l.nick = tmp[2]
	}

	return l
}

func k1st(record []string) []string {
	nicks := make(map[string]string)

	var logs []*log
	for _, r := range record {
		l := parselog(r)
		if l.ltype != "Change" {
			logs = append(logs, l)
		}

		nicks[l.uid] = l.nick
	}

	var result []string
	for _, l := range logs {
		r := nicks[l.uid]
		if l.ltype == "Enter" {
			r += "님이 들어왔습니다."
		} else if l.ltype == "Leave" {
			r += "님이 나갔습니다."
		}
		result = append(result, r)
	}
	return result
}
