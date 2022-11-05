package util

type ListOfStrings []string

func (los *ListOfStrings) Has(str string) bool {
	for _, s := range *los {
		if s == str {
			return true
		}
	}
	return false
}

func (los *ListOfStrings) Join(op string) string {
	res := ""
	length := len(*los)
	if length == 0 {
		return ""
	} else if length == 1 {
		return (*los)[0]
	}

	for i := 0; i < length - 1; i++ {
		res += (*los)[i] + op
	}
	res += (*los)[length - 1]

	return res
}