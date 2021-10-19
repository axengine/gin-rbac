package types

import (
	"strconv"
	"strings"
)

type IntSplitStr string

func (iss IntSplitStr) Unmarshal() []int64 {
	strs := strings.Split(string(iss), ",")
	ints := make([]int64, 0)
	for _, s := range strs {
		v, err := strconv.ParseInt(s, 10, 64)
		if err == nil {
			ints = append(ints, v)
		}
	}
	return ints
}

func (iss IntSplitStr) Marshal(ints []int64) IntSplitStr {
	strs := make([]string, 0)
	for _, i := range ints {
		strs = append(strs, strconv.FormatInt(i, 10))
	}
	return IntSplitStr(strings.Join(strs, ","))
}
