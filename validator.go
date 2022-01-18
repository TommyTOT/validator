package validator

// Form parameter form
type Form interface {
	// Validate form validation, return error message
	Validate() (err error)
}

// Validate implement the `Validate` used by gin `validate.Validate`
type Validate struct{}

// New implement the `New()` used by gin `validate.New()`
func New() (validate *Validate) { return }

// SetTagName implement the `SetTagName(_ string)` used by gin `validate.Validate.SetTagName(name string)`
func (*Validate) SetTagName(_ string) {}

// Struct implement the `Struct(form interface{})` used by gin `validate.Validate.Struct(s interface)`
func (*Validate) Struct(form interface{}) (err error) {
	// judge whether it is not `Form`
	f, status := form.(Form)
	if status {
		// call validation
		err = f.Validate()
	} else {

	}
	return
}
