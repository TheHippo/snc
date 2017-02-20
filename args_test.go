package snc

import "testing"

func TestParseArgs(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   uint32
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				args: []string{"foo", "localhost", "999"},
			},
			want:    "localhost",
			want1:   999,
			wantErr: false,
		},
		{
			name: "invalid int",
			args: args{
				args: []string{"foo", "localhost", "foobar"},
			},
			want:    "",
			want1:   0,
			wantErr: true,
		},
		{
			name: "invalid number of arguments",
			args: args{
				args: []string{"foo", "localhost"},
			},
			want:    "",
			want1:   0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := ParseArgs(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseArgs() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ParseArgs() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
