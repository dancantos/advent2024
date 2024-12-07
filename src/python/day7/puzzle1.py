from .equation import equation

def isSolveable(eq):
    if len(eq.terms) == 1:
        return eq.target == eq.terms[-1]
    last = eq.terms[-1]

    if eq.target % last == 0 and isSolveable(equation(eq.target//last, eq.terms[:-1])):
        return True

    return isSolveable(equation(eq.target-last, eq.terms[:-1]))

if __name__ == "__main__":
    print(isSolveable(equation(10, [1, 9])))
    print(isSolveable(equation(10, [1, 8])))
    print(isSolveable(equation(10, [1, 10])))
    print(isSolveable(equation(10, [3, 3, 1])))