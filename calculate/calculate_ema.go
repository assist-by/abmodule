package calculate

// normalization을 위한 상수
const PRICE_MULTIPLIER = 1000000.0 // 백만 단위로 통일

// SMA 계산 함수 추가
func calculateSMA(prices []float64, period int) float64 {
	sum := 0.0
	for i := 0; i < period; i++ {
		sum += prices[i]
	}
	return sum / float64(period)
}

// EMA 계산 개선
func CalculateEMA(prices []float64, period int) float64 {
	// 모든 가격을 백만 단위로 정규화
	normalizedPrices := make([]float64, len(prices))
	for i := range prices {
		normalizedPrices[i] = prices[i] * PRICE_MULTIPLIER
	}

	k := 2.0 / float64(period+1)
	ema := calculateSMA(normalizedPrices, period) // SMA로 초기값 설정

	for i := period; i < len(normalizedPrices); i++ {
		ema = normalizedPrices[i]*k + ema*(1-k)
	}

	return ema / PRICE_MULTIPLIER // 원래 스케일로 변환
}

// EMASlice 계산 개선
func CalculateEMASlice(prices []float64, period int) []float64 {
	normalizedPrices := make([]float64, len(prices))
	for i := range prices {
		normalizedPrices[i] = prices[i] * PRICE_MULTIPLIER
	}

	k := 2.0 / float64(period+1)
	ema := make([]float64, len(prices))
	ema[0] = calculateSMA(normalizedPrices, period)

	for i := 1; i < len(normalizedPrices); i++ {
		ema[i] = normalizedPrices[i]*k + ema[i-1]*(1-k)
	}

	// 결과값 정규화 해제
	result := make([]float64, len(prices))
	for i := range ema {
		result[i] = ema[i] / PRICE_MULTIPLIER
	}

	return result
}
