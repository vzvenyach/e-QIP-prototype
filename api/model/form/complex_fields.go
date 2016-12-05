package form

// AddressField contains complete address information for an entity
type AddressField struct {
	Address TextField
	Street  TextField
	City    TextField
	State   StateField
	Zipcode ZipcodeField
	County  TextField
	Country CountryField
}

func (a AddressField) Valid() (bool, error) {

	var stack ErrorStack

	if ok, err := a.Address.Valid(); !ok {
		stack.Append("Address", err)
	}

	if ok, err := a.Street.Valid(); !ok {
		stack.Append("Street", err)
	}

	if ok, err := a.City.Valid(); !ok {
		stack.Append("City", err)
	}

	if ok, err := a.State.Valid(); !ok {
		stack.Append("State", err)
	}

	if ok, err := a.Zipcode.Valid(); !ok {
		stack.Append("Zipcode", err)
	}

	if ok, err := a.County.Valid(); !ok {
		stack.Append("County", err)
	}

	if ok, err := a.Country.Valid(); !ok {
		stack.Append("Country", err)
	}

	return !stack.HasErrors(), stack
}

// BirthAddress contains the birth location for a person. This contains a
// subset of the Address information
type BirthAddressField struct {
	City    TextField
	State   StateField
	County  CountryField
	Country CountryField
}

// Valid validates information for a person Birth Address
func (f BirthAddressField) Valid() (bool, error) {
	var stack ErrorStack
	stack.Append("City", ErrFieldRequired{"Field is required"})
	stack.Append("State", ErrFieldRequired{"State is required"})

	return !stack.HasErrors(), stack
}

// PersonField stores information for a particular person
type PersonField struct {
	Firstname  FirstnameField
	Lastname   TextField
	Middlename MiddlenameField
	Suffix     Suffix
}

func (f PersonField) Valid() (bool, error) {
	var stack ErrorStack

	if ok, err := f.Firstname.Valid(); !ok {
		stack.Append("Firstname", err)
	}

	if ok, err := f.Lastname.Valid(); !ok {
		stack.Append("Lastname", err)
	}

	if ok, err := f.Middlename.Valid(); !ok {
		stack.Append("Middlename", err)
	}
	return !stack.HasErrors(), stack
}

// DateRangeField contains from and to dates and ensures they are within a valid range
type DateRangeField struct {
	From DateField
	To   DateField
}

func (d DateRangeField) Valid() (bool, error) {
	return true, nil
}

// EmployerActivity stores when a person has worked during a specific period of time for a particular
// position and supervisor
type EmployerActivity struct {
	PositionTitle     TextField
	Supervisor        TextField
	DatesOfEmployment DateRangeField
}

type EmployerActivities []EmployerActivity
