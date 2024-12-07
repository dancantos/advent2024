from .equation import equation
from math import log10, ceil, pow

def isSolveableWithConcat(eq):
    if eq.target < 0:
        return False

    if len(eq.terms) == 1:
        return eq.target == eq.terms[-1]

    last = eq.terms[-1]

    # try division
    if eq.target % last == 0 and isSolveableWithConcat(equation(eq.target//last, eq.terms[:-1])):
        return True

    # try deconcatenating
    unconcatenated = unconcat(eq.target, last)
    if unconcatenated < eq.target and isSolveableWithConcat(equation(unconcatenated, eq.terms[:-1])):
        return True

    # try subtraction
    return isSolveableWithConcat(equation(eq.target-last, eq.terms[:-1]))

def unconcat(a, b):
    if a == b:
        return 0

    order = int(ceil(log10(b+0.1)))
    exp = int(pow(10, order))
    if a%exp == b:
        return (a-b)/exp
    return a

if __name__ == "__main__":
    print(isSolveableWithConcat(equation(10, [1, 9])))
    print(isSolveableWithConcat(equation(10, [1, 8])))
    print(isSolveableWithConcat(equation(10, [1, 10])))
    print(isSolveableWithConcat(equation(10, [3, 3, 1])))