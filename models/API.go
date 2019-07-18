package models

type API struct {
	URL string
	Token string
	Request Request
	Response Response
	Query Query
}

// Select inicializa a construção da query
func (a *API) NewRequest(fd []string, ft []string) {
	a.Query.New(fd, ft)
	a.Request.New(a.URL + "?token=" + a.Token + a.Query.GetQuery(), "GET")
}

func (a *API) GetRequestURL() string {
	return a.Request.URL
}