package view

import (
	"log"
	"net/http"

	"go.sancus.dev/cms"
)

func (v *View) HandleResourceError(w http.ResponseWriter, r *http.Request, err cms.Error) {
	log.Printf("%T.HandleResourceError: %s%s: %v: %s", v, r.Host, r.URL.Path, err.Status(), err.Error())
	panic(err)
}

func (v *View) HandleError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("%T.HandleError: %s%s: %T: %s", v, r.Host, r.URL.Path, err, err.Error())
	panic(err)
}

func (v *View) HandlePanic(w http.ResponseWriter, r *http.Request, rvr interface{}) {
	log.Printf("%T.HandlePanic: %s%s: %T", v, r.Host, r.URL.Path, rvr)
}
