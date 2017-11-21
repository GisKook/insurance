package http_srv
import(
	"github.com/gorilla/mux"
)

func (h *HttpSrv) init_web_insurance(r *mux.Router) {
	s := r.PathPrefix("/insurance").Subrouter()
	s.HandleFunc("/undertake", h.handler_insurance_undertake)
	s.HandleFunc("/gen_subject", h.handler_insurance_gen_subject_div)
	s.HandleFunc("/subject", h.handler_insurance_subject)
	s.HandleFunc("/subject_add",h.handler_insurance_subject_add)
}