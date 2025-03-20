import sys


class Bank:
    def __init__(self):
        self.accounts = {}

    def DEPOSIT(self, name: str, sum_: int) -> None:
        self.accounts[name] = self.accounts.get(name, 0) + sum_

    def WITHDRAW(self, name: str, sum_: int) -> None:
        self.accounts[name] = self.accounts.get(name, 0) - sum_

    def BALANCE(self, name: str):
        return self.accounts.get(name)

    def TRANSFER(self, name1: str, name2: str, sum_: int) -> None:
        self.accounts[name1] = self.accounts.get(name1, 0) - sum_
        self.accounts[name2] = self.accounts.get(name2, 0) + sum_

    def INCOME(self, p: int) -> None:
        for name, balance in list(self.accounts.items()):
            if balance > 0:
                self.accounts[name] = balance * (100 + p) // 100


def main():
    b = Bank()
    for line in sys.stdin:
        parts = line.split()
        op = parts[0]

        if op == "DEPOSIT":
            name = parts[1]
            sum_ = int(parts[2])
            b.DEPOSIT(name, sum_)
        elif op == "WITHDRAW":
            name = parts[1]
            sum_ = int(parts[2])
            b.WITHDRAW(name, sum_)
        elif op == "BALANCE":
            name = parts[1]
            sum_ = b.BALANCE(name)
            if sum_ is not None:
                print(sum_)
            else:
                print("ERROR")
        elif op == "TRANSFER":
            name1 = parts[1]
            name2 = parts[2]
            sum_ = int(parts[3])
            b.TRANSFER(name1, name2, sum_)
        elif op == "INCOME":
            p = int(parts[1])
            b.INCOME(p)
        else:
            raise RuntimeError(f"unknown operation: {op}")


if __name__ == "__main__":
    main()
