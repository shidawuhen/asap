package limit

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
	"asap/router"
)

func BenchmarkSlide(b *testing.B){
	r := gin.Default()
	router.InitRouter(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/limit/slidewindowsreject", nil)

	for i:=0; i< b.N;i++{
		r.ServeHTTP(w, req)
	}
}
