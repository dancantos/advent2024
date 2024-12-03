import re

mulre = re.compile("mul\(([1-9][0-9]{0,2}),([1-9][0-9]{0,2})\)")

def mulall(text):
    sum = 0
    for m in mulre.finditer(text):
        sum += int(m.groups(0)[0])*int(m.groups(0)[1])
    return sum

def mulall_oneliner(text):
    # cool comprehension but allocates more memory
    return sum([ int(m.groups(0)[0])*int(m.groups(0)[1]) for m in mulre.finditer(text) ])

if __name__ == "__main__":
    print(mulall("""mul(1,2)xx45hhksmul(4,5)adk890w;"""))
    print(mulall_oneliner("""mul(1,2)xx45hhksmul(4,5)adk890w;"""))
