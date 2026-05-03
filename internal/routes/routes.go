package routes

type appRoutes struct {
	Users     string
	Products  string
	Todos     string
	Comments  string
	Posts     string
	Aggregate string
	Carts     string
	Quotes    string
	Recipes   string
}

var PathRoutes = appRoutes{
	Users:     "/users",
	Products:  "/products",
	Todos:     "/todos",
	Comments:  "/comments",
	Posts:     "/posts",
	Aggregate: "/aggregate",
	Carts:     "/carts",
	Quotes:    "/quotes",
	Recipes:   "/recipes",
}
