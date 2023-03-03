package cryptorand

func GetAlphaNumericRandomString(size int) string {
	return randomString(size, true, true, true)
}

func GetRandomString(size int) string {
	return randomString(size, true, true, false)
}

func GetUpperCaseRandomString(size int) string {
	return randomString(size, false, true, false)
}

func GetLowerCaseRandomString(size int) string {
	return randomString(size, true, false, false)
}

func GetNumericRandomString(size int) string {
	return randomString(size, false, false, true)
}

func randomString(size int, includeLowerCase bool, includeUpperCase bool, includeNumeric bool) string {
	source := make([]string, 0)
	result := ""

	if includeLowerCase {
		source = append(source, LowerCaseAlphaCharacters()...)
	}

	if includeUpperCase {
		source = append(source, UpperCaseAlphaCharacters()...)
	}

	if includeNumeric {
		source = append(source, NumericCharacters()...)
	}
	if len(source) > 0 {
		random := Rand()
		if size > 0 {
			for i := 0; i < size; i++ {
				idx := random.Intn(len(source))
				result += source[idx]
			}
		}
	}

	return result
}
