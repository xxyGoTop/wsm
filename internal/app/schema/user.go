package schema

type ProfilePure struct {
	Id  string `json:"id"`
	Username string `json:"username"`
	Nickname *string `json:"nickname"`
	Gender int `json:"gender"`
	Avatar string `json:"avatar"`
}

type Profile struct {
	ProfilePure
	CreatedAt string `json:"created_at"`
	UpdateAt string `json:"update_at"`
}

type ProfileWithToken struct {
	Profile
	Token string `json:"token"`
}