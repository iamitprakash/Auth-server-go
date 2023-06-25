package users

type user struct {
	email        string
	userName     string
	passwordHash string
	fullName     string
	role         int
	createdDate  string
}

// You can get this from Database
var userList = []user{
	{
		email:        "abc@gmail.com",
		userName:     "abc12",
		passwordHash: "hashedme1",
		fullName:     "abc def",
		createdDate:  "1631600786",
		role:         1,
	},
	{
		email:        "chekme@example.com",
		userName:     "checkme34",
		passwordHash: "hashedme2",
		fullName:     "check me",
		createdDate:  "1631600837",
		role:         0,
	},
}

// based on the email id provided, finds the user object
// can be seen as the main constructor to start validation
func GetUserObject(email string) (user, bool) {
	//needs to be replaces using Database
	for _, user := range userList {
		if user.email == email {
			return user, true
		}
	}
	return user{}, false
}

// checks if the password hash is valid
func (u *user) ValidatePasswordHash(pswdhash string) bool {
	return u.passwordHash == pswdhash
}

// this simply adds the user to the list
func AddUserObject(email string, username string, passwordhash string, fullname string, role int) bool {
	// declare the new user object
	newUser := user{
		email:        email,
		passwordHash: passwordhash,
		userName:     username,
		fullName:     fullname,
		role:         role,
	}
	// check if a user already exists
	for _, ele := range userList {
		if ele.email == email || ele.userName == username {
			return false
		}
	}
	userList = append(userList, newUser)
	return true
}
