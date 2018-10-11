package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/bmanth60/DuckTracker/data"
	"github.com/stretchr/testify/suite"
)

func initDB() {
	db, err := data.Connect()
	if err != nil {
		log.Fatal(err)
	}

	err = db.Reset()
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	initDB()
}

// Hook in test suite
func TestFunctional(t *testing.T) {
	suite.Run(t, new(FunctionalTestSuite))
}

type FunctionalTestSuite struct {
	suite.Suite
}

func (test *FunctionalTestSuite) SetupSuite() {
	// Run server
	c := make(chan string)
	setSetupChannel(c)

	go main()

	// Wait for server to start
	output := <-c
	if output != "done" {
		panic("not started")
	}

	// Add delay before sending first request
	time.Sleep(10 * time.Millisecond)
}

func (test *FunctionalTestSuite) TestGetRequest() {
	// I am a request
	// I want to get the form
	// I expect to get 200 and a form in the html

	client := getHttpClient()
	resp, err := client.Get("http://localhost:80/")
	test.NoError(err)

	test.Equal(http.StatusOK, resp.StatusCode)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	document := string(body)

	test.NotEmpty(document)
	test.Contains(document, "form")
}

func (test *FunctionalTestSuite) TestInvalidGetRequest() {
	// I am a request
	// I want to get an invalid url
	// I expect to get 404

	client := getHttpClient()
	resp, err := client.Get("http://localhost:80/badlink")
	test.NoError(err)

	test.Equal(http.StatusNotFound, resp.StatusCode)
}

func (test *FunctionalTestSuite) TestInvalidPutRequest() {
	// I am a request
	// I want to use an invalid method
	// I expect to get 404

	client := getHttpClient()
	req, err := http.NewRequest("PUT", "http://localhost:80/", nil)
	test.NoError(err)

	resp, err := client.Do(req)
	test.NoError(err)

	test.Equal(http.StatusNotFound, resp.StatusCode)
}

func (test *FunctionalTestSuite) TestInvalidPostDataRequest() {
	// I am a request
	// I want to use invalid form data
	// I expect to get 500

	client := getHttpClient()
	req, err := http.NewRequest("POST", "http://localhost:80/", strings.NewReader("data"))
	req.Header.Set("Content-Type", "text/plain; boundary=")
	test.NoError(err)

	resp, err := client.Do(req)
	test.NoError(err)

	test.Equal(http.StatusInternalServerError, resp.StatusCode)
}

func (test *FunctionalTestSuite) TestPostRequest() {
	// I am a form postback
	// I want to insert my form into the database
	// I expect to get 200 and a success message

	client := getHttpClient()
	data := "num_ducks=5&time_fed=2018-10-11T10:11&location=park&food_amount=10&food_name=bread&food_kind=grains"
	req, err := http.NewRequest("POST", "http://localhost:80/", strings.NewReader(data))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	test.NoError(err)

	resp, err := client.Do(req)
	test.NoError(err)

	test.Equal(http.StatusOK, resp.StatusCode)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	document := string(body)

	test.NotEmpty(document)
	test.Contains(document, "Successfully added")
}
