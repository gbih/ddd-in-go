package valueobject001

type name struct {
	firstname string
	lastname  string
}

type address struct {
	streetaddress string
	city          string
	zip           string
}

var Name1 = name{
	firstname: "Alex",
	lastname:  "Adams",
}

var Name2 = name{
	firstname: "Alex",
	lastname:  "Adams",
}

var Name3 = name{
	firstname: "John",
	lastname:  "Smith",
}

var Address1 = address{
	streetaddress: "123 Main St",
	city:          "New York",
	zip:           "90001",
}

var Address2 = address{
	streetaddress: "123 Main St",
	city:          "New York",
	zip:           "90001",
}

type phone interface {
	isphone()
}
type phoneCompany struct{ model string }
type phonePersonal struct{ model string }

func (o phoneCompany) isphone()  {}
func (o phonePersonal) isphone() {}

var (
	_ phone = (*phoneCompany)(nil)
	_ phone = (*phonePersonal)(nil)
)

var Phone1 = phoneCompany{model: "AAA"}
var Phone2 = phoneCompany{model: "AAA"}
var Phone3 = phoneCompany{model: "BBB"}
var Phone4 = phonePersonal{model: "BBB"}
