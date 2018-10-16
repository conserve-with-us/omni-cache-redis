package rediscache

import (
	"net/url"
	"time"

	"github.com/go-redis/redis"
)

// Cache contains options to connect to a badger database
// a TTL of 0 does not expire keys
type Cache struct {
	TTL time.Duration
}

// Conn is a connection to a redis database
type Conn struct {
	TTL    time.Duration
	client *redis.Client
}

// Stats displays stats about redis
type Stats map[string]interface{}

// NewCache creates a new Cache
func NewCache(defaultTimeout time.Duration) (*Cache, error) {
	return &Cache{TTL: defaultTimeout}, nil
}

// Open opens a new connection to redis
func (c Cache) Open(uri string) (*Conn, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return &Conn{}, err
	}
	pass, _ := u.User.Password()
	client := redis.NewClient(&redis.Options{
		Addr:     u.Hostname() + ":" + u.Port(),
		Password: pass,
		DB:       0,
	})
	return &Conn{TTL: c.TTL, client: client}, nil
}

// Close closes the redis connection
func (c *Conn) Close() error {
	return c.client.Close()
}

// Write writes data to the cache with the default cache TTL
func (c *Conn) Write(k, v []byte) error {
	return c.WriteTTL(k, v, c.TTL)
}

// WriteTTL writes data to the cache with an explicit TTL
// a TTL of 0 does not expire keys
func (c *Conn) WriteTTL(k, v []byte, ttl time.Duration) error {
	return c.client.Set(string(k), v, ttl).Err()
}

// Read retrieves data for a key from the cache
func (c *Conn) Read(k []byte) ([]byte, error) {
	return c.client.Get(string(k)).Bytes()
}

// Stats provides stats about the redis database
func (c *Conn) Stats() (map[string]interface{}, error) {
	info, err := c.client.Info().Result()
	if err != nil {
		return map[string]interface{}{}, err
	}
	// TODO: parse info into JSON
	return Stats{
		"Info": info,
	}, nil
}
