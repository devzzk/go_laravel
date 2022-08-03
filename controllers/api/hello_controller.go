package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HelloController struct {
}

func (hc HelloController) index(w *gin.ResponseWriter, r *http.Request) {

}
