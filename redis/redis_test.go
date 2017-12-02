package redis

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-redis/redis"
)

var c = redis.NewClient(&redis.Options{
	Addr: "127.0.0.1:6379",
	DB:   8,
})

func init() {
	if err := c.Ping().Err(); err != nil {
		log.Fatalln(err)
	}
}

func TestStringOps(t *testing.T) {
	c.FlushDB()
	defer c.FlushDB()

	// APPEND key value
	// O(1)
	r, _ := c.Exists("mykey").Result()
	assert.Equal(t, int64(0), r)

	c.Append("mykey", "Hello")
	c.Append("mykey", " World")

	v, _ := c.Get("mykey").Result()
	assert.Equal(t, "Hello World", v)

	/* Pattern: Time series */
	c.Append("ts", "0043")
	c.Append("ts", "0035")
	ts, _ := c.GetRange("ts", 0, 3).Result()
	assert.Equal(t, "0043", ts)
	ts, _ = c.GetRange("ts", 4, 7).Result()
	assert.Equal(t, "0035", ts)

	c.SetRange("ts", 0, "0073")
	ts, _ = c.GetRange("ts", 0, 3).Result()
	assert.Equal(t, "0073", ts)

	tl, _ := c.StrLen("ts").Result()
	assert.Equal(t, int64(8), tl)

	// BITCOUNT key [start end]
	// O(N)
	bc, _ := c.BitCount("bits", &redis.BitCount{Start: 0, End: -1}).Result()
	assert.Equal(t, int64(0), bc)
	c.SetBit("bits", 0, 1)
	c.SetBit("bits", 3, 1)
	bc, _ = c.BitCount("bits", nil).Result()
	assert.Equal(t, int64(2), bc)

	// BITFIELD [GET type offset] [SET type offset value] [INCRBY type offset increment] [OVERFLOW WRAP|SAT|FAIL]
	// O(1)
	// skip ...

	// BITOP operation destkey key [key ...]
	// BITOP AND, [OR, XOR and NOT]
	// O(N)
	c.Set("key1", "foobar", 0)
	c.Set("key2", "abcdef", 0)
	c.BitOpAnd("dest", "key1", "key2")
	bo, _ := c.Get("dest").Result()
	assert.Equal(t, "`bc`ab", bo)

	// BITPOS key bit [start] [end]
	// O(N)
	c.Set("mykey", "\xff\xf0\x00", 0)
	bp, _ := c.BitPos("mykey", 0).Result()
	assert.Equal(t, int64(12), bp)

}
