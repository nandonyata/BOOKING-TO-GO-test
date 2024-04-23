package model

type FamilyList struct {
	Name     string `json:"name"`
	Dob      string `json:"dob"`
	Relation string `json:"relation"`
}

type ReqBodyCreateCustomer struct {
	Nationality_id int          `json:"nationality_id"`
	Name           string       `json:"name"`
	Dob            string       `json:"dob"`
	Phone          string       `json:"phone"`
	Email          string       `json:"email"`
	Family_list    []FamilyList `json:"family_list"`
}
