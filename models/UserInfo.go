package models

type Userinfo struct {
	ID         uint64 `json:"id" gorm:"primary_key"`
	First_name string `json:"first_name"`
	Last_name  string `json:"Last_name"`
	Phone      string `json:"phone"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

// type API_SECRET struct {
// 	Book  	 Book
// 	Username string `json:"username"`
//  	Password string `json:"password"`
// }
