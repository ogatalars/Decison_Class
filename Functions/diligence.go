package Functions

import "strings"

func Diligence(text string, char int) bool {
	var ret = false
	var diligence []string

	var less17 = []string{"conversão do julgamento em diligência", "convertido em diligência", "diligência"}
	var more17 = []string{"acórdão em diligência", "conversão do julgamento em diligência", "conversão do feito em diligência", "convertido em diligência", "decisão em diligência", "julgalmento em diligência", "sentença em diligência"}

	if char < 17 {
		diligence = less17
	} else {
		diligence = more17
	}
	for i := 0; i < len(diligence); i++ {
		if strings.Contains(text, diligence[i]) {
			ret = true
		}

	}
	return ret
}
