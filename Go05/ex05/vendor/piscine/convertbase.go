package piscine

func ConvertBase(nbr, baseForm, baseTo string) string {
	deceimalValue := baseToDecimal(nbr, baseForm)
	result := decimalToBase(decimalValue, baseTo)
	return result
}

func baseToDecimal(nbr, baseFrom string) int {
	baseLen := len(baseFrom)
	baseMap := make(map[rune]int)
	for i, r := range baseFrom {
		baseMap[r] = i
	}
	decimalValue := 0
	for _, r := range nbr {
		decimalValue = decimalValue*baseLen + baseMap[r]
	}
	return decimalValue
}

func decimaltoBase(decimalValue int, baseTo string) string {
	if decimalValue == 0 {
		return string(baseTo[0])
	}
	baseLen := len(baseTo)
	result := ""
	for decimalValue > 0 {
		result = string(baseTo[decimalValue%baseLen]) + result
		decimalValue /= baseLen
	}
	return result
}
