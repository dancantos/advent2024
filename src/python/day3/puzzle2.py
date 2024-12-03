import re

mulre = re.compile("(?:mul\(([1-9][0-9]{0,2}),([1-9][0-9]{0,2})\)|do\(\)|don't\(\))")

def mulWithInstruction(text):
    sum = 0
    active = True
    for m in mulre.finditer(text):
        if m.group(0) == "do()":
            active = True
            continue
        if m.group(0) == "don't()":
            active = False
            continue
        if active:
            sum += int(m.group(1))*int(m.group(2))
    return sum

if __name__ == "__main__":
    print(mulWithInstruction("""mul(1,2)xxdo()45hhdon't()ksmul(4,5)addo()77mul(2,2)k890w;"""))
