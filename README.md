# OmniCache Redis

A [Redis](https://redis.io/) persistence layer for [omni-cache](https://github.com/panoplymedia/omni-cache).
# Local Cache Redis

### Sample Usage

```go
defaultTimeout := time.Minute
cache, err := NewCache(defaultTimeout)
if err != nil {
  fmt.Println(err)
}

// open a connection to badger database
conn, err := cache.Open("redis://user:password@localhost:6379")
defer conn.Close()

// write data to redis (uses defaultTimeout)
err = conn.Write([]byte("key"), []byte("data"))

// write data to redis with custom timeout
err = conn.WriteTTL([]byte("key2"), []byte("data"), 5*time.Minute)

// read data
data, err := conn.Read([]byte("key"))

// log stats
fmt.Println(conn.Stats())
```
