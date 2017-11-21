
package http_srv

import (
	"github.com/giskook/insurance/base"
	"fmt"
	"net/http"
	"log"
)

const (
	HTTP_INSURANCE_SUBJECT_ADD_NAME string = "name"
	HTTP_INSURANCE_SUBJECT_ADD_ATTR string = "attr"
)

func (h *HttpSrv) handler_insurance_subject_add(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if x := recover(); x != nil {
			log.Println(x)
			fmt.Fprint(w, EncodeGeneralResponse(HTTP_REP_INTERAL_ERROR))
		}
	}()

	RecordReq(r)
	r.ParseForm()
	defer r.Body.Close()

	name := r.Form.Get(HTTP_INSURANCE_SUBJECT_ADD_NAME)
	attrs := r.Form[HTTP_INSURANCE_SUBJECT_ADD_ATTR]
	if name == ""{
		fmt.Fprint(w, EncodeGeneralResponse(HTTP_REP_LACK_PARAMETER))
		return 
	}

	for index, attr := range attrs{
		if attr == "" {
			if len(attrs) != 1{
				attrs = append(attrs[:index], attrs[index+1:]...)
			}else{
				attrs = attrs[:0]
			}
		}
	}
	if len(attrs) == 0{
		fmt.Fprint(w, EncodeErrResponse(base.ERR_INSURANCE_SUBJECT_ADD_NO_ATTR_CODE, base.ERR_INSURANCE_SUBJECT_ADD_NO_ATTR_DESC))
		return
	}

	err := h.db.InsuranceSubjectAdd(&base.Subject{
		Name:name,
		Cols:attrs,
	})
	if err != nil{
		fmt.Fprint(w, EncodeErrResponse(err.(*base.PiccError).Code, err.(*base.PiccError).Desc))
		return
	}
	fmt.Fprint(w,EncodeGeneralResponse(HTTP_REP_SUCCESS))
}