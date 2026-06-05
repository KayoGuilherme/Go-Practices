package validators

import "unicode"

func isValidCPF(cpf string) bool {
	if len(cpf) != 11 {
		return false
	}
	for _, r := range cpf {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	allSame := true
	for i := 1; i < len(cpf); i++ {
		if cpf[i] != cpf[0] {
			allSame = false
			break
		}
	}
	if allSame {
		return false
	}

	d1 := cpfCheckDigit(cpf[:9], 10)
	d2 := cpfCheckDigit(cpf[:10], 11)
	return cpf[9] == d1 && cpf[10] == d2
}

func cpfCheckDigit(digits string, weight int) byte {
	sum := 0
	for i := 0; i < len(digits); i++ {
		sum += int(digits[i]-'0') * weight
		weight--
	}
	remainder := sum % 11
	if remainder < 2 {
		return '0'
	}
	return byte('0' + (11 - remainder))
}
