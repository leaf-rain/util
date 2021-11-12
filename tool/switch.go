package tool

type Switch int16

func (s *Switch) TurnOn(index int) {
	*s = 1<<index | *s
}

func (s *Switch) TurnOff(index int) {
	*s = (^(1 << index)) & *s
}

func (s Switch) CheckTurnOn(index int) (result bool) {
	return s == 1<<index|s
}
