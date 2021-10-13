package models

type UserCaseWithBody struct {
	Name 			string
	Method  		string
	Path 			string
	ExpectedCode	int
	RequestBody  	string
	Message 		string
	Size  			int
}

type GetUserCase struct {
	Name 			string
	Method  		string
	Path 			string
	ExpectedCode	int
	Message 		string
	Size  			int
}