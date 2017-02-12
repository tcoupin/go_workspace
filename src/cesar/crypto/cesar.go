package crypto

import "strings"
var alphabet = [26]string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}

func Cesar(text string, delta int) string {
	result := make([]string, len(text))
	for i:=0;i<len(text);i++ {
		ind := letter2indice(string(text[i])) + delta
		if ind >= 26 {
			ind -= 26
		}
		result[i] = alphabet[ind]
	}
	return strings.Join(result,"")
}

func letter2indice(letter string) int {
	for i:=0;i<len(alphabet);i++ {
		if letter == alphabet[i] {
			return i
		}
	}
	return -1
}