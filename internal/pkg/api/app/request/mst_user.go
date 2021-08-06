package request

type UserRequest struct {
	ID       int    `json:"id" uri:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}
