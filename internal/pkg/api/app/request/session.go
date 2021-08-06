package request

type SessionRequest struct {
	ID          int    `json:"id" uri:"id"`
	UserID      int    `json:"user_id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Start       string `json:"start"`
	Duration    int    `json:"duration"`
	Created     string `json:"created"`
	Updated     string `json:"updated"`
	Page        int    `form:"page"`
	Size        int    `form:"size"`
	Order       string `form:"order"`
}
