// decrypt package consist of all the decryption algorithms
package decrypt

// decrypts by reducing the ascii code by 3 for each character
func Nimbus(str string) string {
	decryptedStr := ""

	for _, c := range str {
		asciiCode := byte(c)
		character := string(asciiCode - 3)
		decryptedStr += character
	}

	return decryptedStr
}
