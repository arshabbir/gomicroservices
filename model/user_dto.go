package model

var users map[int]User = map[int]User{121: User{Id: 7,
	Name: "arshabbir",
	Age:  39}}

func GetUser(id int) User {

	users[122] = User{Id: 8, Name: "Lover", Age: 30}
	users[123] = User{Id: 87, Name: "Lover1", Age: 20}

	var user User

	user = users[id]
	return user

}
