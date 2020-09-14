package models


import (
	"time"
)

type User struct {
	Id 					string 			`bson:"_id" 		json:"id"`
	NickName 			string 			`bson:"nick_name"	json:"nick_name"`
	Account 			string 			`bson:"account" 	json:"account"`
	Password 			string 			`bson:"password" 	json:"password"`
	BasicInfo			struct {
		Name 				string 			`bson:"name" 		json:"name"`
		Email 				string 			`bson:"email" 		json:"email"`
		Address				string  		`bson:"address" 	json:"address"`
		Phone  				string  		`bson:"phone" 		json:"phone"`
		App 				struct {
			AppName 			string 			`bson:"app_name" 	json:"app_name"`
			AppAccount 			string 			`bson:"app_account 	json:"app_account"`
		}									`bson:"app" 		json:"app"`
	}									`bson:"basic_info" 	json:"basic_info"`
	Photo 				struct{
		Path 				string 			`bson:"path" 		json:"path"`
		UpdateTime 			*time.Time 		`bson:"update_time" json:"update_time"`
	}									`bson:"photo" 		json:"photo"`
	Role 				map[string]interface{} 	`bson:"role" 	json:"role"`
	CreateTime 			*time.Time 		`bson:"create_time"	json:"create_time"`
}

type SearchUser struct {
	Payload 			struct{
		Account 			string 			`bson:"account" 	json:"account"`
		NickName 			string 			`bson:"nick_name"	json:"nick_name"`
	}									`bson:"payload" 	json:"payload"`
	Ticket 				string 			`bson:"ticket" 		json:"ticket"`
}

type AddUser struct {
	Payload 			struct{
		NewUser 			User 			`bson:"new_user" 	json:"new_user"`
 	} 									`bson:"payload" 	json:"payload"`
	Ticket 				string 			`bson:"ticket" 		json:"ticket"`
}

type ModifyUser struct {
	Payload 			struct{
		Account 			string 			`bson:"account" 	json:"account"`
		Password 			string 			`bson:"password" 	json:"password"`
		ChangeField 		map[string]interface{} `bson:"change_field" json:"change_field"`
	}									`bson:"payload" 	json:"payload"`
	Ticket 				string 			`bson:"ticket" 		json:"ticket"`
}

type DeleteUser struct {
	Payload 			struct {
		Account 			string 			`bson:"account" 	json:"account"`
		Password 			string 			`bson:"password" 	json:"password"`
	}									`bson:"payload" 	json:"payload"`
	Ticket 				string 			`bson:"ticket" 		json:"ticket"`
}














