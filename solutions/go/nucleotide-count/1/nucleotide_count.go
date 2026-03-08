package nucleotidecount

import "errors"

type DNA string

type Histogram map[rune]int

func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, element := range d {
        switch element {
            case 'A':
            	h['A'] += 1
            case 'C':
            	h['C'] += 1
            case 'G':
            	h['G'] += 1
            case 'T':
            	h['T'] += 1
            default:
            	return h, errors.New("invalid nucleotides")
        }
    }
	return h, nil
}
