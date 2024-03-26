package cryptoes

import "unicode/utf8"

func DesensitizeName(name string) string {
	runeCount := utf8.RuneCountInString(name)
	if runeCount == 0 {
		return ""
	}
	first, _ := utf8.DecodeRuneInString(name)
	newName := string(first)
	for i := 1; i < runeCount-1; i++ {
		newName += "*"
	}
	if runeCount > 2 {
		last, _ := utf8.DecodeLastRuneInString(name)
		newName += string(last)
	} else {
		newName += "*"
	}
	return newName
}

func desensitizeNumber(number string, firstCount, lastCount int) string {
	if firstCount+lastCount >= len(number) {
		return number
	}
	newNumber := number[:firstCount]
	end := len(number) - lastCount
	for i := firstCount; i < end; i++ {
		newNumber += "*"
	}
	if end < len(number) {
		newNumber += number[end:]
	}

	return newNumber
}

func DesensitizePhone(phone string) string {
	return desensitizeNumber(phone, 3, 4)
}

func DesensitizeIdentityCard(identityCard string) string {
	return desensitizeNumber(identityCard, 6, 4)
}
