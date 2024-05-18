package ascii

func IsEmpty(slice []string) bool {
	for _, v := range slice {
		if v != "" {
			return false
		}
	}
	return true
}
func IsPrintable(char rune) bool {
	return (char < 127 && char > 31) 
}
func DefaultFont(font, defaultValue string) string {
	if font == "" {
		return defaultValue
	}
	return font
}