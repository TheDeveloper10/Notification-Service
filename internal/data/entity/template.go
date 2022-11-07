package entity

type TemplateBody struct {
	Email *string `json:"email"`
	SMS   *string `json:"sms"`
	Push  *string `json:"push"`
}

type TemplateEntity struct {
	Id       int          `json:"id"`
	Body     TemplateBody `json:"contents"`
	Language string       `json:"language"`
	Type     string 	  `json:"type"`
}
