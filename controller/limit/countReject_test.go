package limit

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
	"asap/router"
	"github.com/stretchr/testify/assert"
)

func TestCountReject(t *testing.T) {
	r := gin.Default()
	router.InitRouter(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/limit/countreject", nil)
	r.Run(":8082")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
