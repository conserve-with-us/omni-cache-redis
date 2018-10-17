package rediscache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRedisInfo(t *testing.T) {
	info := `
redis_version:4.0.9
redis_git_sha1:00000000
redis_git_dirty:0
redis_build_id:d3ebfc7feabc1290
used_cpu_sys:84.433
	`
	p, err := parseRedisInfo(info)
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{
		"redis_version":   "4.0.9",
		"redis_git_sha1":  "00000000",
		"redis_git_dirty": "0",
		"redis_build_id":  "d3ebfc7feabc1290",
		"used_cpu_sys":    "84.433",
	}, p)
}
