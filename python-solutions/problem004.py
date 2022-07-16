from time import time

def is_pallindrome(n):
    return str(n) == str(n)[::-1]

def compute():
    l = []
    for i in range(999, 100, -1):
        for j in range(999, 100, -1):
            if is_pallindrome(i*j):
                l.append(i*j)

    return max(l)


if __name__ == "__main__":
    start = time()
    answer = compute()
    end = time() - start

    print(f"\nAnswer for Problem 4: {answer}")
    print( f"Time Taken: {end} seconds\n")
