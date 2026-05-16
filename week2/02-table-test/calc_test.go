package calc

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{3, 4, 7},
		{-5, 9, 4},
		{0, 3, 3},
		{1, -1, 0},
	}

	for index, test := range tests {
		if Add(test.a, test.b) != test.want {
			t.Errorf("testNo.%d Add(%d, %d) is incorrect.", index, test.a, test.b)
		}
	}
}

func TestIsEven(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{-1, false},
		{0, true},
		{1, false},
		{2, true},
		{3, false},
		{4, true},
		{5, false},
	}

	for index, test := range tests {
		if IsEven(test.n) != test.want {
			t.Errorf("testNo.%d isEven(%d) is incorrect.", index, test.n)
		}
	}
}

func TestValidateAge(t *testing.T) {
	tests := []struct {
		age     int
		wantErr bool
	}{
		{-1, true},
		{0, false},
		{1, false},
		{2, false},
		{3, false},
		{4, false},
		{5, false},
	}

	for index, test := range tests {
		err := ValidateAge(test.age)

		if err != nil != test.wantErr {
			t.Errorf("testNo.%d == ValidateAge(%d) is incorrect.", index, test.age)
		}
	}
}
