package authdata

type AuthData struct {
	Guid       int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	IPadd      string `json:"ip"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	Session_id int    `json:"session_id"`
}
