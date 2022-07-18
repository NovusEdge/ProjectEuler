from time import time

def sum_of_squares(n):
    return n * (n + 1) * (2 * n + 1) / 6

def square_of_sum(n):
    return (n * (n + 1) / 2) ** 2

def compute():
    return square_of_sum(100) - sum_of_squares(100)

if __name__ == "__main__":
    start = time()
    answer = compute()
    end = time() - start

    print(f"\nAnswer for Problem 6: {answer}")
    print( f"Time Taken: {end} seconds\n")
