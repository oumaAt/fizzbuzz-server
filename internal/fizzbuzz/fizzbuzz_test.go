package fizzbuzz

import (
	"errors"
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name string
		int1, int2 int
		limit int
		str1, str2 string
		want []string
		wantErr error
	} {
		{
			name: "fizzbuzz up to 15",
			int1: 3, int2 : 5, limit: 15,
			str1: "fizz", str2: "buzz",
			want: []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz"},
		},
		{
			name: "tictac example",
			int1: 2, int2: 3, limit: 6,
			str1: "tic", str2: "tac",
			want: []string{"1", "tic", "tac", "tic", "5", "tictac"},
		},
		{
			name: "limit of 1",
			int1: 3, int2: 5, limit: 1,
			str1: "fizz", str2: "buzz",
			want: []string{"1"},
		},
				{
			name: "identical divisors",
			int1: 2, int2: 2, limit: 4,
			str1: "a", str2: "b",
			want: []string{"1", "ab", "3", "ab"},
		},
		{
			name: "int1 is zero",
			int1: 0, int2: 5, limit: 10,
			str1: "fizz", str2: "buzz",
			wantErr: ErrInvalidDivisor,
		},
		{
			name: "int2 is negative",
			int1: 3, int2: -5, limit: 10,
			str1: "fizz", str2: "buzz",
			wantErr: ErrInvalidDivisor,
		},
		{
			name: "limit is zero",
			int1: 3, int2: 5, limit: 0,
			str1: "fizz", str2: "buzz",
			wantErr: ErrInvalidLimit,
		},
		{
			name: "limit is negative",
			int1: 3, int2: 5, limit: -1,
			str1: "fizz", str2: "buzz",
			wantErr: ErrInvalidLimit,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err:= Generate(tt.int1, tt.int2, tt.limit, tt.str1, tt.str2)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("error = %v, want %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		}) 
	}
}