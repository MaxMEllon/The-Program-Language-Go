package popcount

var pc1 [256]byte
var pc2 [256]byte

func init() {
	for i := range pc1 {
		pc1[i] = pc1[i/2] + byte(i&1)
	}
	for i := range pc2 {
		pc2[i] = pc2[i/2] + byte(i&1)
	}
}

func PopCountByVerbose(x uint64) int {
	return int(pc1[byte(x>>(0*8))] +
		pc1[byte(x>>(1*8))] +
		pc1[byte(x>>(2*8))] +
		pc1[byte(x>>(3*8))] +
		pc1[byte(x>>(4*8))] +
		pc1[byte(x>>(5*8))] +
		pc1[byte(x>>(6*8))] +
		pc1[byte(x>>(7*8))])
}

func PopCountByFor(x uint64) int {
	result := 0
	for i := 0; i < 8; i++ {
		result += int(pc2[byte(x>>uint(i*8))])
	}
	return result
}
