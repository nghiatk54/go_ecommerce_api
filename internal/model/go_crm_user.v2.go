package model

const TableNameGoCrmUserV2 = "go_crm_user_v2"

// GoCrmUser Account
type GoCrmUserV2 struct {
	UsrID            int64  `gorm:"column:usr_id;primaryKey;autoIncrement:true;comment:Account Id" json:"usr_id"`                         // Account Id
	UsrEmail         string `gorm:"column:usr_email;type:varchar(255);not null;comment:Email" json:"usr_email"`                           // Email
	UsrPhone         string `gorm:"column:usr_phone;not null;comment:Phone number" json:"usr_phone"`                                      // Phone number
	UsrUsername      string `gorm:"column:usr_username;not null;comment:User name" json:"usr_username"`                                   // User name
	UsrPassword      string `gorm:"column:usr_password;not null;comment:Password" json:"usr_password"`                                    // Password
	UsrCreateAt      int64  `gorm:"column:usr_create_at;not null;comment:Creation time" json:"usr_create_at"`                             // Creation time
	UsrUpdateAt      int64  `gorm:"column:usr_update_at;not null;comment:Update time" json:"usr_update_at"`                               // Update time
	UsrCreateIPAt    string `gorm:"column:usr_create_ip_at;not null;comment:Creation IP" json:"usr_create_ip_at"`                         // Creation IP
	UsrLastLoginAt   int64  `gorm:"column:usr_last_login_at;not null;comment:Last login time" json:"usr_last_login_at"`                   // Last login time
	UsrLastLoginIPAt string `gorm:"column:usr_last_login_ip_at;not null;comment:Last login IP" json:"usr_last_login_ip_at"`               // Last login IP
	UsrLoginTimes    int64  `gorm:"column:usr_login_times;not null;comment:Login times" json:"usr_login_times"`                           // Login times
	UsrStatus        bool   `gorm:"column:usr_status;not null;comment:Status: 1 - enabled, 0 - disabled, -1 - deleted" json:"usr_status"` // Status: 1 - enabled, 0 - disabled, -1 - deleted
}

// TableName GoCrmUser's table name
func (*GoCrmUserV2) TableName() string {
	return TableNameGoCrmUserV2
}
