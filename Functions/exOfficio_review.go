package Functions

import "strings"

func ExOfficioReview(text string, char int) bool {
	var ret = false
	var exOfficioReview []string

	var less17 = []string{"reexame necessário", "xame necessário"}
	var more17 = []string{"reexame necessário"}

	if char < 17 {
		exOfficioReview = less17
	} else {
		exOfficioReview = more17
	}
	for i := 0; i < len(exOfficioReview); i++ {
		if strings.Contains(text, exOfficioReview[i]) {
			ret = true
		}
	}
	return ret
}
