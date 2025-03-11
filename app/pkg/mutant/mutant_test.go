package mutant

import "testing"


func TestIsMutant(t *testing.T) {
	tests := []struct {
		name string
		dna  []string
		want bool
	}{
		{
			name: "Non-mutant DNA (no sequences)",
			dna: []string{
				"ATGCGA",
				"CAGTGC",
				"TTATTT",
				"AGACGG",
				"GCGTCA",
				"TCACTG",
			},
			want: false,
		},
		
		{
			name: "Mutant DNA (horizontal sequence)",
			dna: []string{
				"AAAA",
				"GTCG",
				"GTCG",
				"GTCG",
			},
			want: true,
		},
		{
			name: "Mutant DNA (vertical sequence)",
			dna: []string{
				"ATCG",
				"ATCG",
				"ATCG",
				"ATCG",
			},
			want: true,
		},
		{
			name: "Mutant DNA (diagonal sequence)",
			dna: []string{
				"ATCG",
				"GATC",
				"CGAT",
				"TCGA",
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsMutant(tt.dna)
			if got != tt.want {
				t.Errorf("IsMutant() = %v, want %v", got, tt.want)
			}
		})
	}
}