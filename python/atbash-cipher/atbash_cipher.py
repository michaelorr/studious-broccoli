from string import maketrans, translate


alpha = 'abcdefghijklmnopqrstuvwxyz'
rev = alpha[::-1] + '1234567890'
alpha += '1234567890'


def decode(text):
    table = maketrans(rev, alpha)
    text = text.lower().replace(" ", "")
    return translate(text, table)


def encode(text):
    table = maketrans(alpha, rev)
    text = ''.join(ch for ch in text.lower() if ch in alpha)
    cipher = translate(text, table)
    return ' '.join(cipher[i:i + 5] for i in range(0, len(cipher), 5))
