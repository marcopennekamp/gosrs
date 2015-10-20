package handlers
import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/marcopennekamp/gosrs/core"
)


func Index(ctx *core.Context, w http.ResponseWriter, r *http.Request, _ httprouter.Params) (core.HttpStatus, error) {
	w.Write([]byte("おはよう！"))
	return http.StatusOK, nil
}
