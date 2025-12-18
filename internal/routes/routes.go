package routes

type appRoutes struct {
	Users    string
	Products string
}

var PathRoutes = appRoutes{
	Users:    "/users",
	Products: "/products",
}
