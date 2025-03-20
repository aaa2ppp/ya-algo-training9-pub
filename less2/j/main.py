import sys

def solve(g, s, m, n):
    gf = [0] * 64
    for i in range(m):
        gf[g[i] - 64] += 1

    sf = [0] * 64
    cnt = 0
    for v in gf:
        if v == 0:
            cnt += 1

    ans = 0
    # первое окно
    for i in range(m):
        idx = s[i] - 64
        if sf[idx] == gf[idx]:
            cnt -= 1
        sf[idx] += 1
        if sf[idx] == gf[idx]:
            cnt += 1
    if cnt == 64:
        ans += 1

    # скользящее окно
    for r in range(m, n):
        l = r - m
        # удаление левого символа
        idx = s[l] - 64
        if sf[idx] == gf[idx]:
            cnt -= 1
        sf[idx] -= 1
        if sf[idx] == gf[idx]:
            cnt += 1
        # добавление правого символа
        idx = s[r] - 64
        if sf[idx] == gf[idx]:
            cnt -= 1
        sf[idx] += 1
        if sf[idx] == gf[idx]:
            cnt += 1
        if cnt == 64:
            ans += 1

    return ans

def main():
    data = sys.stdin.buffer.readline().split()
    if not data:
        return
    m, n = map(int, data)
    g = sys.stdin.buffer.readline()
    s = sys.stdin.buffer.readline()
    print(solve(g, s, m, n))

if __name__ == "__main__":
    main()
