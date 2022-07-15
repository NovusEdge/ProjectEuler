from time import time

def solution():
    return sum(filter(lambda x: x%3 == 0 or x%5 == 0, range(0, 1000)))

if __name__ == "__main__":
    start  = time()
    answer = solution()
    print(f"\nAnswer to Problem 1: {answer}")
    print(  f"Time Taken: {time()-start} seconds\n")
