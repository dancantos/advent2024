package grid

type Bitmask struct {
	width, height int
	mask          []uint64
}

func NewBitmask(width, height int) Bitmask {
	return Bitmask{
		width:  width,
		height: height,
		mask:   make([]uint64, width*height/64+1),
	}
}

func (b Bitmask) IsSet(x, y int) bool {
	mapped := y*b.width + x
	row, col := mapped/64, uint64(1)<<(mapped%64)
	return b.mask[row]&col > 0
}

func (b Bitmask) Set(x, y int) {
	row, col := b.rowcol(x, y)
	b.mask[row] |= col
}

func (b Bitmask) Flip(x, y int) {
	row, col := b.rowcol(x, y)
	b.mask[row] ^= col
}

func (b Bitmask) Unset(x, y int) {
	row, col := b.rowcol(x, y)
	b.mask[row] &= ^col
}

func (b Bitmask) CountSet() int {
	count := 0
	for _, mask := range b.mask {
		count += countBits(mask)
	}
	return count
}

func countBits(n uint64) int {
	var count int
	for count = 0; n != 0; count++ {
		n &= n - 1
	}
	return count
}

func (b Bitmask) rowcol(x, y int) (int, uint64) {
	mapped := y*b.width + x
	return mapped / 64, uint64(1) << (mapped % 64)
}
