package verificationservice

type CustomerEmail interface {
	isCustomerEmail()
}

type Unverified struct {
	Email string
}

type VerifiedEmailAddress struct {
	email string
}

type Verified VerifiedEmailAddress

func (u Unverified) isCustomerEmail() {}
func (v Verified) isCustomerEmail()   {}

var (
	_ CustomerEmail = (*Unverified)(nil)
	_ CustomerEmail = (*Verified)(nil)
)

func CheckVerification(i string) Verified {
	// do check here
	// assume it passes
	if i == "jiko@omame.com" {
		var VEA = Verified{
			email: i,
		}
		return VEA
	}
	return Verified{}
}
