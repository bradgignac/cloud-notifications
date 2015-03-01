package config

import (
	"os"
	"testing"

	. "gopkg.in/check.v1"
)

func TestOptions(t *testing.T) {
	TestingT(t)
}

type OptionsSuite struct{}

var _ = Suite(&OptionsSuite{})

func (s *OptionsSuite) TestOptionFromKey(c *C) {
	os.Setenv("one", "")
	config := map[string]interface{}{
		"one": "from key",
	}

	val, err := ReadOption(Option{Key: "one", Env: "one"}, config)

	c.Assert(err, IsNil)
	c.Assert(val, Equals, "from key")
}

func (s *OptionsSuite) TestOptionFromEnv(c *C) {
	os.Setenv("one", "from env")
	config := map[string]interface{}{
		"one": "from key",
	}

	val, err := ReadOption(Option{Key: "one", Env: "one"}, config)

	c.Assert(err, IsNil)
	c.Assert(val, Equals, "from env")
}

func (s *OptionsSuite) TestMissingOption(c *C) {
	os.Setenv("one", "")
	config := map[string]interface{}{}

	val, err := ReadOption(Option{Key: "one", Env: "one"}, config)

	c.Assert(err, ErrorMatches, "Missing option \"one\"")
	c.Assert(val, Equals, "")
}

func (s *OptionsSuite) TestReturnsAllOptions(c *C) {
	config := map[string]interface{}{
		"one": "from key",
		"two": "from key",
	}

	val, err := ReadOptions([]Option{
		Option{Key: "one"},
		Option{Key: "two"},
	}, config)

	c.Assert(err, IsNil)
	c.Assert(val, DeepEquals, map[string]string{
		"one": "from key",
		"two": "from key",
	})
}

func (s *OptionsSuite) TestReturnsFirstError(c *C) {
	config := map[string]interface{}{
		"exists": "from key",
	}

	val, err := ReadOptions([]Option{
		Option{Key: "missing"},
		Option{Key: "exists"},
	}, config)

	c.Assert(err, ErrorMatches, "Missing option \"missing\"")
	c.Assert(val, IsNil)
}
