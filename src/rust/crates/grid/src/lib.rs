#[derive(Debug)]
struct Bitmask {
    width: u64,
    height: u64,
    mask: Vec<u64>,
}

impl Bitmask {
    fn new(width: u64, height: u64) -> Bitmask {
        let rows: usize = (width*height / 64).try_into().unwrap();
        let mask = vec![0;rows];
        Bitmask { width: width, height: height, mask: mask }
    }

    fn set(&mut self, x: u64, y: u64) {
        let entry: usize = (y*self.width + x).try_into().unwrap();
        let row: usize = entry / 64;
        let col = entry % 64;
        self.mask[row] |= 1<<col;
    }

    fn unset(&mut self, x: u64, y: u64) {
        let entry: usize = (y*self.width + x).try_into().unwrap();
        let row: usize = entry / 64;
        let col = entry % 64;
        self.mask[row] &= !(1<<col);
    }

    fn flip(&mut self, x: u64, y: u64) {
        let entry: usize = (y*self.width + x).try_into().unwrap();
        let row: usize = entry / 64;
        let col = entry % 64;
        self.mask[row] ^= 1<<col;
    }

    fn is_set(&self, x: u64, y: u64) -> bool {
        let entry: usize = (y*self.width + x).try_into().unwrap();
        let row: usize = entry / 64;
        let col = entry % 64;
        return (self.mask[row] & 1<<col) > 0;
    }

    fn inside(&self, x: u64, y: u64) -> bool {
        x < self.width && y < self.height
    }
}


#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn set_works() {
        let mut m = Bitmask::new(10, 10);
        m.set(1, 1);
        m.set(2, 2);
        m.set(9, 7);
        assert!(m.is_set(1, 1));
        assert!(m.is_set(2, 2));
        assert!(m.is_set(9, 7));
        assert!(!m.is_set(1, 2));
        assert!(!m.is_set(9, 8));
    }

    #[test]
    fn unset_works() {
        let mut m = Bitmask::new(10, 10);
        m.set(1, 1);
        m.set(2, 2);
        m.set(9, 7);
        m.set(9, 9);
        m.unset(1, 1);
        m.unset(9, 7);
        m.unset(9, 8);
        assert!(!m.is_set(1, 1));
        assert!(!m.is_set(9, 7));
        assert!(!m.is_set(9, 8));
        assert!(m.is_set(2, 2));
        assert!(m.is_set(9, 9));
    }

    #[test]
    fn flip_works() {
        let mut m = Bitmask::new(10, 10);
        m.set(1, 1);
        m.set(9, 7);
        m.flip(1, 1);
        m.flip(1, 2);
        m.flip(9, 7);
        m.flip(9, 8);
        assert!(!m.is_set(1, 1));
        assert!(m.is_set(1, 2));
        assert!(!m.is_set(9, 7));
        assert!(m.is_set(9, 8));
    }

}
