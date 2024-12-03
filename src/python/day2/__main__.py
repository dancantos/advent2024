import modules.timer as timer
from .puzzle1 import safe1
from .puzzle2 import safe2
from .input import input

def run():
    print("puzzle1: count safe:", countSafe(input, safe1))
    for _ in timer.run(100):
        countSafe(input, safe1)

    print()

    print("puzzle2: count safe with dampening:", countSafe(input, safe2))
    for _ in timer.run(100):
        countSafe(input, safe2)

def countSafe(input, safe):
    counter = 0
    for v in input:
        if safe(v):
            counter+=1
    return counter

if __name__ == "__main__":
    run()