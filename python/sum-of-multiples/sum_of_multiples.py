def sum_of_multiples(max, mults=None):
    if not mults:
        mults = [3, 5]

    return sum(i for i in range(max) if is_mult(i, mults))


def is_mult(i, mults):
    for n in mults:
        if n:
            if i % n == 0:
                return True
    return False
