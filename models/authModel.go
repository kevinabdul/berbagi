package models

type Resource struct {
	ID 		uint  	`gorm:"primaryKey" json:"id"`
	Path    string  `json:"path"`
}

type Action struct {
	ID 		uint  	`gorm:"primaryKey" json:"id"`
	Name    string  `json:"name"`
}

type Permission struct {
	ID  		uint   		`gorm:"primaryKey"`
	ActionID 	uint 		`gorm:"primaryKey" json:"action_id"`
	ResourceID	uint 		`gorm:"primaryKey" json:"resource_id"`
	Action 		Action 		`gorm:"foreignKey:ActionID" json:"action_id"`
	Resource  	Resource 	`gorm:"foreignKey:ResourceID" json:"resource_id"`
}

type Role struct {
	ID  		uint  	`gorm:"primaryKey" json:"id"`
	Name 		string	`gorm:"unique" json:"role_name"`
}

type RolePermission struct {
	RoleID  		uint      	`gorm:"primaryKey"`
	PermissionID	uint  		`gorm:"primaryKey"`
	Role 			Role
	Permission  	Permission
}

type PermissionAPI struct {
	RoleName  	string `json:"role_name"` 
	ActionName  string `json:"action_name"`
	Path     	string `json:"path"`
}