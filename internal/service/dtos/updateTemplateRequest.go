package dtos

type UpdateTemplateRequest struct {
	AbstractRequest
	Id 		 *int `json:"id"`
	Template *string `json:"template"`
}

func (utr *UpdateTemplateRequest) Validate() (bool, string) {
	if utr.Id == nil {
		return false, "'id' must be given!"
	} else if utr.Template == nil {
		return false, "'template' must be given!"
	} else if (*utr.Id) <= 0 {
		return false, "'id' must be greater than 0"
	} else if len(*utr.Template) <= 0 || len(*utr.Template) > 2048 {
		return false, "'template' must have a length greater than 0 and lesser than 2048!"
	}
	return true, ""
}
