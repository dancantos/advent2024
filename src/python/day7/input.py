from .equation import equation
import os

def readFile(file):
    equations = []
    with open(file) as f:
        for line in f:
            eq = line.split()
            equations.append(equation(
                int(eq[0][:-1]),
                list(map(int, eq[1:]))
            ))
    return equations

input = readFile(os.path.join(os.path.dirname(__file__), "input.txt"))

if __name__ == "__main__":
    print(input)