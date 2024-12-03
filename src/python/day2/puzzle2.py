def safe2(arr):
    if len(arr) < 3:
        return True
    return _safe2(arr, 1) or _safe2(arr, -1)

def _safe2(arr, direction, dampened=False):
    for i in range(len(arr)-1):
        diff = direction*(arr[i+1]-arr[i])
        if 1 <= diff <= 3:
            continue
        if dampened:
            return False
        # cut out either i or i+1 by repeating on i-1
        if i == 0:
            return _safe2(arr[1:], direction, dampened=True) or _safe2([arr[0]] + arr[2:], direction, dampened=True)
        return _safe2(arr[:i] + arr[i+1:], direction, dampened=True) or _safe2(arr[:i+1] + arr[i+2:], direction, dampened=True)
    return True

if __name__ == "__main__":
    print(safe2([1, 2, 3]))
    print(safe2([1, 2, 5]))
    print(safe2([1, 2, 6]))
    print(safe2([1, 2, 1]))
    print(safe2([3, 2, 1]))
    print(safe2([1, 2, 4, 3]))
    print(safe2([1, 2, 4, 10]))
    print(safe2([1, 2, 2, 1]))
    print(safe2([4, 1, 2, 6]))
    print(safe2([4, 1, 2, 3]))