package calculate

// 정규화된 가격으로 EMA 계산
func CalculateEMA(prices []float64, period int) float64 {
	normalizedPrices := make([]float64, len(prices))
	for i, p := range prices {
		normalizedPrices[i] = normalizePrice(p)
	}

	k := 2.0 / float64(period+1)
	ema := normalizedPrices[0]
	for i := 1; i < len(normalizedPrices); i++ {
		ema = normalizedPrices[i]*k + ema*(1-k)
	}
	return ema
}

// 정규화된 가격으로 EMA 슬라이스 계산
func CalculateEMASlice(prices []float64, period int) []float64 {
	normalizedPrices := make([]float64, len(prices))
	for i, p := range prices {
		normalizedPrices[i] = normalizePrice(p)
	}

	k := 2.0 / float64(period+1)
	ema := make([]float64, len(normalizedPrices))
	ema[0] = normalizedPrices[0]
	for i := 1; i < len(normalizedPrices); i++ {
		ema[i] = normalizedPrices[i]*k + ema[i-1]*(1-k)
	}
	return ema
}
