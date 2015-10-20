package core
import (
	"github.com/jmoiron/sqlx"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"strconv"
)

type Context struct {
	Conf		*Config
	Db			*sqlx.DB
	Templates	*TemplateLoader
}

type HttpStatus int

type AppHandle func(*Context, http.ResponseWriter, *http.Request, httprouter.Params) (HttpStatus, error)

func CreateHandle(ctx *Context, ah AppHandle) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		status, err := ah(ctx, w, r, ps)
		statusValue := int(status)
		if err != nil {
			println("HTTP", strconv.Itoa(statusValue) + ":", err.Error())
			switch statusValue {
			case http.StatusNotFound:
				http.NotFound(w, r)
			case http.StatusInternalServerError:
				http.Error(w, http.StatusText(statusValue), statusValue)
			default:
				http.Error(w, http.StatusText(statusValue), statusValue)
			}
		}
	}
}
