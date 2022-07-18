from time import time


def prime_sieve(n):
    size = n//2
    sieve = [1]*size
    limit = int(n**0.5)
    for i in range(1,limit):
        if sieve[i]:
            val = 2*i+1
            tmp = ((size-1) - i)//val 
            sieve[i+val::val] = [0]*tmp
    return [2] + [i*2+1 for i, v in enumerate(sieve) if v and i>0]


def compute():
    return prime_sieve(10000000)[10000] 

if __name__ == "__main__":
    start = time()
    answer = compute()
    end = time() - start

    print(f"\nAnswer for Problem 7: {answer}")
    print( f"Time Taken: {end} seconds\n")
