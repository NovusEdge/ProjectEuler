from time import time
from math import sqrt

def F(n):
    return ((1+sqrt(5))**n-(1-sqrt(5))**n)/(2**n*sqrt(5))

def solution():
    i, sum = 1, 0
    while True:
        t = int(F(i))
        if t > 4_000_000: break
        if t%2 == 0: 
            sum += t
        i += 1
        
    return sum



if __name__ == "__main__":
    start  = time()
    answer = solution()
    print(f"\nAnswer to Problem 2: {answer}")
    print(  f"Time Taken: {time()-start} seconds\n")
