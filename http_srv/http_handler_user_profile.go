package http_srv

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

const (
	UNDERTAKE       string = "ud"
	VERIFICATION    string = "ve"
	LOSS            string = "lo"
	STATISTICS      string = "st"
	LARGE_AREA_LOSS string = "ll"
	USER_MANAGEMENT string = "ad"
	SUBJECT         string = "su"
	POLICY          string = "pm"
)

type business_management struct {
	Undertake    string
	Verification string
	Loss         string
}

type statistic_management struct {
	Statistic string
	LargeLoss string
}

type system_management struct {
	UserManagement string
	Subject        string
}

type mine struct {
	Setting   string
	Insurance string
}

type header struct {
	Name string
	BM   *business_management
	SM   *statistic_management
	SYM  *system_management
	Mine *mine
}

type profile struct {
	Title  string
	Header *header
}

func (h *HttpSrv) gen_menu(name string, auth string) *header {
	head := new(header)
	head.Name = name
	if strings.Contains(auth, UNDERTAKE) {
		head.BM = new(business_management)
		head.BM.Undertake = UNDERTAKE
	}
	if strings.Contains(auth, VERIFICATION) {
		if head.BM == nil {
			head.BM = new(business_management)
		}
		head.BM.Verification = VERIFICATION
	}
	if strings.Contains(auth, LOSS) {
		if head.BM == nil {
			head.BM = new(business_management)
		}
		head.BM.Loss = LOSS
	}

	if strings.Contains(auth, STATISTICS) {
		head.SM = new(statistic_management)
		head.SM.Statistic = STATISTICS
	}
	if strings.Contains(auth, LARGE_AREA_LOSS) {
		if head.SM == nil {
			head.SM = new(statistic_management)
		}
		head.SM.LargeLoss = LARGE_AREA_LOSS
	}

	if strings.Contains(auth, USER_MANAGEMENT) {
		head.SYM = new(system_management)
		head.SYM.UserManagement = USER_MANAGEMENT
	}
	if strings.Contains(auth, SUBJECT) {
		if head.SYM == nil {
			head.SYM = new(system_management)
		}
		head.SYM.Subject = SUBJECT
	}

	head.Mine = new(mine)
	head.Mine.Setting = "mine"
	head.Mine.Insurance = "insurane"

	return head
}

func (h *HttpSrv) handler_user_profile(w http.ResponseWriter, r *http.Request) {
	RecordReq(r)
	token, ok := r.Context().Value(xkey).(JwtAuth)
	if !ok {
		http.NotFound(w, r)
		return
	}
	name, err := h.db.UserGetInfo(token.Userid)
	if err != nil {
		log.Println(err.Error())
	}
	p := profile{Title: name,
		Header: h.gen_menu(name, token.Auth),
	}
	tmpl := template.Must(template.ParseFiles("./web/tmpl/main.tmpl", "./web/tmpl/common/header.tmpl"))
	err = tmpl.Execute(w, p)
	if err != nil {
		panic(err)
	}
}
