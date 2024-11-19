package calculate

import "math"

// 정규화된 가격으로 Parabolic SAR 계산
func CalculateParabolicSAR(highs, lows []float64) float64 {
	normalizedHighs := make([]float64, len(highs))
	normalizedLows := make([]float64, len(lows))

	for i := range highs {
		normalizedHighs[i] = normalizePrice(highs[i])
		normalizedLows[i] = normalizePrice(lows[i])
	}

	af := 0.02
	maxAf := 0.2
	sar := normalizedLows[0]
	ep := normalizedHighs[0]
	isLong := true

	for i := 1; i < len(normalizedHighs); i++ {
		if isLong {
			sar = sar + af*(ep-sar)
			if normalizedHighs[i] > ep {
				ep = normalizedHighs[i]
				af = math.Min(af+0.02, maxAf)
			}
			if sar > normalizedLows[i] {
				isLong = false
				sar = ep
				ep = normalizedLows[i]
				af = 0.02
			}
		} else {
			sar = sar - af*(sar-ep)
			if normalizedLows[i] < ep {
				ep = normalizedLows[i]
				af = math.Min(af+0.02, maxAf)
			}
			if sar < normalizedHighs[i] {
				isLong = true
				sar = ep
				ep = normalizedHighs[i]
				af = 0.02
			}
		}
	}
	return sar
}
