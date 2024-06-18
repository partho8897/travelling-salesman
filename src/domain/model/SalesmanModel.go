package model

type SalesPerson struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Area struct {
	XCord int `json:"x_cord"`
	YCord int `json:"y_cord"`
}

type GetUniqueDestinationResponse struct {
	UniqueDestinations int `json:"unique_destinations"`
}
