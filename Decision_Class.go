package Decison_Class

import (
	"github.com/Darklabel91/Decison_Class/Classifier"
	"github.com/Darklabel91/Decison_Class/Struct"
)

func Decision_Classifier(summary string, identifier string, court string) Struct.Infered_decision {

	ret := Classifier.ClassDecision(summary, identifier, court, 16)

	if ret.Class == "Não Mapeado" {
		ret = Classifier.ClassDecision(summary, identifier, court, 70)
		if ret.Class == "Não Mapeado" || ret.Class == "Sem informaçao" {
			ret = Classifier.ClassDecision(summary, identifier, court, len(summary)/2)
			if ret.Class == "Não Mapeado" {
				ret = Classifier.ClassDecision(summary, identifier, court, len(summary))
			}
		}
	}

	return ret

}
