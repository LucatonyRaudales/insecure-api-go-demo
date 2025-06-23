package user

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

var Users = []User{
	{1, "Alice", "alice@example.com", "123456", "tok_alice"},
	{2, "Bob", "bob@example.com", "password", "tok_bob"},
}