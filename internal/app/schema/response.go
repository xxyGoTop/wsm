package schema

type Meta struct {
	Limit 	int `json:"limit"`
	Page 	int `json:"page"`
	Total 	int `json:"total"`
	Num 	int `json:"num"`
	Sort 	int `json:"sort"`
	Platform *string `json:"platform,omitempty"`
}

type Response struct {
	Message string 		`json:"message"`
	Data 	interface{} `json:"data"`
	Status 	int 		`json:"status"`
	Meta 	*Meta 		`json:"meta"`
}

const (
	StatusSuccess = 1
	StatusFail = 0
)