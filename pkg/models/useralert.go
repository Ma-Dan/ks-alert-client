package models

import (
	"github.com/carmanzhang/ks-alert-client/pkg/utils/dbutil"
	"github.com/jinzhu/gorm"
)

type ResourceType string

const (
	Cluster     ResourceType = "cluster"
	Node        ResourceType = "node"
	Workspace   ResourceType = "workspace"
	Namespace   ResourceType = "namespace"
	Application ResourceType = "application"
	Workload    ResourceType = "workload"
	Pod         ResourceType = "pod"
	Container   ResourceType = "container"
)

// table `user_alert_bindings` record alerts created by user
// associated table of table `user` and table `alert`
type UserAlertBinding struct {
	gorm.Model
	UserID        string `gorm:"type:varchar(50);not null;"`
	AlertConfigID string `gorm:"type:varchar(50);not null;"`
	ResourceType  string `gorm:"type:varchar(50);not null;"`
	ResourceName  string `gorm:"type:varchar(50);not null;"`
	ProductID     string `gorm:"type:varchar(50);not null;"`
	// local cluster
	Cluster   string
	Node      string
	Workspace string
	Namespace string
	Workload  string
	Pod       string
}

func GetAlertByResourceName(userAlertBind *UserAlertBinding) (*[]UserAlertBinding, error) {
	db, err := dbutil.DBClient()
	if err != nil {
		panic(err)
	}

	var userAlerts []UserAlertBinding
	db.Model(&UserAlertBinding{}).Where(userAlertBind).Find(&userAlerts)
	return &userAlerts, db.Error
}

func GetAlertByUserID(userID string) (*[]UserAlertBinding, error) {
	db, err := dbutil.DBClient()
	if err != nil {
		panic(err)
	}

	var userAlerts []UserAlertBinding
	db.Model(&UserAlertBinding{}).Where(&UserAlertBinding{UserID: userID}).Find(&userAlerts)
	return &userAlerts, db.Error
}

func GetAlertConfigIDs(userAlerts *[]UserAlertBinding) []string {
	var alertConfigIDs []string
	for i := 0; i < len(*userAlerts); i++ {
		if (*userAlerts)[i].AlertConfigID != "" {
			alertConfigIDs = append(alertConfigIDs, (*userAlerts)[i].AlertConfigID)
		}
	}

	return alertConfigIDs
}
