package fangless_test

import (
	"math"
	"testing"
	"time"
	"github.com/opalmer/fangless"
	"github.com/spf13/viper"
	. "gopkg.in/check.v1"
)

type SnakeTest struct{}

var _ = Suite(&SnakeTest{})

func Test(t *testing.T) { TestingT(t) }

func (s *SnakeTest) TestLockUnlock(c *C) {
	snake := fangless.New(viper.New())
	snake.Lock()
	snake.Unlock() // Should not panic
}

func (s *SnakeTest) TestUnsafe(c *C) {
	snake := fangless.New(viper.New())
	snake.Viper()
	snake.Unlock() // Should not panic
}

func (s *SnakeTest) TestGetFunctions(c *C) {
	now := time.Now().UTC()
	snake := fangless.New(viper.New())
	values := map[string]interface{}{
		"string": "string",
		"int": 1,
		"bool": true,
		"int64": math.MaxInt64,
		"float64": 1.000000000000000000000000000000000000000000000000001,
		"time": now,
		"duration": "5s",
	}
	for key, value := range values {
		snake.Set(key, value)
	}
	c.Assert(snake.Get("string"), Equals, "string")
	c.Assert(snake.GetString("string"), Equals, "string")
	c.Assert(snake.GetBool("bool"), Equals, true)
	c.Assert(snake.GetInt("int"), Equals, 1)
	c.Assert(snake.GetInt64("int64"), Equals, int64(math.MaxInt64))
	c.Assert(
		snake.GetFloat64("float64"), Equals,
		float64(1.000000000000000000000000000000000000000000000000001))
	c.Assert(snake.GetTime("time"), Equals, now)
	c.Assert(
		snake.GetDuration("duration"), Equals, time.Second * 5)
	// TODO map function tests
}

func (s *SnakeTest) TestIsSet(c *C) {
	snake := fangless.New(viper.New())
	c.Assert(snake.IsSet("foo"), Equals, false)
	snake.Set("foo", "")
	c.Assert(snake.IsSet("foo"), Equals, true)
}

func (s *SnakeTest) TestSetDefault(c *C) {
	snake := fangless.New(viper.New())
	c.Assert(snake.Get("foo"), Equals, nil)
	snake.SetDefault("foo", "bar")
	c.Assert(snake.Get("foo"), Equals, "bar")
	snake.Set("foo", "bar3")
	c.Assert(snake.Get("foo"), Equals, "bar3")
}

