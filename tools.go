package ascii

func IsEmpty(slice []string) bool {
	for _, v := range slice {
		if v != "" {
			return false
		}
	}
	return true
}
func IsPrintable(s string) bool {
	for _, char := range s {
		if( char < 32 || char > 126 ) && char != '\n' && char != '\r' && char != '\t' && char != ' ' && char != '\v' && char != '\f' {
			return false
		}
	}
	return true
}