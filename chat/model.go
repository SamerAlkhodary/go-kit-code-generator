package chatService

type User struct {
	firstName string `json:"firstName"`

	lastName string `json:"lastName"`

	profilePic string `json:"profilePic"`

	token string `json:"token"`
}

func MakeNewUser(firstName string, lastName string, profilePic string, token string) *User {
	return &User{

		firstName: firstName,

		lastName: lastName,

		profilePic: profilePic,

		token: token,
	}
}
func (user *User) GetFirstName() string {
	return user.firstName
}

func (user *User) SetFirstName(newFirstName string) {
	user.firstName = newFirstName
}

func (user *User) GetLastName() string {
	return user.lastName
}

func (user *User) SetLastName(newLastName string) {
	user.lastName = newLastName
}

func (user *User) GetProfilePic() string {
	return user.profilePic
}

func (user *User) SetProfilePic(newProfilePic string) {
	user.profilePic = newProfilePic
}

func (user *User) GetToken() string {
	return user.token
}

func (user *User) SetToken(newToken string) {
	user.token = newToken
}
