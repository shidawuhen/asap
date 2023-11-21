package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	f "asap/framework"

	"github.com/stretchr/testify/assert"
)

func pong(c *f.Context) {
	c.String("%s", "pong")
}

func TestPingRoute(t *testing.T) {
	r := f.New()
	r.AddRoute("GET", "/ping", pong)

	ts := httptest.NewServer(r)
	defer ts.Close()

	{
		res, err := http.Get(fmt.Sprintf("%s/ping", ts.URL))
		if err != nil {
			log.Println(err)
		}
		resp, _ := ioutil.ReadAll(res.Body)
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, "pong", string(resp))
	}
}
