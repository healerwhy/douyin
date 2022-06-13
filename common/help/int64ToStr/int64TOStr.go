package int64ToStr

import (
	"strconv"
	"strings"
)

func Int64ToStr(arr []int64) string {
	ids := make([]string, 0, len(arr))
	for _, id := range arr {
		ids = append(ids, strconv.FormatInt(id, 10))
	}
	idsStr := strings.Join(ids, ",")
	return idsStr
}
