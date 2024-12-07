from .equation import equation
from .input import input
from .puzzle1 import isSolveable
from .puzzle2 import isSolveableWithConcat
import modules.timer as timer


def main():
    # min:2.892ms     max:3.239ms     mean:3.055ms
    print(f"Puzzle 1 (solveable with +*): {countSolveable(input, isSolveable)}")
    for _ in timer.run(100):
        countSolveable(input, isSolveable)

    print()

    # min:5.937ms     max:6.683ms     mean:6.195ms
    print(f"Puzzle 1 (solveable with +*||): {countSolveable(input, isSolveableWithConcat)}")
    for _ in timer.run(100):
        countSolveable(input, isSolveableWithConcat)

def countSolveable(equations, filter):
    count = 0
    for eq in equations:
        if filter(eq.target, eq.terms):
            count += eq.target
    return count

if __name__ == "__main__":
    main()
