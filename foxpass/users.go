package users

import "time"

type Users struct {
	Status string `json:"status"`
	Data   []struct {
		Username       string      `json:"username"`
		Email          string      `json:"email"`
		UID            int         `json:"uid"`
		Created        time.Time   `json:"created"`
		Gid            int         `json:"gid"`
		IsEngUser      bool        `json:"is_eng_user"`
		IsPosixUser    bool        `json:"is_posix_user"`
		IsActive       bool        `json:"is_active"`
		Active         bool        `json:"active"`
		Shell          interface{} `json:"shell"`
		FirstName      string      `json:"first_name"`
		LastName       string      `json:"last_name"`
		GithubUsername interface{} `json:"github_username"`
		CustomFields   struct {
			DevUserName interface{} `json:"devUserName"`
		} `json:"custom_fields"`
	} `json:"data"`
	Page interface{} `json:"page"`
}
