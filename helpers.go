package rediscache

import (
	"bufio"
	"strings"
)

func parseRedisInfo(redisInfo string) (map[string]interface{}, error) {
	data := map[string]interface{}{}
	scanner := bufio.NewScanner(strings.NewReader(redisInfo))
	for scanner.Scan() {
		t := scanner.Text()
		s := strings.SplitN(t, ":", 2)
		if len(s) == 2 {
			data[s[0]] = s[1]
		}
	}
	if err := scanner.Err(); err != nil {
		return data, err
	}
	return data, nil
}
