package q4

//Dado um array não vazio de números inteiros "nums", cada elemento aparece duas vezes, exceto um. Encontre esse único
//elemento.
//
//Você deve implementar uma solução com complexidade de tempo linear e sem memória extra.

func SingleNumber(nums []int) int {
	Map := make(map[int]int)
	for _, num := range nums {
		Map[num]++
	}
	for i, num := range nums {
		if Map[num] == 1 {
			return i
		}
	}
	return -1
}
