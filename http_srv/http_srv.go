package http_srv

import (
	"github.com/giskook/insurance/conf"
	"github.com/giskook/insurance/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type HttpSrv struct {
	conf   *conf.Conf
	router *mux.Router
	db     *db.DBSocket
}

func NewHttpSrv(conf *conf.Conf) *HttpSrv {
	db, err := db.NewDBSocket(conf.DB)
	if err != nil {
		log.Println(err.Error())

		return nil
	}
	return &HttpSrv{
		conf:   conf,
		router: mux.NewRouter(),
		db:     db,
	}
}

func (h *HttpSrv) Start() {
	s := h.router.PathPrefix("/web").Subrouter()
	h.init_web_user(s)
	h.init_web_insurance(s)

	h.router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(h.conf.Http.Path))))

	if err := http.ListenAndServe(h.conf.Http.Addr, h.router); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
