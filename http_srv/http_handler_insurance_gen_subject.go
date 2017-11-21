package http_srv

import (
	"fmt"
	"log"
	"net/http"
)

func (h *HttpSrv) handler_insurance_gen_subject_div(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if x := recover(); x != nil {
			log.Println(x)
			fmt.Fprint(w, EncodeGeneralResponse(HTTP_REP_INTERAL_ERROR))
		}
	}()

	RecordReq(r)
	r.ParseForm()
	defer r.Body.Close()
}
