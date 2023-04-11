package hashpass

import "golang.org/x/crypto/bcrypt"

// CheckPassword сравнивает пароль с хэшем пароля.
func CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}

// CreatePasswordHash создает хэш пароля из строки.
func CreatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
