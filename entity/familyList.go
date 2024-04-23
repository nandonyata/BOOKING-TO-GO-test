package entity

type FamilyList struct {
	Id          int    `json:"id"`
	Cst_id      int    `json:"cst_id"`
	Fl_relation string `json:"fl_relation"`
	Fl_name     string `json:"fl_name"`
	Fl_dob      string `json:"fl_dob"`
}
