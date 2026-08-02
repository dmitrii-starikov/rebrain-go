package analys

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAnalys(t *testing.T) {
	req := require.New(t)

	cases := map[string]struct {
		file string
		want *AnalysResult
	}{
		"simple": {
			file: "../../assets/astanalys/simple.go",
			want: &AnalysResult{0, 12, 6, 2},
		},
		"empty": {
			file: "../../assets/astanalys/empty.go",
			want: &AnalysResult{0, 0, 0, 0},
		},
		"more info": {
			file: "../../assets/astanalys/astanalys.go",
			want: &AnalysResult{3, 3, 7, 4},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			res, err := Analys(cs.file)
			req.NoError(err)
			req.Equal(cs.want, res)
		})
	}
}
