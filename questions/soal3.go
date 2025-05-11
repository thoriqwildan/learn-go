package questions

type User struct {
	Name string
	Age int
}

func UpdateProfile(user *User, name string, age int) {
	user.Name = name
	user.Age = age
}