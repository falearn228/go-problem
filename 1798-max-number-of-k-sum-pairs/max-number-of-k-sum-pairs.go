func maxOperations(nums []int, k int) int {
    freq := make(map[int]int)
    operations := 0

    
    for _, num := range nums {
        complement := k - num
        
        // Проверяем, есть ли для нас пара в карте
        if freq[complement] > 0 {
            // Нашли!
            operations++
            // Уменьшаем счетчик пары, так как мы ее "использовали"
            freq[complement]--
        } else {
            // Пары нет, добавляем себя в карту в ожидании своей пары
            freq[num]++
        }
    }
    return operations
}