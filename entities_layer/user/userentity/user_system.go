package userentity

// UserSystem is worker of organisation
type UserSystem struct {
	ID          string
	IDCl        string
	IDsEst      []string
	Login       string
	Password    string
	Email       string
	PhoneNumber string
	Name        string
	Surname     string
	Patronymic  string
	Roles       []string
}
