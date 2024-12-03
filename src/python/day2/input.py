import os

def readInput(text):
    result = []
    for line in text:
        if len(line) == 0:
            continue
        lineVals = []
        for v in line.split():
            lineVals.append(int(v))
        result.append(lineVals)
    return result

def readFile(file):
    with open(file) as f:
        return readInput(f.readlines())

input = readFile(os.path.join(os.path.dirname(__file__), "input.txt"))

if __name__ == "__main__":
    print(input)