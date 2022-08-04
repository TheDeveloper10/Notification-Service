package dto

type TemplateRecord struct {
	Id int
	ContactType string
	Template string
}

func validateContactType(contactType *string) bool {
	return *contactType == "email" || *contactType == "sms" || *contactType == "push"
}