package convert

func CToF(c float64) float64 { return float64(c*9/5 + 32) }

func FToC(f float64) float64 { return float64((f - 32) * 5 / 9) }
