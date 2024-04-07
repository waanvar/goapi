package models

import "gorm.io/gorm"

type Useracl struct {
	gorm.Model
	UserID       string
	UserEmail    string
	UserWho      string
	UserAcl      string
	UserName     string
	UserPassword string
	IPerson      string
	UserCom      string
	IEmploy      string
	UserGroup    string
	UserPo       string
	UserLoad     string
}
