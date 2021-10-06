package iteration

func Repeat(char string, repeatTimes int) string {
	result := ""
	for i := 0; i < repeatTimes; i++ {	
		result += char
	}
	return result;
}

