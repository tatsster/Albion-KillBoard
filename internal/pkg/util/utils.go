package util

import "strconv"

func If[T any](cond bool, vtrue, vfalse T) T {
	if cond {
		return vtrue
	}
	return vfalse

}

func FormatInt(number int) string {
	output := strconv.Itoa(number)
	startOffset := 3
	if number < 0 {
		startOffset++
	}
	for outputIndex := len(output); outputIndex > startOffset; {
		outputIndex -= 3
		output = output[:outputIndex] + "." + output[outputIndex:]
	}
	return output
}
