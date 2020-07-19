package limit

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
	"asap/router"
	"github.com/stretchr/testify/assert"

)

func TestRandomReject(t *testing.T) {
	r := gin.Default()
	router.InitRouter(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/limit/randomreject", nil)
	r.ServeHTTP(w, req)

	if w.Code == http.StatusOK {
		assert.Equal(t, "ok", w.Body.String())
	}else if w.Code == http.StatusBadGateway{
		assert.Equal(t, "reject", w.Body.String())
	}
}
