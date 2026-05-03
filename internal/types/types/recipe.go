package types

type Recipe struct {
	ID                 int      `json:"id"`
	Name               string   `json:"name"`
	Ingredients        []string `json:"ingredients"`
	Instructions       []string `json:"instructions"`
	PrepTimeMinutes    int      `json:"prepTimeMinutes"`
	CookTimeMinutes    int      `json:"cookTimeMinutes"`
	Servings           int      `json:"servings"`
	Difficulty         string   `json:"difficulty"`
	Cuisine            string   `json:"cuisine"`
	CaloriesPerServing int      `json:"caloriesPerServing"`
	Tags               []string `json:"tags"`
	UserID             int      `json:"userId"`
	Image              string   `json:"image"`
	Rating             float64  `json:"rating"`
	ReviewCount        int      `json:"reviewCount"`
	MealType           []string `json:"mealType"`
}

func (r *Recipe) GetParentID() int {
	return r.UserID
}
func (r *Recipe) GetChildKey() int {
	return r.ID
}
func (r *Recipe) GetChildType() string {
	return "recipes"
}
