def to_rna(dna):
    rna = ''
    mapping = {'G': 'C', 'C': 'G', 'T': 'A', 'A': 'U'}
    for nucleotide in dna:
        rna += mapping[nucleotide]
    return rna
