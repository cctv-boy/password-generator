package password_generator

import "math/rand"

func GeneratePassword(length int, charset string) string {
	password := make([]byte, length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}
	return string(password)
}
