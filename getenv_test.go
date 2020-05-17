package getenv

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	expectedDefaultValue = ":5000"
)

type GetEnvSuite struct {
	suite.Suite
}

func (s *GetEnvSuite) TestGetEnvReturnsDefault_Succeeds() {
	value := EnvOrDefault("PORT", String(expectedDefaultValue))

	s.Equal(expectedDefaultValue, value)
}

func (s *GetEnvSuite) TestGetEnvReturnsEnvOverDefault_Succeeds() {
	os.Setenv("PORT", ":9000")

	value := EnvOrDefault("PORT", String(expectedDefaultValue))

	s.Equal(":9000", value)
}

func (s *GetEnvSuite) TestGetEnvPanics_Succeeds() {
	s.Assert().Panics(func() { EnvOrDefault("PORT", Nil()) }, "Expected panic of missing environment variable")
}

func TestGetEnvSuite(t *testing.T) {
	suite.Run(t, new(GetEnvSuite))
}
