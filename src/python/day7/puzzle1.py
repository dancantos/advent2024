from .equation import equation

def isSolveable(target, terms):
    return _isSolveable(target, terms, len(terms)-1)


def _isSolveable(target, terms, head):
    if head == 0:
        return target == terms[head]
    last = terms[head]

    if target % last == 0 and _isSolveable(target//last, terms, head-1):
        return True

    return _isSolveable(target-last, terms[:-1], head-1)

if __name__ == "__main__":
    print(isSolveable(10, [1, 9]))
    print(isSolveable(10, [1, 8]))
    print(isSolveable(10, [1, 10]))
    print(isSolveable(10, [3, 3, 1]))