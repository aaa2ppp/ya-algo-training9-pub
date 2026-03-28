import sys

MOD = 1_000_000_007

def solve(a):
    n = len(a)
    if n < 3:
        return 0

    b = [0] * (n + 1)
    for i, val in enumerate(a):
        b[i + 1] = (b[i] + val) % MOD

    ans = 0
    for i in range(1, n - 1):
        v = a[i]
        v = (v * b[i]) % MOD
        v = (v * ((b[n] - b[i + 1]))) % MOD
        ans = (ans + v) % MOD

    return ans


def main():
    n = int(input())
    a = list(map(int, input().split()))
    ans = solve(a)
    print(ans)


if __name__ == "__main__":
    main()
