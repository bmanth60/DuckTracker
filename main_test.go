package main

import (
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// Hook in test suite
func TestIntegration(t *testing.T) {
	suite.Run(t, new(FunctionalTestSuite))
}

type FunctionalTestSuite struct {
	suite.Suite
}

func (test *FunctionalTestSuite) SetupSuite() {
	// Load env
	//setConfig(getTestEnvironment())

	//Run server
	c := make(chan string)
	setSetupChannel(c)

	go main()

	//Wait for server to start
	output := <-c
	if output != "done" {
		panic("not started")
	}

	//Add delay before sending first request
	time.Sleep(10 * time.Millisecond)
}

func (test *FunctionalTestSuite) TestGetRequest() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Get("http://localhost:80/")
	test.NoError(err)

	test.Equal(http.StatusOK, resp.StatusCode)
}
