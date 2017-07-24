package protein

const testVersion = 1

var codonMapping = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func FromCodon(codon string) string {
	return codonMapping[codon]
}

func FromRNA(rna string) (proteins []string) {
	return translateEach(splitCodons(rna))
}

func splitCodons(rna string) (codons []string) {
	runes := []rune(rna)
	for i := 0; i < len(runes); i += 3 {
		codons = append(codons, string(runes[i:i+3]))
	}
	return
}

func translateEach(codons []string) (proteins []string) {
	for _, codon := range codons {
		protein := FromCodon(codon)
		if protein == "STOP" {
			break
		}
		proteins = append(proteins, protein)
	}
	return
}
