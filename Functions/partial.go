package Functions

import "strings"

func Partial(text string, char int) bool {
	var ret = false
	var partial []string

	var less17 = []string{"adequação parcial", "almente mantid", "arte conhecid", "cial provimento", "dência parcial", "em parte", "ialmente", "lmente", "mente", "nessa parte", "nesta parte", "neste tópico", "parcial", "parcial provimento", "parcialmente modificada", "parcialmente modificado", "parcialmente provid", "parcialmente", "parte acolhid", "parte apreciad", "parte chonhecid", "parte e provid", "parte provid", "parte, portanto", "provido, em parte", "provido em parte", "rcial provimento", "reformada parcialmente", "te provid"}
	var more17 = []string{"adequação parcial", "em parte", "nessa parte", "nesta parte", "parcialmente", "conhecido neste tópico", "em parte acolhid", "em parte apreciad", "em parte conhecid", "em parte e provid", "em parte provid", "em parte, portanto", "parcial provimento", "parcial provimento", "parcialmente mantid", "parcialmente modificada", "parcialmente modificado", "paricalmente provid", "parcialmente provid", "parte conhecid", "procedência parcial", "provido, em parte", "provido em parte", "reformada parcialmente"}

	if char < 17 {
		partial = less17
	} else {
		partial = more17
	}
	for i := 0; i < len(partial); i++ {
		if strings.Contains(text, partial[i]) {
			ret = true
		}
	}
	return ret
}
