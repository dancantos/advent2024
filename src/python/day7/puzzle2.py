from .equation import equation
from math import log10, ceil, pow

def isSolveableWithConcat(target, terms):
    return _isSolveableWithConcat(target, terms, len(terms)-1)

def _isSolveableWithConcat(target, terms, head):
    if target < 0:
        return False

    if head == 0:
        return target == terms[head]

    last = terms[head]

    # try division
    if target % last == 0 and _isSolveableWithConcat(target//last, terms, head-1):
        return True

    # try deconcatenating
    unconcatenated = unconcat(target, last)
    if unconcatenated < target and _isSolveableWithConcat(unconcatenated, terms, head-1):
        return True

    # try subtraction
    return _isSolveableWithConcat(target-last, terms, head-1)

def unconcat(a, b):
    if a == b:
        return 0

    order = int(ceil(log10(b+0.1)))
    exp = int(pow(10, order))
    if a%exp == b:
        return (a-b)/exp
    return a

if __name__ == "__main__":
    print(isSolveableWithConcat(10, [1, 9]))
    print(isSolveableWithConcat(10, [1, 8]))
    print(isSolveableWithConcat(10, [1, 10]))
    print(isSolveableWithConcat(10, [3, 3, 1]))