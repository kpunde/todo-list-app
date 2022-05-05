package utility

import "golang.org/x/crypto/bcrypt"

func HashAndSalt(pass string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		panic(err.Error())
	}

	return string(hashed)
}

func IsPasswordMatch(hash string, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err != nil {
		return false
	}

	return true
}
