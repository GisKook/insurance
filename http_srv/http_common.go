package http_srv

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

type ctx_key int

type Code int

const (
	HTTP_REP_SUCCESS        Code    = 0
	HTTP_REP_LACK_PARAMETER Code    = 1
	HTTP_REP_INTERAL_ERROR  Code    = 2
	xkey                    ctx_key = 0
)

var HTTP_REQUEST_DESC []string = []string{
	"成功",
	"缺少参数",
	"服务器内部错误"}

func (c Code) Desc() string {
	return HTTP_REQUEST_DESC[c]
}

type GeneralResponse struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
}

func EncodeGeneralResponse(code Code) string {
	general_response := &GeneralResponse{
		Code: int(code),
		Desc: code.Desc(),
	}

	resp, _ := json.Marshal(general_response)

	return string(resp)
}

func EncodeErrResponse(code int, desc string) string{
	gr := &GeneralResponse{
		Code:code,
		Desc:desc,
	}

	resp, _ := json.Marshal(gr)

	return string(resp)
}

func RecordReq(r *http.Request) {
	v, e := httputil.DumpRequest(r, true)
	if e != nil {
		log.Println(e.Error())
		return
	}
	log.Println(string(v))
}

func (h *HttpSrv) show_panel(panel string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				fmt.Fprint(w, EncodeGeneralResponse(HTTP_REP_INTERAL_ERROR))
			}
		}()

		RecordReq(r)
		tmpl := template.Must(template.ParseFiles(panel))

		tmpl.Execute(w, nil)

	})
}

func (h *HttpSrv) validate(page http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie("Auth")
		if err != nil {
			http.NotFound(res, req)
			return
		}

		token, err := jwt.ParseWithClaims(cookie.Value, &JwtAuth{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method")
			}
			return []byte(h.conf.Http.Secret), nil
		})
		if err != nil {
			http.NotFound(res, req)
			return
		}
		if claims, ok := token.Claims.(*JwtAuth); ok && token.Valid {
			ctx := context.WithValue(req.Context(), xkey, *claims)
			page(res, req.WithContext(ctx))
		} else {
			http.NotFound(res, req)
			return
		}
	})
}

// MarshalJson 把对象以json格式放到response中
func marshal_json(w http.ResponseWriter, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	fmt.Fprint(w, string(data))
	return nil
}

// UnMarshalJson 从request中取出对象
func unmarshal_json(req *http.Request, v interface{}) error {
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	json.Unmarshal([]byte(bytes.NewBuffer(result).String()), v)
	return nil
}
