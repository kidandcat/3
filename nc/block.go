package nc

import ()

// Block is a [][][]float32 with square layout and contiguous underlying storage:
// 	len(block[0]) == len(block[1]) == ...
// 	len(block[i][0]) == len(block[i][1]) == ...
// 	len(block[i][j][0]) == len(block[i][j][1]) == ...
type Block [][][]float32

// Make a block of float32's of size N[0] x N[1] x N[2]
// with contiguous underlying storage.
func MakeBlock(size [3]int) Block {
	checkSize(size[:])
	sliced := make([][][]float32, size[0])
	for i := range sliced {
		sliced[i] = make([][]float32, size[1])
	}
	storage := make([]float32, size[0]*size[1]*size[2])
	for i := range sliced {
		for j := range sliced[i] {
			sliced[i][j] = storage[(i*size[1]+j)*size[2]+0 : (i*size[1]+j)*size[2]+size[2]]
		}
	}
	return Block(sliced)
}

// Total number of scalar elements.
func (b Block) NFloat() int {
	return len(b) * len(b[0]) * len(b[0][0])
}

// BlockSize is the size of the block (N0, N1, N2)
// as was passed to MakeBlock()
func (b Block) BlockSize() [3]int {
	return [3]int{len(b), len(b[0]), len(b[0][0])}
}

// Returns the contiguous underlying storage.
func (b Block) Contiguous() []float32 {
	return ([][][]float32)(b)[0][0][:b.NFloat()]
}
