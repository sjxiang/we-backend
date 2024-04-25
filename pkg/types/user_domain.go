package types

 import (
	"time"
)


type User struct {
	ID        int64            

	Email     string           
	Password  string           

	Nickname  string           
	Mobile    string           
	Intro     string           
	Avatar    string
	Birthday  int64   
	Gender    string  
	Role      int8     

	CreatedAt time.Time
	UpdatedAt time.Time
}


type CreateUserParams struct {
	Password string   
	Email    string   
}