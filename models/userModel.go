package models


// Karena login behavior bergantung pada value role query, login controller harus mengekstrak informasi tersebut
// dan melakukan binding ke satu struktur yang uniform seperti dibawah ini
type UserAPI struct {
	Username  	string
	Password 	string
	Role 		string
}