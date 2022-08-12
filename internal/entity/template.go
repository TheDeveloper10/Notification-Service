package entity

type TemplateEntity struct {
	Id 			int    `json:"id"`
	ContactType string `json:"contactType"`
	Template 	string `json:"template"`
	Language    string `json:"language"`
	Type        string `json:"type"`
}

func (te *TemplateEntity) GetRespectiveContactInfoType() string {
	switch te.ContactType {
	case ContactTypeEmail:
		return "email"
	case ContactTypePush:
		return "fcmRegistrationToken"
	case ContactTypeSMS:
		return "phoneNumber"
	default:
		return ""
	}
}
