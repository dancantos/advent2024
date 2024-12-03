import timeit
from input import a1, a2
from puzzle1 import sumdiffs
from puzzle2 import similarityScore

def run():
    print(sumdiffs(a1, a2))
    print(similarityScore(a1, a2))
    diffsTime = timeit.timeit("sumdiffs(a1, a2)","from __main__ import sumdiffs, a1, a2", number=10)
    print("sumdiffs time: ", diffsTime/10*1000, "ms")
    simTime = timeit.timeit("similarityScore(a1, a2)","from __main__ import similarityScore, a1, a2", number=10)
    print("similarityScore time: ", simTime/10*1000, "ms")
    return

if __name__ == "__main__":
    run()