package ingestor

import (
	"testing"

	. "gopkg.in/check.v1"
)

func TestIngestor(t *testing.T) {
	TestingT(t)
}

type IngestorSuite struct{}

var _ = Suite(&IngestorSuite{})

func (s *IngestorSuite) TestRackspace(c *C) {
	n, _ := New("rackspace", nil)

	c.Assert(n, FitsTypeOf, &Rackspace{})
}
