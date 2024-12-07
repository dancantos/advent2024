from .equation import equation
from .input import input
from .puzzle1 import isSolveable
from .puzzle2 import isSolveableWithConcat
import modules.timer as timer


def main():
    # min:6.993ms     max:31.727ms     mean:7.772ms
    print(f"Puzzle 1 (solveable with +*): {countSolveable(input, isSolveable)}")
    for _ in timer.run(100):
        countSolveable(input, isSolveable)

    print()

    # min:11.098ms     max:12.658ms     mean:11.471ms
    print(f"Puzzle 1 (solveable with +*||): {countSolveable(input, isSolveableWithConcat)}")
    for _ in timer.run(100):
        countSolveable(input, isSolveableWithConcat)

def countSolveable(equations, filter):
    count = 0
    for eq in equations:
        if filter(eq):
            count += eq.target
    return count

if __name__ == "__main__":
    main()
