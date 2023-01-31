package b_summa_k_oplate

import "testing"

func Test_testing_sum(t *testing.T) {
	type args struct {
		a   int
		b   int
		res []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "base test",
			args: args{
				a:   6,
				b:   12,
				res: []int{2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3},
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testing_sum(tt.args.a, tt.args.b, tt.args.res); got != tt.want {
				t.Errorf("testing_sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
