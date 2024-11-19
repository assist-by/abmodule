package calculate

// 정규화된 가격으로 MACD 계산
func CalculateMACD(prices []float64) ([]float64, []float64) {
	if len(prices) < 26 {
		return nil, nil // Not enough data
	}

	// 전체 기간에 대한 EMA 계산
	ema12Slice := CalculateEMASlice(prices, 12)
	ema26Slice := CalculateEMASlice(prices, 26)

	// MACD Line 계산
	macdLine := make([]float64, len(prices))
	for i := 0; i < len(prices); i++ {
		macdLine[i] = ema12Slice[i] - ema26Slice[i]
	}

	// Signal Line 계산 (9일 EMA of MACD Line)
	signalLine := CalculateEMASlice(macdLine, 9)

	return macdLine, signalLine
}
