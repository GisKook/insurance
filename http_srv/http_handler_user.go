package http_srv

import (
	//"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type UserListReq struct {
	Page  *int
	Count *int
}

func (h *HttpSrv) check_user_list_req_paramters(paramters *UserListReq) bool {
	if paramters.Page == nil || paramters.Count == nil {
		return false
	}

	return true
}

func (h *HttpSrv) handler_user_list(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Fprint(w, EncodeGeneralResponse(HTTP_REP_INTERAL_ERROR))
		}
	}()

	RecordReq(r)
	//r.ParseForm()
	user_list_req := &UserListReq{}
	err := unmarshal_json(r, user_list_req)
	if err != nil {
		log.Println(err)
		return
	}
	if !h.check_user_list_req_paramters(user_list_req) {
		marshal_json(w, &GeneralResponse{Code: int(HTTP_REP_LACK_PARAMETER), Desc: HTTP_REP_LACK_PARAMETER.Desc()})
		return
	}
}
