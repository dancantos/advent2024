def safe1(arr):
    if len(arr) < 2:
        return True
    return _safe1(arr, 1) or _safe1(arr, -1)

def _safe1(arr, direction):
    for i in range(len(arr)-1):
        diff = direction*(arr[i+1]-arr[i])
        if diff < 1 or 3 < diff:
            return False
    return True

if __name__ == "__main__":
    print(safe1([1, 2, 3]))
    print(safe1([1, 2, 5]))
    print(safe1([1, 2, 6]))
    print(safe1([1, 2, 1]))
    print(safe1([3, 2, 1]))
    print(safe1([1, 2, 4, 3]))
    print(safe1([1, 2, 4, 10]))