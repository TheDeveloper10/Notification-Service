package dto

type TemplateRecord struct {
	Id int
	ContactType int8
	Template string
}

type TemplateReadable struct {
	Id int `json:"id"`
	ContactType string `json:"contactType"`
	Template string `json:"template"`
}

func (rec *TemplateRecord) ToReadable() (TemplateReadable) {
	contactType := ""

	// make a repository for templates
	switch (rec.ContactType) {
	case 1:
		contactType = "email"
	case 2:
		contactType = "sms"
	case 3:
		contactType = "push"
	}

	return TemplateReadable{
		Id: rec.Id,
		ContactType: contactType,
		Template: rec.Template,
	}
}