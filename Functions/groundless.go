package Functions

import "strings"

func Groundless(text string, char int) bool {
	var ret = false
	var groundless []string

	var less17 = []string{"a-se provimento", "acórdão confirmado", "afastad", "ão conhecid", "ção do decidido", "confirmad", "de nenhum deles se conhece", "decisão confirmada", "degenada", "denega", "denegada", "denegado", "deprovid", "desacolhid", "desacolhida", "desacolhido", "descabid", "descabida", "descabido", "descabidos", "descabimento", "desconheci", "desprovi", "desprovid", "desprovida", "desprovido ", "desprovido", "desprovido", "desprovimento do recurso", "desprovimento", "ega conhecimento", "gado provimento", "impossibilidade", "improcedência mantida", "improcedência", "improcedente", "improvid", "impróvid", "improvida", "improvido", "improvimen", "improvimento", "inaplicabilidade do princípio da fungibilidade recursal", "inadimissível", "inadimitida", "inadimitido", "inadmitido", "indefere-se", "indefere", "indeferi", "indeferid", "indeferimento", "indevid", "julgamento confirmado", "mantém-se", "mantém", "mantid", "mantida", "mantido", "manutenção do decidido", "manutenção", "não autorizado", "não cabe conhecer", "não cabe reconhecer", "não conheceram", "não conhecid", "não conhecida", "não conhecido", "não conhecimento", "não provid", "não provimento", "não se conhece", "não se há de", "não há fundamento para acolhimento", "não", "nega conhecimento", "nega provimento", "nega-se provimento", "nega", "negado provimento", "negam-se", "negaram provimento", "nego provimento", "nego", "rejei", "rejeição ", "rejeição", "rejeitad", "rejeitada", "rejeitam-se", "rejeitado", "rejeito", "repelid", "repelido", "repelidos", "retratação não exercid", "se os embargos", "sem alteração do resultado", "sem alteração dos resultados", "sem efeitos modificativos", "sem efeito modificativo", "sentença confirmada"}
	var more17 = []string{"acórdão confirmado", "afastad", "de nenhum deles se conhece", "decisão confirmada", "degenada", "denega", "denegada", "denegado", "deprovid", "desacolhid", "desacolhida", "desacolhido", "descabid", "descabida", "descabido", "descabimento", "desconheci", "desprovid", "desprovida", "desprovido ", "desprovido", "desprovimento", "impossibilidade", "improcedência mantida", "improcedência", "improcedente", "impróvid", "improvida", "improvido", "improvimento", "inaplicabilidade do princípio da fungibilidade recursal", "inadimissível", "inadimitida", "inadimitido", "indefere-se", "indefere", "indeferid", "indeferimento", "indevid", "julgamento confirmado", "mantém-se", "mantém", "mantida", "mantido", "manutenção do decidido", "manutenção", "não autorizado", "não cabe conhecer", "não conheceram", "não conhecid", "não conhecida", "não conhecido", "não conhecimento", "não provid", "não provimento", "não se conhece", "não se há de", "não há fundamento para acolhimento", "nega conhecimento", "nega provimento", "nega-se provimento", "negado provimento", "negado conhecimento", "negam-se", "negaram provimento", "nego provimento", "rejeição ", "rejeição", "rejeitad", "repelido", "rejeitam-se", "repelidos", "retratação não exercid", "sem alteração do resultado", "sem alteração dos resultados", "sem efeito modificativo", "sem efeitos modificativos", "sentença confirmada"}

	if char < 17 {
		groundless = less17
	} else {
		groundless = more17
	}
	for i := 0; i < len(groundless); i++ {
		if strings.Contains(text, groundless[i]) {
			ret = true
		}
	}
	return ret
}
