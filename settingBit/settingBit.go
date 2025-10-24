package settingBit

// SetBit устанавливает бит с индексом pos в значение value (0 или 1).
// Возвращает новое число и старое значение этого бита.
func SetBit(n int64, pos int, value int) (int64, int) {
	mask := int64(1) << pos
	oldBit := int((n >> pos) & 1) // сохраним старое значение

	if value == 1 {
		n |= mask // установка бита в 1
	} else {
		n &^= mask // установка бита в 0 (AND NOT)
	}
	return n, oldBit
}
