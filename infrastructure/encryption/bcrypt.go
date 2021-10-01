package encryption

import "golang.org/x/crypto/bcrypt"

func Hash(plaintText string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plaintText), 10)
	if err != nil {
		panic(err)
	}

	return string(hashed)
}

func Compare(hashed string, plaintext string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plaintext)); err != nil {
		return false
	}
	return true
}
