package calculate

/// MACD 계산
func CalculateMACD(prices []float64) (float64, float64) {
	if len(prices) < 26 {
		return 0, 0
	}

	// 가격이 1보다 작으면 1000을 곱해서 스케일업
	scaleFactor := 1.0
	if prices[0] < 1.0 {
		scaleFactor = 1000.0
	}

	scaledPrices := make([]float64, len(prices))
	for i := 0; i < len(prices); i++ {
		scaledPrices[i] = prices[i] * scaleFactor
	}

	ema12 := CalculateEMA(scaledPrices, 12)
	ema26 := CalculateEMA(scaledPrices, 26)
	macdLine := (ema12 - ema26) / scaleFactor // 결과를 다시 원래 스케일로

	ema12Slice := CalculateEMASlice(scaledPrices, 12)
	ema26Slice := CalculateEMASlice(scaledPrices, 26)
	macdSlice := make([]float64, len(prices))
	for i := 0; i < len(prices); i++ {
		macdSlice[i] = (ema12Slice[i] - ema26Slice[i]) / scaleFactor
	}

	signalLine := CalculateEMA(macdSlice, 9)
	return macdLine, signalLine
}
