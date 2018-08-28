package config

import (
	"testing"

	"github.com/aphistic/sweet"
	"github.com/aphistic/sweet-junit"
	. "github.com/onsi/gomega"
)

func TestMain(m *testing.M) {
	RegisterFailHandler(sweet.GomegaFail)

	sweet.Run(m, func(s *sweet.S) {
		s.RegisterPlugin(junit.NewPlugin())

		s.AddSuite(&BuildSuite{})
		s.AddSuite(&ConfigSuite{})
		s.AddSuite(&PlanSuite{})
		s.AddSuite(&PushSuite{})
		s.AddSuite(&RemoveSuite{})
		s.AddSuite(&ResolverSuite{})
		s.AddSuite(&RunSuite{})
		s.AddSuite(&StageSuite{})
		s.AddSuite(&UtilSuite{})
	})
}
