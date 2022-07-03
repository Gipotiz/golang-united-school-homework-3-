package homework

func average(input [15]float32) (result float32) {
	//Place your code here

	var sum float32

	for i := range input {
		sum += input[i]
	}
	result = sum / float32(len(input))

	return
}
