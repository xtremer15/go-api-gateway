package types

type Cart struct {
	ID       int       `json:"id"`
	UserID   int       `json:"user_id"`
	Products []Product `json:"products"`
}

func (cart *Cart) GetResourceID() int {
	return cart.ID
}

func (cart *Cart) AttachChildren(children any, childType string) any {

	var concreteProducts []Product
	unpackedProducts, _ := children.([]MergeableChild)

	for idx := range unpackedProducts {
		concreteProducts = append(concreteProducts, *unpackedProducts[idx].(*Product))
	}

	cart.Products = concreteProducts

	return nil
}

func (cart *Cart) GetChildren(childrenType string, childrenID int) any {
	return cart.Products[childrenID]
}

func (cart *Cart) GetChildType() string {
	return "cart"
}

func (cart *Cart) GetParentID() int {
	return cart.UserID
}

func (cart *Cart) GetChildKey() int {
	return cart.ID
}
