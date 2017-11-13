package base

type UserListRep struct {
	Code    int    `json:"code"`
	Desc    string `json:"desc"`
	Content struct {
		User []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			Tel  string `json:"tel"`
			Addr string `json:"addr"`
		} `json:"user"`
	} `json:"content"`
}

type User struct {
	ID     string
	Name   string
	Tel    string
	Prov   string
	City   string
	County string
}
