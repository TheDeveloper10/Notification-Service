package dtos

type TemplateIdRequest struct {
	AbstractRequest
	Id *int `json:"id"`
}

func (ctr *TemplateIdRequest) Validate() (bool, string) {
	if ctr.Id == nil {
		return false, "'id' must be given!"
	} else if (*ctr.Id) <= 0 {
		return false, "'id' must be greater than 0"
	}
	return true, ""
}
