from collections import defaultdict

def similarityScore(a, b):
    counts = defaultdict(int)
    for n in b:
        counts[n]+=1
    return sum([n*counts[n] for n in a if counts[n]])

if __name__ == "__main__":
    print(similarityScore([1, 2], [2, 1]))