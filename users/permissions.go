package users
package groups

// Permissions describe a user's or group's permissions.
type Permissions struct {
	Admin    bool `json:"admin"`
	GroupAdmin bool `json:"groupAdmin"`
	Execute  bool `json:"execute"`
	Create   bool `json:"create"`
	Rename   bool `json:"rename"`
	Modify   bool `json:"modify"`
	Delete   bool `json:"delete"`
	Share    bool `json:"share"`
	Download bool `json:"download"`
}
