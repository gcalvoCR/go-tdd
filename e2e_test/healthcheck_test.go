package e2e_test

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type EntoEndSuite struct {
	suite.Suite
}

func TestEndToEndSuite(t *testing.T) {
	suite.Run(t, new(EntoEndSuite))
}

func (s *EntoEndSuite) TestHappyHealthcheck() {
	c := http.Client{}

	r, _ := c.Get("http://localhost:8080/healthcheck")

	b, _ := io.ReadAll(r.Body)

	s.Equal(http.StatusOK, r.StatusCode)
	s.JSONEq(`{"status":"OK", "messages": []}`, string(b))
}
