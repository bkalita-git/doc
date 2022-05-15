import time


def check(input):
    if input == "password":
        return True


def alpha(input):
    start = time.time_ns()
    check(input)
    end = time.time_ns()
    print(end-start)


alpha("aassword")
alpha("bassword")
print("\n\n")

alpha("aassword")
alpha("bassword")
alpha("password")
alpha("cassword")
alpha("dassword")
