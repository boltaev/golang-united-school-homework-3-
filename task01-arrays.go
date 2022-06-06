package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	var sum float32
	for i := 0; i < len(input); i++ {
		sum += input[i]
	}
	lenOfArray := float32(len(input))
	result = sum / lenOfArray
	return result
}
