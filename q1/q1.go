package q1

//Dado um array de números inteiros "nums" e um número inteiro "target", retorne os índices dos dois números cuja soma
//seja igual a "target".
//
//Você pode assumir que cada entrada terá exatamente uma solução e não poderá usar o mesmo elemento duas vezes.
//
//Você pode retornar a resposta em qualquer ordem.

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if _, ok := m[complement]; ok {
			return []int{m[complement], i}
		}
		m[num] = i
	}
	return nil
}
