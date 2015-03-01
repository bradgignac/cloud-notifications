package notifier

import (
	"testing"

	. "gopkg.in/check.v1"
)

func TestNotifier(t *testing.T) {
	TestingT(t)
}

type NotifierSuite struct{}

var _ = Suite(&NotifierSuite{})

func (s *NotifierSuite) TestConsole(c *C) {
	n, _ := New("console", nil)

	c.Assert(n, FitsTypeOf, &Console{})
}

func (s *NotifierSuite) TestTwilio(c *C) {
	n, _ := New("twilio", map[string]interface{}{
		"account": "account",
		"token":   "token",
		"from":    "from",
		"to":      "to",
	})

	c.Assert(n, FitsTypeOf, &Twilio{})
}
