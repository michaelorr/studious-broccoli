from operator import mul


def largest_product(series, num):
    if not series:
        return 1
    return max(reduce(mul, slice) for slice in slices(series, num))


def slices(series, num):
    res = []
    while len(series) >= num:
        res.append(series[:num])
        series = series[1:]
    if not res:
        raise ValueError
    return [map(int, list(i)) for i in res]
