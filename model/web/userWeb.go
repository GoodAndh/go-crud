package web

type UserRegisterPayload struct {
	Username string `json:"username" validate:"required,min=8,max=30"`
	Password string `json:"password" validate:"required,min=8"`
	CPassword string `json:"cpassword" validate:"required,min=8,eqfield=Password"`
	Email    string `json:"email" validate:"required"`
	Name     string `json:"name" validate:"required"`
}
type UserLoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}


type UserWeb struct {
	Username string `json:"username"`
	Password string `json:"-"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}