package main

import "testing"

func TestSearch(t *testing.T) {
	type args struct {
		numbers        []int
		numberToSearch int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Success Scenario", args{numbers: []int{1,4,5,7,9,10}, numberToSearch: 7}, true},
		{"Failure Scenario", args{numbers: []int{1,4,5,7,9,10}, numberToSearch: 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Search(tt.args.numbers, tt.args.numberToSearch); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
