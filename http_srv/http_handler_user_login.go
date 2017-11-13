package http_srv

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type JwtAuth struct {
	Auth   string `json:"au"`
	Userid string `json:"id"`
	jwt.StandardClaims
}

func (h *HttpSrv) init_user_api(r *mux.Router) {
	s := r.PathPrefix("/user").Subrouter()
	s.HandleFunc("/login", h.handler_user_login)
	s.HandleFunc("/logout", h.handler_user_logout)
	s.HandleFunc("/profile", h.validate(h.handler_user_profile))
	s.HandleFunc("/list", h.handler_user_list)
	s.HandleFunc("/show_panel", h.handler_user_show_panel)
}

func (h *HttpSrv) gen_jwt_token(auth string, id string) (string, error) {
	claims := JwtAuth{
		auth,
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(h.conf.Http.Expire)).Unix(),
			Issuer:    "zk",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(h.conf.Http.Secret))
}

func (h *HttpSrv) set_token(res http.ResponseWriter, req *http.Request, auth string, userid string) error {
	token, err := h.gen_jwt_token(auth, userid)
	if err != nil {
		return err
	}
	exp_cookie_time := time.Now().Add(time.Second * time.Duration(h.conf.Http.Expire))
	cookie := http.Cookie{Name: "Auth", Value: token, Expires: exp_cookie_time, HttpOnly: true}
	http.SetCookie(res, &cookie)

	http.Redirect(res, req, "/api/user/profile", 307)

	return nil
}

func (h *HttpSrv) handler_user_login(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Fprint(w, EncodeGeneralResponse(HTTP_REP_INTERAL_ERROR))
		}
	}()

	RecordReq(r)
	r.ParseForm()
	defer r.Body.Close()

	user := r.Form.Get("user")
	passwd := r.Form.Get("passwd")
	//var login Login
	if user == "" ||
		passwd == "" {
		fmt.Fprint(w, EncodeGeneralResponse(HTTP_REP_LACK_PARAMETER))
		return
	}

	auth, id, err := h.db.UserValid(user, passwd)
	if err != nil {
		log.Println(err.Error())
		return
	}

	err = h.set_token(w, r, auth, id)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func (h *HttpSrv) handler_user_logout(w http.ResponseWriter, r *http.Request) {
	RecordReq(r)
	deleteCookie := http.Cookie{Name: "Auth", Value: "none", Expires: time.Now()}
	http.SetCookie(w, &deleteCookie)
	return
}
