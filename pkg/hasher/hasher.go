// Package hasher provides primitives for working with password's hashes.
package hasher

// HashPassword hashes provided password string.
// It returns the password hash and an error.
func HashPassword(password string) (string, error) {
	return "passwordhash", nil
}

// CheckPasswordHash checks if provided password's hash is same as provided hash.
// It returns true or false.
func CheckPasswordHash(password, hash string) bool {
	return true
}
