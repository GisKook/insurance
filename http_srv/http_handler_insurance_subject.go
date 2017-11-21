package http_srv

import (
	"fmt"
	"github.com/giskook/insurance/base"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

const (
	INSURANCE_SUBJECT_PAGE_SIZE       int    = 20
	HTTP_INSURANCE_SUBJECT_NAME       string = "name"
	HTTP_INSURANCE_SUBJECT_PAGE_INDEX string = "page_index"
)

type subject_list struct {
	Subject []*base.Subject
	PageSize   int
	Total      int
	SearchName string
}

func (h *HttpSrv) handler_insurance_subject(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if x := recover(); x != nil {
			log.Println(x)
			fmt.Fprint(w, EncodeGeneralResponse(HTTP_REP_INTERAL_ERROR))
		}
	}()

	RecordReq(r)
	r.ParseForm()
	defer r.Body.Close()

	tmpl := template.Must(template.ParseFiles("./web/tmpl/insurance/insurance_subject.tmpl", "./web/tmpl/insurance/insurance_subject_dlg_add.tmpl"))

	name := r.Form.Get(HTTP_INSURANCE_SUBJECT_NAME)
	page_index := r.Form.Get(HTTP_INSURANCE_SUBJECT_PAGE_INDEX)

	pgi, _ := strconv.Atoi(page_index)
	if pgi >= 1 {
		pgi -= 1
	}
	offset := pgi * INSURANCE_SUBJECT_PAGE_SIZE

	wg := &sync.WaitGroup{}
	wg.Add(2)
	var s []*base.Subject
	var err, err2 error
	var count int
	go func(wg *sync.WaitGroup) {
		s, err = h.db.InsuranceSubjectGet(INSURANCE_SUBJECT_PAGE_SIZE, offset, name)
		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		count, err2 = h.db.InsuranceSubjectGetCount()
		wg.Done()
	}(wg)

	wg.Wait()

	if err != nil {
		panic(err)
	}
	if err2 != nil {
		panic(err2)
	} 
	sl := &subject_list{ 
		Subject:s,
		PageSize:INSURANCE_SUBJECT_PAGE_SIZE,
		Total:count, 
		SearchName:name,
	}

	err = tmpl.Execute(w,sl)
	panic(err)
}