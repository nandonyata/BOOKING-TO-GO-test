package entity

type Customer struct {
	Id             int    `json:"id"`
	Nationality_id int    `json:"nationality_id"`
	Cst_name       string `json:"cst_name"`
	Cst_dob        string `json:"cst_cst_dob"` // DD-MM-YY
	Cst_phoneNum   string `json:"cst_phoneNum"`
	Cst_email      string `json:"cst_email"`
}
