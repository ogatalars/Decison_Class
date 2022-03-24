package Functions

import "strings"

func Affected(text string, char int) bool {
	var ret = false
	var affected []string

	var less17 = []string{"dicado", "extinção da punibilidade", "icado", "perda do objeto", "prejudicad", "prejudicado", "prejudicada", "perda superveniente do objeto", "prejudicando o exame"}
	var more17 = []string{"extinção da punibilidade", "perda do objeto", "prejudicad", "prejudicado", "prejudicada", "perda superveniente do objeto", "prejudicando o exame"}

	if char < 17 {
		affected = less17
	} else {
		affected = more17
	}
	for i := 0; i < len(affected); i++ {
		if strings.Contains(text, affected[i]) {
			ret = true
		}
	}
	return ret
}
