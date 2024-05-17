package presenters

type GetListActivitiesRequest struct {
	Service string `query:"service"`
	Date    string `query:"date" validate:"required"`
	Page    int    `query:"page"`
	Limit   int    `query:"limit"`
}
