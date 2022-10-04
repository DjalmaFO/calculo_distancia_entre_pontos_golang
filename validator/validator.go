package validator

type Validator struct {
	Valor    interface{}
	Esperado bool
}

func Validate(validate []Validator, f func(interface{}) bool) (ok bool) {
	for _, v := range validate {
		if v.Esperado != f(v.Valor) {
			return
		}
	}
	return true
}
