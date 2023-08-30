package user

type User struct {
	Uid             uint64 `gorm:"primaryKey;autoIncrement;not null"`
	Mobile          string
	Password        string `gorm:"->:false"`
	salt            string `gorm:"->:false"`
	Nickname        string
	Udid            string
	UserToken       string
	FromChannel     uint8
	FromMarket      uint8
	Product         uint8
	Gender          uint8
	AutoSub         uint8
	IsVip           uint8
	Avatar          string
	RegType         uint8
	AgentAdminId    uint8
	RegIp           string
	AreaId          int
	RegTime         int
	LastLoginIp     int
	LastLoginTime   int
	Status          uint8
	SensitiveStatus uint8
	DeviceId        uint8
	UserType        uint8
	SafetyCode      string
	InviteCode      string
	MasterUid       int
	Idfa            string
	CreatedAt       int
	UpdatedAt       int
	DeletedAt       int
}
