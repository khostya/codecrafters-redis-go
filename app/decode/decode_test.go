package decode

import (
	"slices"
	"testing"
)

func TestArray(t *testing.T) {
	t.Parallel()

	test := []struct {
		name string
		data []byte
		exp  []string
	}{
		{
			name: "list \"PING\n",
			data: []byte("*1\r\n$4\r\nPING\r\n"),
			exp:  []string{"PING"},
		},
		{
			name: "list \"ECHO orange\"",
			data: []byte("*2\r\n$4\r\nECHO\r\n$6\r\norange\r\n"),
			exp:  []string{"ECHO", "orange"},
		},
		{
			name: "list \"GET strawberry\"",
			data: []byte("*2\r\n$3\r\nGET\r\n$10\r\nstrawberry\r\n"),
			exp:  []string{"GET", "strawberry"},
		},
		{
			name: "list \"SET strawberry grape\"",
			data: []byte("*3\r\n$3\r\nSET\r\n$10\r\nstrawberry\r\n$5\r\ngrape\r\n"),
			exp:  []string{"SET", "strawberry", "grape"},
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, _, err := Decode(tt.data)

			if err != nil {
				t.Error(err)
			}
			if !slices.Equal(tt.exp, got) {
				t.Fail()
			}
		})
	}
}
