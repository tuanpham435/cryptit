package encrypt

func Nimbus(str string) string {
	encryptedStr := ""
	for _, c := range str {
		asciiCode := byte(c)
		character := string(asciiCode + 3)
		encryptedStr += character
	}

	return encryptedStr
}
