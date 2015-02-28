package main

import (
	"testing"

	. "gopkg.in/check.v1"
)

func TestConfig(t *testing.T) {
	TestingT(t)
}

type ConfigSuite struct{}

var _ = Suite(&ConfigSuite{})

func (s *ConfigSuite) TestMainConfig(c *C) {
	config, err := LoadYAML("../support/config.yaml")

	c.Assert(err, IsNil)
	c.Assert(config.Endpoints, DeepEquals, []string{
		"https://dfw.feeds.api.rackspacecloud.com",
		"https://iad.feeds.api.rackspacecloud.com",
		"https://ord.feeds.api.rackspacecloud.com",
	})
	c.Assert(config.Feeds, DeepEquals, []string{
		"backup_events_obs",
		"bigdata_events_obs",
		"ssl_usagesummary_events_obs",
	})
}

func (s *ConfigSuite) TestIngestorConfig(c *C) {
	config, err := LoadYAML("../support/config.yaml")
	ingestor := config.Ingestor

	c.Assert(err, IsNil)
	c.Assert(ingestor.Type, Equals, "rackspace")
	c.Assert(ingestor.Options, DeepEquals, map[string]interface{}{
		"user": "foo",
		"key":  "bar",
	})
}

func (s *ConfigSuite) TestNotifierConfig(c *C) {
	config, err := LoadYAML("../support/config.yaml")
	notifier := config.Notifier

	c.Assert(err, IsNil)
	c.Assert(notifier.Type, Equals, "twilio")
	c.Assert(notifier.Options, DeepEquals, map[string]interface{}{
		"user": "foo",
		"key":  "bar",
		"from": "from",
		"to":   "to",
	})
}

func (s *ConfigSuite) TestMissingConfig(c *C) {
	_, err := LoadYAML("../support/missing.yaml")

	c.Assert(err, ErrorMatches, "open ../support/missing.yaml: no such file or directory")
}
