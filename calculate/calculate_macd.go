package calculate

// MACD 계산 개선
func CalculateMACD(prices []float64) (float64, float64) {
	if len(prices) < 26 {
		return 0, 0
	}

	// 모든 가격을 백만 단위로 정규화
	normalizedPrices := make([]float64, len(prices))
	for i := range prices {
		normalizedPrices[i] = prices[i] * PRICE_MULTIPLIER
	}

	ema12 := CalculateEMA(normalizedPrices, 12)
	ema26 := CalculateEMA(normalizedPrices, 26)
	macdLine := (ema12 - ema26) / PRICE_MULTIPLIER

	ema12Slice := CalculateEMASlice(normalizedPrices, 12)
	ema26Slice := CalculateEMASlice(normalizedPrices, 26)
	macdSlice := make([]float64, len(prices))
	for i := range prices {
		macdSlice[i] = (ema12Slice[i] - ema26Slice[i]) / PRICE_MULTIPLIER
	}

	signalLine := CalculateEMA(macdSlice, 9)
	return macdLine, signalLine
}
