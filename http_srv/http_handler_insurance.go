package http_srv
import(
	"github.com/gorilla/mux"
)

func (h *HttpSrv) init_web_insurance(r *mux.Router) {
	s := r.PathPrefix("/insurance").Subrouter()
	s.HandleFunc("/undertake", h.handler_insurance_undertake)
}