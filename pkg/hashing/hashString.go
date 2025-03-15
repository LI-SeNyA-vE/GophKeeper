package hashing

import "golang.org/x/crypto/bcrypt"

// HashString хеширует пароль с bcrypt
func HashString(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CheckString сравнивает пароль с хешем
func CheckString(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
