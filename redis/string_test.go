package redis_test

import (
	"testing"
	"time"

	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

// APPEND key value
// Available since 2.0.0.
// Time complexity: O(1).
func TestAppend(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()

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
}

// BITCOUNT key [start end]
// Available since 2.6.0.
// Time complexity: O(N)
func TestBitCount(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()

	bc, _ := c.BitCount("bits", &redis.BitCount{Start: 0, End: -1}).Result()
	assert.Equal(t, int64(0), bc)
	c.SetBit("bits", 0, 1)
	c.SetBit("bits", 3, 1)
	bc, _ = c.BitCount("bits", nil).Result()
	assert.Equal(t, int64(2), bc)
}

// BITOP operation destkey key [key ...]
// Available since 2.6.0.
// Time complexity: O(N)
func TestBitOp(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()

	c.Set("key1", "foobar", 0)
	c.Set("key2", "abcdef", 0)
	c.BitOpAnd("dest", "key1", "key2")
	bo, _ := c.Get("dest").Result()
	assert.Equal(t, "`bc`ab", bo)
}

// BITPOS key bit [start] [end]
// Available since 2.8.7.
// Time complexity: O(N)
func TestBitPos(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()

	c.Set("mykey", "\xff\xf0\x00", 0)
	bp, _ := c.BitPos("mykey", 0).Result()
	assert.Equal(t, int64(12), bp)
}

// DECR key
// Available since 1.0.0.
// Time complexity: O(1)
func TestDecr(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()

	c.Set("mykey", 10, 0)
	r, _ := c.Decr("mykey").Result()
	assert.Equal(t, int64(9), r)
	c.Set("mykey", "234293482390480948029348230948", 0)
	assert.Errorf(t, c.Decr("mykey").Err(), "ERR value is not an integer or out of range")
}

// DECRBY key decrement
// Available since 1.0.0.
// Time complexity: O(1)
func TestDecrBy(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()

	c.Set("mykey", 10, 0)
	r, _ := c.DecrBy("mykey", 3).Result()
	assert.Equal(t, int64(7), r)
}

// GET key
// Available since 1.0.0.
// Time complexity: O(1)
func TestGet(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()

	// r should be nil here, but string can not be nil in Go.
	r, _ := c.Get("nonexisting").Result()
	assert.Empty(t, r)
}

// GETBIT key offset
// Available since 2.2.0.
// Time complexity: O(1)
func TestGetBit(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()

	c.SetBit("mykey", 7, 1)
	b, _ := c.GetBit("mykey", 0).Result()
	assert.Equal(t, int64(0), b)
	b, _ = c.GetBit("mykey", 7).Result()
	assert.Equal(t, int64(1), b)
	b, _ = c.GetBit("mykey", 100).Result()
	assert.Equal(t, int64(0), b)
}

// GETRANGE key start end
// Available since 2.4.0.
// Time complexity: O(N)
func TestGetRange(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()

	c.Set("mykey", "This is a string", 0)
	// NOTE: it is [start, end] not [start, end)
	sub, _ := c.GetRange("mykey", 0, 3).Result()
	assert.Equal(t, "This", sub)
	sub, _ = c.GetRange("mykey", 10, 100).Result()
	assert.Equal(t, "string", sub)
}

// GETSET key value
// Available since 1.0.0.
// Time complexity: O(1)
func TestGetSet(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()

	c.Set("mykey", "Hello", 0)
	c.GetSet("mykey", "World")
	s, _ := c.Get("mykey").Result()
	assert.Equal(t, "World", s)
}

// INCR key
// Available since 1.0.0.
// Time complexity: O(1)
func TestIncr(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()

	k := "127.0.0.1:" + time.Now().Format("2006.01.02")
	if cur, _ := c.Get(k).Int64(); cur == 0 {
		// use MULTI and EXEC to avoid race condition.
		c.Pipelined(func(c redis.Pipeliner) error {
			c.Incr(k)
			c.Expire(k, time.Second)
			return nil
		})
		cur, _ = c.Get(k).Int64()
		assert.Equal(t, int64(1), cur)
	} else {
		c.Incr(k)
	}
}

// INCRBY key increment
// Available since 1.0.0.
// Time complexity: O(1)
func TestIncrBy(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()

	c.Set("mykey", 10, 0)
	c.IncrBy("mykey", 5)
	v, _ := c.Get("mykey").Int64()
	assert.Equal(t, int64(15), v)
}

// INCRBYFLOAT key increment
// Available since 2.6.0.
// Time complexity: O(1)
func TestIncrByFloat(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()

	v, _ := c.IncrByFloat("mykey", 9.0).Result()
	assert.Equal(t, float64(9), v)
}

// MGET key [key ...]
// Available since 1.0.0.
// Time complexity: O(N)
func TestMGet(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()

	c.Set("key1", "Hello", 0)
	c.Set("key2", "World", 0)
	v, _ := c.MGet("key1", "key2", "nonexisting").Result()
	assert.Equal(t, []interface{}{"Hello", "World", nil}, v)
}

// MSET key value [key value ...]
// Available since 1.0.1.
// Time complexity: O(N)
func TestMSet(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()

	c.MSet("key1", "Hello", "key2", "World")
	v, _ := c.Get("key1").Result()
	assert.Equal(t, "Hello", v)
}

// MSETNX key value [key value ...]
// Available since 1.0.1.
// Time complexity: O(N)
func TestMSetNX(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()

	f, _ := c.MSetNX("key1", "Hello", "key2", "there").Result()
	assert.True(t, f)
	f, _ = c.MSetNX("key2", "there", "key3", "World").Result()
	assert.False(t, f)
	v, _ := c.MGet("key1", "key2", "key3").Result()
	assert.Equal(t, []interface{}{"Hello", "there", nil}, v)
}

// PSETEX key milliseconds value
// Available since 2.6.0.
// Time complexity: O(1)
func TestPSetEX(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()

	// go-redis skip
}

// SETNX key value
// Available since 1.0.0.
// Time complexity: O(1)
// go-redis use new syntax: SET key value [EX seconds] [PX milliseconds] [NX|XX],
// if expiration is not 0;
// otherwise use old: SETNX key value.
func TestSetNX(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()
	// Design pattern: Locking with SETNX
	// SETNX lock.foo <current Unix time + lock timeout + 1>
	k, timeout := "lock.foo", time.Now().Unix()+1+1
	ok, _ := c.SetNX(k, timeout, 2*time.Second).Result()
	for time.Now().Unix() < timeout {
		if ok {
			// hold lock
			break
		} else {
			if t, _ := c.Get(k).Int64(); t < time.Now().Unix() {
				if t, _ = c.GetSet(k, time.Now().Unix()+1+1).Int64(); t < time.Now().Unix() {
					// hold lock
					ok = true
				}
			} else {
				time.Sleep(time.Millisecond)
				continue
			}
		}
	}
	if ok {
		// do sth
		// then release the lock
		c.Del(k)
	} else {
		// retry or return error
	}
}
