package types

type User struct {
	ID         int     `json:"id" validate:"required"`
	FirstName  string  `json:"firstName"`
	LastName   string  `json:"lastName"`
	MaidenName string  `json:"maidenName"`
	Age        int     `json:"age"`
	Gender     string  `json:"gender"`
	Email      string  `json:"email" validate:"required,min=1,max=100"`
	Phone      string  `json:"phone"`
	Username   string  `json:"username"`
	Password   string  `json:"password"`
	BirthDate  string  `json:"birthDate"`
	Image      string  `json:"image"`
	BloodGroup string  `json:"bloodGroup"`
	Height     float64 `json:"height"`
	Weight     float64 `json:"weight"`
	EyeColor   string  `json:"eyeColor"`
	Hair       Hair    `json:"hair"`
	IP         string  `json:"ip"`
	Address    Address `json:"address"`
	MacAddress string  `json:"macAddress"`
	University string  `json:"university"`
	Bank       Bank    `json:"bank"`
	Company    Company `json:"company"`
	EIN        string  `json:"ein"`
	SSN        string  `json:"ssn"`
	UserAgent  string  `json:"userAgent"`
	Crypto     Crypto  `json:"crypto"`
	Role       string  `json:"role"`
	Comment    Comment `json:"comment"`
}

type Hair struct {
	Color string `json:"color"`
	Type  string `json:"type"`
}

type Address struct {
	Address     string      `json:"address"`
	City        string      `json:"city"`
	State       string      `json:"state"`
	StateCode   string      `json:"stateCode"`
	PostalCode  string      `json:"postalCode"`
	Coordinates Coordinates `json:"coordinates"`
	Country     string      `json:"country"`
}

type Coordinates struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Bank struct {
	CardExpire string `json:"cardExpire"`
	CardNumber string `json:"cardNumber"`
	CardType   string `json:"cardType"`
	Currency   string `json:"currency"`
	IBAN       string `json:"iban"`
}

type Company struct {
	Department string  `json:"department"`
	Name       string  `json:"name"`
	Title      string  `json:"title"`
	Address    Address `json:"address"`
}

type Crypto struct {
	Coin    string `json:"coin"`
	Wallet  string `json:"wallet"`
	Network string `json:"network"`
}
