/**
@author: Jason Pang
@desc:
@date: 2021/12/30
**/
package algorithm

var digToLetter map[string][]string
var digStr []string
var digStrArray []string

func letterCombinations(digits string) []string {
	digStrArray = make([]string, 0)
	if digits == "" {
		return digStrArray
	}

	digToLetter = make(map[string][]string)
	digToLetter["2"] = []string{"a", "b", "c"}
	digToLetter["3"] = []string{"d", "e", "f"}
	digToLetter["4"] = []string{"g", "h", "i"}
	digToLetter["5"] = []string{"j", "k", "l"}
	digToLetter["6"] = []string{"m", "n", "o"}
	digToLetter["7"] = []string{"p", "q", "r", "s"}
	digToLetter["8"] = []string{"t", "u", "v"}
	digToLetter["9"] = []string{"w", "x", "y", "z"}
	dig := make([]string, 0)
	for _, item := range digits {
		dig = append(dig, string(item))
	}
	digStr = make([]string, len(dig))
	letterComb(dig, 0, digStr)
	return digStrArray
}
func joinLetter(str []string) string {
	s := ""
	for _, item := range str {
		s += item
	}
	return s
}

func letterComb(dig []string, index int, str []string) {
	if index >= len(dig) {
		digStrArray = append(digStrArray, joinLetter(str))
		return
	}
	for _, letter := range digToLetter[dig[index]] {
		str[index] = letter
		letterComb(dig, index+1, str)
	}
}
