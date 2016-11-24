def slices(orig_string, n):
    if n > len(orig_string) or n == 0:
        raise ValueError
    result = []
    for i in xrange((len(orig_string) + 1) - n):
        result.append(orig_string[i:i + n])
    return [map(int, el) for el in map(list, result)]
