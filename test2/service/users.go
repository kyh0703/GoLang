package service

// defined collection name
const UserCollection = "users"

// instance
var UserCache = NewCacheMap[User](UserCollection)

// Register Hook API
func init() {
	Hook[UserCollection] = UserCache
}

type User struct {
	ID            string        `json:"_id"`
	Tenant        string        `json:"tntId"`
	Extension     string        `json:"email"`
	DidNum        string        `json:"didNum"`
	BillNum       string        `json:"billNum"`
	PhoneID       string        `json:"phoneId"`
	ServiceOption ServiceOption `json:"serviceOption"`
}

type ServiceOption struct {
	OutgoingBlock   bool    `json:"outgoingBlock,omitempty"`
	IncomingBlock   bool    `json:"incomingBlock,omitempty"`
	ForwardUse      bool    `json:"forwardUse,omitempty"`
	Forward         Forward `json:"forward,omitempty"`
	ReleaseToneUse  bool    `json:"releaseToneUse,omitempty"`
	TransferToneUse bool    `json:"transferedToneUse,omitempty"`
	EnableCallWait  bool    `json:"enableCallWait,omitempty"`
	Auth            Auth    `json:"auth,omitempty"`
}

type Forward struct {
	Type        string `json:"type,omitempty"`
	NoAnswerSec int    `json:"noAnswerSec,omitempty"`
	TransferNum string `json:"transferNum,omitempty"`
}

type Auth struct {
	Monitor      bool `json:"monitor,omitempty"`
	Coaching     bool `json:"coaching,omitempty"`
	AvoidMonitor bool `json:"avoidMonitor,omitempty"`
}
