import modules.timer as timer
from .puzzle1 import mulall
from .puzzle2 import mulWithInstruction
from .input import input

def run():
    print("puzzle1: mulall:", mulall(input))
    for _ in timer.run(100):
        mulall(input)

    print()

    print("puzzle2: mulWithInstruction:", mulWithInstruction(input))
    for _ in timer.run(100):
        mulWithInstruction(input)

if __name__ == "__main__":
    run()