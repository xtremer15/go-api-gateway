package types

type User struct {
	ID         int      `json:"id" validate:"required"`
	FirstName  string   `json:"firstName"`
	LastName   string   `json:"lastName"`
	MaidenName string   `json:"maidenName"`
	Age        int      `json:"age"`
	Gender     string   `json:"gender"`
	Email      string   `json:"email" validate:"required,min=1,max=100"`
	Phone      string   `json:"phone"`
	Username   string   `json:"username"`
	Password   string   `json:"password"`
	BirthDate  string   `json:"birthDate"`
	Image      string   `json:"image"`
	BloodGroup string   `json:"bloodGroup"`
	Height     float64  `json:"height"`
	Weight     float64  `json:"weight"`
	EyeColor   string   `json:"eyeColor"`
	Hair       hair     `json:"hair"`
	IP         string   `json:"ip"`
	Address    address  `json:"address"`
	MacAddress string   `json:"macAddress"`
	University string   `json:"university"`
	Bank       bank     `json:"bank"`
	Company    company  `json:"company"`
	EIN        string   `json:"ein"`
	SSN        string   `json:"ssn"`
	UserAgent  string   `json:"userAgent"`
	Crypto     crypto   `json:"crypto"`
	Role       string   `json:"role"`
	Posts      []Post   `json:"posts"`
	Cart       Cart     `json:"cart"`
	Quotes     []Quote  `json:"quotes"`
	Recipes    []Recipe `json:"recipes"`
}

type hair struct {
	Color string `json:"color"`
	Type  string `json:"type"`
}

type address struct {
	Address     string      `json:"address"`
	City        string      `json:"city"`
	State       string      `json:"state"`
	StateCode   string      `json:"stateCode"`
	PostalCode  string      `json:"postalCode"`
	Coordinates coordinates `json:"coordinates"`
	Country     string      `json:"country"`
}

type coordinates struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type bank struct {
	CardExpire string `json:"cardExpire"`
	CardNumber string `json:"cardNumber"`
	CardType   string `json:"cardType"`
	Currency   string `json:"currency"`
	IBAN       string `json:"iban"`
}

type company struct {
	Department string  `json:"department"`
	Name       string  `json:"name"`
	Title      string  `json:"title"`
	Address    address `json:"address"`
}

type crypto struct {
	Coin    string `json:"coin"`
	Wallet  string `json:"wallet"`
	Network string `json:"network"`
}

// Interfaces Implementation
func (user *User) GetResourceID() int {
	return user.ID
}

func (user *User) AttachChildren(children any, relation string) any {

	switch relation {
	case "posts":
		var concretePosts []Post
		unpackedPosts, _ := children.([]MergeableChild)

		for idx := range unpackedPosts {
			concretePosts = append(concretePosts, *unpackedPosts[idx].(*Post))
		}

		user.Posts = concretePosts
	case "quotes":
		var concreteQuotes []Quote
		unpackedQuotes, _ := children.([]MergeableChild)

		for idx := range unpackedQuotes {
			concreteQuotes = append(concreteQuotes, *unpackedQuotes[idx].(*Quote))
		}

		user.Quotes = concreteQuotes
	case "recipes":
		var concreteRecipes []Recipe
		unpackedRecipes, _ := children.([]MergeableChild)

		for idx := range unpackedRecipes {
			concreteRecipes = append(concreteRecipes, *unpackedRecipes[idx].(*Recipe))
		}

		user.Recipes = concreteRecipes

	case "cart":
		var concreteCart Cart
		unpackedCart, _ := children.([]MergeableChild)

		for idx := range unpackedCart {
			concreteCart = *unpackedCart[idx].(*Cart)
		}

		user.Cart = concreteCart
	}

	return nil
}

func (user *User) GetChildren(childrenType string, childrenID int) any {
	return user.Posts[childrenID].GetResourceID()
}
