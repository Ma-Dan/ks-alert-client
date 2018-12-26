package models

import "github.com/carmanzhang/ks-alert-client/pkg/utils/dbutil"

func init() {

	db, err := dbutil.DBClient()

	if err != nil {
		panic(err)
	}

	if !db.HasTable(&UserAlertBinding{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&UserAlertBinding{}).Error; err != nil {
			panic(err)
		}
	}
}
