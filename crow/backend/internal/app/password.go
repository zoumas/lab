package app

import "golang.org/x/crypto/bcrypt"

func hashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}

func passwordsMatch(hashedPassword, givenPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(givenPassword))
}
