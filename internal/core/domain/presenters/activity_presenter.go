package presenters

type GetListActivitiesRequest struct {
	Service string `query:"service"`
	Created string `query:"created"`
	Page    int    `query:"page"`
	Limit   int    `query:"limit"`
}
