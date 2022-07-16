from time import time
from functools import reduce
from math import gcd


if __name__ == "__main__":
    start = time()

    answer = int(reduce(
        lambda a, b: a * b / gcd(int(a), 
        int(b)), range(1, 21)
    ))

    end = time() - start

    print(f"\nAnswer for Problem 4: {answer}")
    print( f"Time Taken: {end} seconds\n")
