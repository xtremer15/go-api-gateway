package routes

type appRoutes struct {
	Users     string
	Products  string
	Todos     string
	Comments  string
	Aggregate string
}

var PathRoutes = appRoutes{
	Users:     "/users",
	Products:  "/products",
	Todos:     "/todos",
	Comments:  "/comments",
	Aggregate: "/aggregate",
}
