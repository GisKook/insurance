package http_srv

import (
	"fmt"
	"github.com/giskook/insurance/base"
	"github.com/giskook/insurance/db"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

const (
	PAGE_SIZE       int    = 20
	HTTP_NAME       string = "name"
	HTTP_ID         string = "id"
	HTTP_TEL        string = "tel"
	HTTP_PAGE_INDEX string = "page_index"
)

type user_info struct {
	ID   string
	Name string
	Tel  string
	Addr string
}

type user_list struct {
	User       []*user_info
	PageSize   int
	Total      int
	SearchName string
	SearchID   string
	SearchTel  string
}

func convert_user_info(base_user []*base.User) []*user_info {
	list := make([]*user_info, 0)
	for _, v := range base_user {
		list = append(list, &user_info{
			ID:   v.ID,
			Name: v.Name,
			Tel:  v.Tel,
			Addr: v.Prov + v.City + v.County,
		})
	}

	return list
}

func (h *HttpSrv) handler_user_show_panel(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if x := recover(); x != nil {
			log.Println(x)
			fmt.Fprint(w, EncodeGeneralResponse(HTTP_REP_INTERAL_ERROR))
		}
	}()

	RecordReq(r)
	r.ParseForm()
	defer r.Body.Close()

	tmpl := template.Must(template.ParseFiles("./web/tmpl/user/user.tmpl","./web/tmpl/user/user_dlg_add.tmpl"))

	name := r.Form.Get(HTTP_NAME)
	id := r.Form.Get(HTTP_ID)
	tel := r.Form.Get(HTTP_TEL)
	page_index := r.Form.Get(HTTP_PAGE_INDEX)

	pgi, _ := strconv.Atoi(page_index)
	if pgi >= 1 {
		pgi -= 1
	}
	offset := pgi * PAGE_SIZE

	wc := h.db.User_gen_where_clause(name, id, tel)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	var u []*base.User
	var err, err2 error
	var count int
	go func(wg *sync.WaitGroup) {
		u, err = h.db.UserGet(PAGE_SIZE, offset, wc)
		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		count, err2 = h.db.CommonGetTableRowCount(db.TABLE_USER, wc)
		wg.Done()
	}(wg)

	wg.Wait()

	if err != nil {
		panic(err)
	}
	if err2 != nil {
		panic(err2)
	}
	user := convert_user_info(u)

	uu := user_list{
		User:       user,
		PageSize:   PAGE_SIZE,
		Total:      count,
		SearchName: name,
		SearchID:   id,
		SearchTel:  tel,
	}

	err = tmpl.Execute(w, uu)
	panic(err)
}
