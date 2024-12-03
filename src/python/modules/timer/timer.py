from time import perf_counter_ns
import math

# timer code courtesy of baely (https://github.com/baely)
def run(n: int) -> iter:
    times = [perf_counter_ns()]
    for i in range(n):
        yield i
        times.append(perf_counter_ns())
    deltas = [b - a for a, b in zip(times[:-1], times[1:])]
    precision = min(3 * math.floor(math.log(min(deltas)) / math.log(10) / 3), 9)
    unit = ["ns", "μs", "ms", "s"][precision//3]
    print("=== Performance Stats ===")
    print(f" min:  {min(deltas)/10**precision:-10.3f}{unit}")
    print(f" max:  {max(deltas)/10**precision:-10.3f}{unit}")
    print(f" mean: {sum(deltas) / len(deltas)/10**precision:-10.3f}{unit}")