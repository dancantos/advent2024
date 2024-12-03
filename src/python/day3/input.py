import os

def readFile(file):
    text = ""
    with open(file) as f:
        for line in f:
            text += line
    return text

input = readFile(os.path.join(os.path.dirname(__file__), "input.txt"))

if __name__ == "__main__":
    print(input)