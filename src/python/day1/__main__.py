import modules.timer as timer
from .input import a1, a2
from .puzzle1 import sumdiffs
from .puzzle2 import similarityScore

def run():
    print("puzzle1: sumdiffs: ", sumdiffs(a1, a2))
    for _ in timer.run(100):
        sumdiffs(a1, a2)

    print()

    print("puzzle2: similarity: ", similarityScore(a1, a2))
    for _ in timer.run(100):
        similarityScore(a1, a2)

if __name__ == "__main__":
    run()