package rediscache

import (
	"os"
	"testing"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert"
)

func TestNewCache(t *testing.T) {
	c, err := NewCache(time.Second)
	assert.Nil(t, err)

	assert.Equal(t, time.Second, c.TTL)
}

func TestWrite(t *testing.T) {
	c, err := NewCache(time.Second)
	assert.Nil(t, err)
	conn, err := c.Open(os.Getenv("REDIS_URL"))
	assert.Nil(t, err)
	defer conn.Close()

	key := []byte("set")

	// cache miss
	b := []byte{1, 2, 3}
	err = conn.Write(key, b)
	assert.Nil(t, err)

	// cache hit
	b2, err := conn.Read(key)
	assert.Nil(t, err)
	assert.Equal(t, b, b2)

	// default ttl timeout (cache miss)
	time.Sleep(time.Second)
	_, err = conn.Read(key)
	assert.Errorf(t, err, "Key not found")
}

func TestWriteTTL(t *testing.T) {
	c, err := NewCache(time.Second)
	assert.Nil(t, err)
	conn, err := c.Open(os.Getenv("REDIS_URL"))
	assert.Nil(t, err)
	defer conn.Close()

	key := []byte("set")

	// cache miss
	b := []byte{1, 2, 3}
	err = conn.WriteTTL(key, b, time.Second)
	assert.Nil(t, err)

	// cache hit
	b2, err := conn.Read(key)
	assert.Nil(t, err)
	assert.Equal(t, b, b2)

	// default ttl timeout (cache miss)
	time.Sleep(time.Second)
	_, err = conn.Read(key)
	assert.Errorf(t, err, "Key not found")
}

func TestRead(t *testing.T) {
	c, err := NewCache(time.Second)
	assert.Nil(t, err)
	conn, err := c.Open(os.Getenv("REDIS_URL"))
	assert.Nil(t, err)
	defer conn.Close()

	// creates initial key
	key := []byte("my-key")
	// cache miss
	b, err := conn.Read(key)
	assert.Errorf(t, err, "Key not found")

	// cache hit
	v := []byte{1, 2}
	err = conn.Write(key, v)
	assert.Nil(t, err)
	b, err = conn.Read(key)
	assert.Nil(t, err)
	assert.Equal(t, v, b)
}

func TestStats(t *testing.T) {
	c, err := NewCache(time.Second)
	assert.Nil(t, err)
	conn, err := c.Open(os.Getenv("REDIS_URL"))
	assert.Nil(t, err)
	defer conn.Close()

	s, err := conn.Stats()
	assert.Nil(t, err)
	assert.NotEqual(t, "", s["redis_version"])
}
