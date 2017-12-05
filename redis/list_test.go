package redis_test // for go-redis black test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/go-redis/redis"
)

// BLPOP key [key ...] timeout
// Available since 2.0.0.
// Time complexity: O(1)
// BLPOP is a blocking list pop primitive.
func TestBLPop(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()

	k, hk := "key", "keyHelp"
	// Pattern: Event notification
	for i := 1; i < 10; i++ {
		go func(i int) {
			c.Pipelined(func(c redis.Pipeliner) error {
				c.SAdd(k, i)
				// LPUSH helper_key x
				c.LPush(hk, i)
				return nil
			})
		}(i)
	}

	// for {
	// 	if v, _ := c.SPop(k).Int64(); v != 0 {
	// 		fmt.Println(v)
	// 	}
	// 	c.BRPop(100*time.Millisecond, hk).Err()
	// }
}

// BRPOP key [key ...] timeout
// Available since 2.0.0.
// Time complexity: O(1)
func TestBRPop(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()

	c.RPush("list1", "a", "b", "c")
	s, _ := c.BRPop(time.Second, "list1", "list2").Result()
	assert.Equal(t, []string{"list1", "c"}, s)
}

// BRPOPLPUSH source destination timeout
// Available since 2.2.0.
// Time complexity: O(1)
func TestBRPopLPush(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()

	c.RPush("key", 1, 2)
	c.BRPopLPush("key", "key1", time.Second)
	v, _ := c.LPop("key1").Int64()
	assert.Equal(t, int64(2), v)
}

// LINDEX key index
// Available since 1.0.0.
// Time complexity: O(N)
func TestLIndex(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()

	c.LPush("key1", 1, 2, 3, 4)         // 4 3 2 1
	v, _ := c.LIndex("key1", 2).Int64() // index is zero based, so is 2.
	assert.Equal(t, int64(2), v)
}

// LINSERT key BEFORE|AFTER pivot value
// Available since 2.2.0.
// Time complexity: O(N)
// Return value
// Integer reply: the length of the list after the insert operation, or -1 when the value pivot was not found.
func TestLInsert(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()

	c.RPush("key1", 1, 2, 3, 4, 5)
	c.LInsertBefore("key1", 2, 6)
	v, _ := c.LIndex("key1", 1).Int64()
	assert.Equal(t, int64(6), v)
}

// LLEN key
// Available since 1.0.0.
// Time complexity: O(1)
func TestLLen(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()

	c.LPush("mykey", 1)
	v, _ := c.LLen("mykey").Result()
	assert.Equal(t, int64(1), v)
}

// LPOP key
// Available since 1.0.0.
// Time complexity: O(1)
// Removes and returns the first element of the list stored at key.
func TestLPop(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()

	c.RPush("mylist", "one", "two")
	v, _ := c.LPop("mylist").Result()
	assert.Equal(t, "one", v)
}

// LPUSH key value [value ...]
// Available since 1.0.0.
// Time complexity: O(1)
func TestLPush(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()

	c.LPush("mylist", 1, 2, 3, 4)
	var v []int64
	// c.LRange does not support other slices except string ones.
	c.LRange("mylist", 0, -1).ScanSlice(&v)
	assert.Equal(t, []int64{4, 3, 2, 1}, v)
}

// LPUSHX key value
// Available since 2.2.0.
// Time complexity: O(1)
func TestLPushX(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()

	v, _ := c.LPushX("mykey", 1).Result()
	assert.Equal(t, int64(0), v)
}

// LRANGE key start stop
// Available since 1.0.0.
// Time complexity: O(S+N)
func TestLRange(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()

	c.RPush("mylist", 1, 2, 3, 4, 5, 6, 7, 8)
	var v []int
	// NOTE: here it is [start, end] not [start, end)
	c.LRange("mylist", 0, 3).ScanSlice(&v)
	assert.Equal(t, []int{1, 2, 3, 4}, v)
	var v1 []int
	c.LRange("mylist", -3, 7).ScanSlice(&v1)
	assert.Equal(t, []int{6, 7, 8}, v1)
}

// LREM key count value
// Available since 1.0.0.
// Time complexity: O(N)
func TestLRem(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()

	c.RPush("mylist", "hello", "hello", "foo", "hello")
	c.LRem("mylist", -2, "hello")
	v, _ := c.LRange("mylist", 0, -1).Result()
	assert.Equal(t, []string{"hello", "foo"}, v)
}

// LSET key index value
// Available since 1.0.0.
// Time complexity: O(N)
func TestLSet(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()

	c.RPush("mylist", "one", "two", "three")
	c.LSet("mylist", 0, "four")
	c.LSet("mylist", -2, "five")
	v, _ := c.LRange("mylist", 0, -1).Result()
	assert.Equal(t, []string{"four", "five", "three"}, v)
}

// LTRIM key start stop
// Available since 1.0.0.
// Time complexity: O(N)
func TestLTrim(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()

	// A common use of LTRIM is together with LPUSH / RPUSH. For example:

	c.LPush("mylist", 0, 1, 2, 3, 4)
	c.LPush("mylist", 5)
	c.LTrim("mylist", 0, 4) // keep mylist cap is 5
	v, _ := c.LLen("mylist").Result()
	assert.Equal(t, int64(5), v)
}

// RPOP key
// Available since 1.0.0.
// Time complexity: O(1)
func TestRPop(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()

	c.RPush("mylist", "one", "two", "three")
	c.RPop("mylist")
	v, _ := c.LRange("mylist", 0, -1).Result()
	assert.Equal(t, []string{"one", "two"}, v)
}

// RPOPLPUSH source destination
// Available since 1.2.0.
// Time complexity: O(1)
func TestRPopLPush(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()

	c.RPush("mylist", "one", "two", "three")
	c.RPopLPush("mylist", "mylist1")
	v, _ := c.LRange("mylist", 0, -1).Result()
	assert.Equal(t, []string{"one", "two"}, v)
	v, _ = c.LRange("mylist1", 0, -1).Result()
	assert.Equal(t, []string{"three"}, v)
}

// RPUSH key value [value ...]
// Available since 1.0.0.
// Time complexity: O(1)
func TestRPush(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()
	c.RPush("mylist", "hello", "world")
	v, _ := c.LRange("mylist", 0, -1).Result()
	assert.Equal(t, []string{"hello", "world"}, v)
}

// RPUSHX key value
// Available since 2.2.0.
// Time complexity: O(1)
func TestRPushX(t *testing.T) {
	safeFlushDB()
	defer safeFlushDB()

	v, _ := c.RPushX("mykey", "hello").Result()
	assert.Equal(t, int64(0), v)
}
