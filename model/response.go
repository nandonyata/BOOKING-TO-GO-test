package model

type ResponseFetchCustomer struct {
	Nationality string       `json:"nationality"`
	Name        string       `json:"name"`
	Dob         string       `json:"dob"`
	Phone       string       `json:"phone"`
	Email       string       `json:"email"`
	Family_list []FamilyList `json:"family_list"`
}
