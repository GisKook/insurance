package http_srv

import (
	"github.com/giskook/insurance/base"
	"fmt"
	//"html/template"
	"log"
	"net/http"
	"crypto/md5"
	"io"
)

const (
	HTTP_USER_ADD_NAME       string = "name"
	HTTP_USER_ADD_ID         string = "id"
	HTTP_USER_ADD_TEL        string = "tel"
	HTTP_USER_ADD_PROV        string = "prov"
	HTTP_USER_ADD_CITY        string = "city"
	HTTP_USER_ADD_COUNTY     string = "county"
	HTTP_USER_ADD_POSTCODE     string = "post_code"
	HTTP_USER_ADD_SUPER     string = "super"
	HTTP_USER_ADD_UNDERTAKE     string = "undertake"
	HTTP_USER_ADD_VERIFICATION string = "verification"
	HTTP_USER_ADD_LOSS string = "loss"
)

func (h *HttpSrv) handler_user_add(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if x := recover(); x != nil {
			log.Println(x)
			fmt.Fprint(w, EncodeGeneralResponse(HTTP_REP_INTERAL_ERROR))
		}
	}()

	RecordReq(r)
	r.ParseForm()
	defer r.Body.Close()

//	tmpl := template.Must(template.ParseFiles("./web/tmpl/user/user_dlg_add.tmpl"))
//	log.Println(tmpl)

	name := r.Form.Get(HTTP_USER_ADD_NAME)
	id := r.Form.Get(HTTP_USER_ADD_ID)
	tel := r.Form.Get(HTTP_USER_ADD_TEL)
	prov := r.Form.Get(HTTP_USER_ADD_PROV)
	city := r.Form.Get(HTTP_USER_ADD_CITY)
	county := r.Form.Get(HTTP_USER_ADD_COUNTY)
	post_code := r.Form.Get(HTTP_USER_ADD_POSTCODE)

	role_super:= r.Form.Get(HTTP_USER_ADD_SUPER)
	role_undertake := r.Form.Get(HTTP_USER_ADD_UNDERTAKE)
	role_verification := r.Form.Get(HTTP_USER_ADD_VERIFICATION)
	role_loss := r.Form.Get(HTTP_USER_ADD_LOSS)

	if name == ""||
	id == ""||
	tel == ""||
	prov == ""||
	city == ""||
	county == ""||
	post_code == ""{ 
		fmt.Fprint(w, EncodeGeneralResponse(HTTP_REP_LACK_PARAMETER))
		return 
	}
	if role_super == "" &&
	role_undertake == "" &&
	role_verification == ""&&
	role_loss == ""{
		fmt.Fprint(w, EncodeGeneralResponse(HTTP_REP_LACK_PARAMETER))
		return 
	}
	roles := make([]int, 0)
	if role_super != ""{
		roles = append(roles, base.DB_USER_ROLE_SUPER)
	}else{
		if role_undertake != ""{
			roles = append(roles, base.DB_USER_ROLE_UNDERTAKE)
		}
		if role_verification != ""{
			roles = append(roles, base.DB_USER_ROLE_VERIFICATION)
		}
		if role_loss != ""{
			roles = append(roles, base.DB_USER_ROLE_LOSS)
		}
	}

	m := md5.New()
	io.WriteString(m, id)

	passwd := fmt.Sprintf("%x", m.Sum([]byte(h.conf.DB.Secret)))

	err := h.db.UserAdd(&base.User{
		ID:id, 
		Name:name,
		Tel:tel,
		Prov:prov,
		City:city,
		County:county,
		PostCode:post_code,
		Roles:roles,
		Passwd:passwd,
	})

	if err != nil{
		fmt.Fprint(w, EncodeErrResponse(err.(*base.PiccError).Code, err.(*base.PiccError).Desc))
		return
	}

	fmt.Fprint(w,EncodeGeneralResponse(HTTP_REP_SUCCESS))
}
