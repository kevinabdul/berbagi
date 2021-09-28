package models


// Karena login behavior bergantung pada value role query, login controller harus mengekstrak informasi tersebut
// dan melakukan binding ke satu struktur yang uniform seperti dibawah ini
type LoginUserAPI struct {
	Username  	string
	Email  		string
	Password 	string
	Role 		string
}

type LoginSearchAPI struct {
	ID     		uint
	Name 		string
	Email  		string
	Password  	string
}

type LoginResponseAPI struct {
	Message 	string
	Token 		string
}

