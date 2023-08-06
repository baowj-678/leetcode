package main

import "math"

func removeComments(source []string) []string {
	res := []string{}

	blockComments := false

	checkBlockCommentLeft := func(s string) int {
		for i := 0; i < len(s)-1; i++ {
			if s[i] == '/' && s[i+1] == '*' {
				return i
			}
		}
		return math.MaxInt
	}

	checkBlockCommentRight := func(s string) (bool, string) {
		for i := 0; i < len(s)-1; i++ {
			if s[i] == '*' && s[i+1] == '/' {
				if i+2 < len(s) {
					return true, s[i+2:]
				} else {
					return true, ""
				}
			}
		}
		return false, ""
	}

	checkLineComment := func(s string) int {
		for i := 0; i < len(s)-1; i++ {
			if s[i] == '/' && s[i+1] == '/' {
				return i
			}
		}
		return math.MaxInt
	}

	for i := 0; i < len(source); i++ {
		line := source[i]
		if blockComments {
			if ok, s := checkBlockCommentRight(line); ok {
				source[i] = s
				i--
				blockComments = false
			}
		} else {
			blockIdx := checkBlockCommentLeft(line)
			lineIdx := checkLineComment(line)

			if blockIdx != math.MaxInt && blockIdx < lineIdx {
				s := line[:blockIdx]
				res = append(res, s)
				blockComments = true
				if blockIdx+2 < len(line) {
					if ok, s := checkBlockCommentRight(line[blockIdx+2:]); ok {
						res[len(res)-1] += s
						blockComments = false
					}
				}
				continue
			}
			if lineIdx != math.MaxInt && lineIdx < blockIdx {
				res = append(res, line[:lineIdx])
			} else {
				res = append(res, line)
			}
		}
	}

	finalRes := []string{}
	for _, s := range res {
		if len(s) > 0 {
			finalRes = append(finalRes, s)
		}
	}
	return finalRes
}

func main() {
	source := []string{"a/*/b//*c", "blank", "d/*/e*//f"}
	removeComments(source)
}
