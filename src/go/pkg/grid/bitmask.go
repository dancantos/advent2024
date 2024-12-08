package grid

// Bitmask represents a grid of numbers where each cell can be set to true or false.
type Bitmask struct {
	width, height int
	mask          []uint64
}

// NewBitmask creates a bitmask representing a grid of numbers.
func NewBitmask(width, height int) Bitmask {
	return Bitmask{
		width:  width,
		height: height,
		mask:   make([]uint64, width*height/64+1),
	}
}

// IsSet returns true if a bit is set.
func (b Bitmask) IsSet(x, y int) bool {
	mapped := y*b.width + x
	row, col := mapped/64, uint64(1)<<(mapped%64)
	return b.mask[row]&col > 0
}

// Set sets a grid bit, returning true if the bit was flipped to true.
func (b Bitmask) Set(x, y int) bool {
	row, col := b.rowcol(x, y)
	wasFalse := b.mask[row]&col == 0
	b.mask[row] |= col
	return wasFalse
}

// Flip flips a grid bit, returning true if the bit was flipped to true.
func (b Bitmask) Flip(x, y int) bool {
	row, col := b.rowcol(x, y)
	wasFalse := b.mask[row]&col == 0
	b.mask[row] ^= col
	return wasFalse
}

// Unset unsets a grid bit, returning true if the bit was flipped to false.
func (b Bitmask) Unset(x, y int) bool {
	row, col := b.rowcol(x, y)
	wasEmpty := b.mask[row]&col == 0
	b.mask[row] &= ^col
	return !wasEmpty
}

// CountSet counts the number of set bits in the mask
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
