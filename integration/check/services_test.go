package composeit

import (
	"testing"

	"github.com/go-check/check"
	compose "github.com/libkermit/compose/check"
)

// Hook up gocheck into the "go test" runner
func TestServices(t *testing.T) { check.TestingT(t) }

type CheckServicesSuite struct{}

var _ = check.Suite(&CheckServicesSuite{})

func (s *CheckServicesSuite) TestServicesProject(c *check.C) {
	project := compose.CreateProject(c, "services", "./assets/services.yml")
	project.Start(c, "hello")

	container := project.Container(c, "hello")
	c.Assert(container.Name, check.Equals, "/services_hello_1")

	//"No container found for '%s' service
	project.NoContainer(c, "other")

	project.Stop(c, "hello")

	project.Start(c)

	container = project.Container(c, "other")
	c.Assert(container.Name, check.Equals, "/services_other_1")

	project.Stop(c)
}
