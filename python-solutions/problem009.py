from time import time

def compute():
    for a in range(1, 500):
        for b in range(1, 500):
            c = 1000 - (a + b)
            if a**2 + b**2 == c**2:
                return a*b*c

if __name__ == "__main__":
    start = time()
    answer = compute()
    end = time() - start

    print(f"\nAnswer for Problem 8: {answer}")
    print( f"Time Taken: {end} seconds\n")
