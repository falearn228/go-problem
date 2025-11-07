// Last updated: 11/7/2025, 2:47:36 PM
func canPlaceFlowers(flowerbed []int, n int) bool {
    if n == 0 {
        return true
    }
    if len(flowerbed) == 1 && n == 1 {
        if flowerbed[0] == 1 {
            return false
        }
        return true
    }
	for i, v := range flowerbed {
		if v == 0 {
			if i == 0 && flowerbed[i+1] != 1 {
				flowerbed[i] = 1
				n--
			} else if i == len(flowerbed)-1 && flowerbed[i-1] != 1 {
				flowerbed[i] = 1
				n--

			} else if i > 0 && i < len(flowerbed) - 1 && flowerbed[i-1] != 1 && flowerbed[i+1] != 1 {
				flowerbed[i] = 1
				n--
			}
		}
		if n == 0 {
			return true
		}
	}
	return false
}