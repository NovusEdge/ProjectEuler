from time import time
from functools import reduce
from math import sqrt

def factors(n):
    return reduce(list.__add__, 
            [[n//i, i] for i in range(1, int(sqrt(n))+1) if n%i == 0])

def is_prime(n):
    return len(factors(n)) == 2

if __name__ == "__main__":
    start = time()
    answer = max(filter(is_prime, factors(600851475143)))
    end = time() - start

    print(f"\nAnswer for Problem 3: {answer}")
    print( f"Time Taken          : {end}\n")
