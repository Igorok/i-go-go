package serviceentity

// Delivery is the entity of delivery
type Delivery struct {
	ID         string
	IDCl       string
	IDEst      string
	IDUsr      string
	IDCst      string
	IDsOrd     []string
	Address    string
	Status     string
	DateStart  string
	DateFinish string
	DateCreate string
}
