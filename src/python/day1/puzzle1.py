def sumdiffs(a, b):
    a = sorted(a)
    b = sorted(b)
    sum = 0
    for i in range(len(a)):
        sum += abs(a[i]-b[i])
    return sum

if __name__ == "__main__":
    print(sumdiffs([1, 2], [2, 1]))