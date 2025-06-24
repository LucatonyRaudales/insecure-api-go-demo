package user

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

var UnsecureUsers = []User{
	{ID: 1, Name: "Alice", Email: "alice@example.com", Password: "123456", Token: "tok_alice"},
	{ID: 2, Name: "Bob", Email: "bob@example.com", Password: "password", Token: "tok_bob"},
}

var Users = []User{}