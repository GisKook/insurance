package http_srv

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func (h *HttpSrv) handler_insurance_undertake(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if x := recover(); x != nil {
			log.Println(x)
			fmt.Fprint(w, EncodeGeneralResponse(HTTP_REP_INTERAL_ERROR))
		}
	}()

	RecordReq(r)
	r.ParseForm()
	defer r.Body.Close()

	tmpl := template.Must(template.ParseFiles("./web/tmpl/insurance/insurance.tmpl","./web/tmpl/insurance/insurance_insured.tmpl"))

	err := tmpl.Execute(w,nil) 
	panic(err)
}
