package calculate

// 가격 정규화 함수
func normalizePrice(price float64) float64 {
	if price < 1 {
		return price * 100 // 작은 가격을 증폭
	}
	return price
}
