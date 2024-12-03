import os

def readInput(text):
    a1 = []
    a2 = []
    for line in text:
        nums = line.split()
        a1.append(int(nums[0]))
        a2.append(int(nums[1]))
    return a1, a2

def readFile(file):
    with open(file) as f:
        return readInput(f.readlines())

a1, a2 = readFile(os.path.join(os.path.dirname(__file__), "input.txt"))

if __name__ == "__main__":
    print(a1, a2)