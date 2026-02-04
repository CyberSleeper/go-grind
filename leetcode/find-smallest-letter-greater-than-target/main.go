package main

func nextGreatestLetter(letters []byte, target byte) byte {
	st := byte('z')
	st++
	ans := st
	for _, v := range letters {
		if v < ans && v > target {
			ans = v
		}
	}
	if ans == st {
		ans = letters[0]
	}

	return ans
}
