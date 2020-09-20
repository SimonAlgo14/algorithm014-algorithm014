package main

func numDecodings(s string) int {
	return decodings(s, 0)
}

func decodings(s string, index int) int {
	if len(s) == index {
		return 1
	}

	if s[index] == '0' {
		return 0
	}

	num := decodings(s, index+1)
	if index+1 < len(s) && (s[index] == '1' || (s[index] == '2' && s[index+1] < '7')) {
		num += decodings(s, index+2)
	}

	return num
}


func numDecodings_dp(s string) int {
	if len(s) <= 0 || s[0] == '0' {
		return 0
	}

	prev, curr := 1, 1

	for i := 1; i < len(s); i++ {
		tmp := curr
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				curr = prev  //10 or 20
			} else {
				return 0
			}
		}else if s[i-1] == '1' || ( s[i-1] == '2' && s[i] < '7' ) {
			curr += prev
		}

		prev = tmp
	}

	return curr
}