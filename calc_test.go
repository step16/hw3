package main

import "testing"

func TestReadNumber(t *testing.T) {
	for _, tc := range []struct {
		line       string
		index      int
		wantNumber float64
		wantIndex  int
	}{
		{
			line:       "42",
			wantNumber: 42,
			wantIndex:  2,
		},
		{
			line:       "3.1415926535897932384626433832795028841971",
			wantNumber: 3.1415926535897932384626433832795028841971,
			wantIndex:  42,
		},
		{
			line:       "6*7",
			wantNumber: 6,
			wantIndex:  1,
		},
		{
			line:       "6*7",
			index:      2,
			wantNumber: 7,
			wantIndex:  3,
		},
	} {
		wantToken := token{
			kind:   kindNumber,
			number: tc.wantNumber,
		}
		if gotToken, gotIndex := readNumber(tc.line, tc.index); gotToken != wantToken || gotIndex != tc.wantIndex {
			t.Errorf("readNumber(%q) = (%v, %q) want (%v, %q)", tc.line, gotToken, gotIndex, wantToken, tc.wantIndex)
		}
	}
}

func approxEqual(a, b float64) bool {
	const epsilon = 1e-5
	return a-b < epsilon && b-a < epsilon
}

func TestEvalLine(t *testing.T) {
	for _, tc := range []struct {
		in   string
		want float64
	}{
		{
			in:   "42",
			want: 42,
		},
		{
			in:   "3.14159",
			want: 3.14159,
		},
		{
			in:   "3.2+1",
			want: 4.2,
		},
		{
			in:   "5.3-1.1",
			want: 4.2,
		},
		{
			in:   "41+1",
			want: 42,
		},
	} {
		if got := evalLine(tc.in); !approxEqual(got, tc.want) {
			t.Errorf("evalLine(%q) = %v want %v", tc.in, got, tc.want)
		}
	}
}
