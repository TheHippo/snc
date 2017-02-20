package utils

import (
	"testing"
)

func TestOptionalIntValue_String(t *testing.T) {
	type fields struct {
		Init   bool
		Parsed bool
		Value  uint32
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "no init",
			fields: fields{
				Init:   false,
				Parsed: false,
				Value:  0,
			},
			want: "8888",
		},
		{
			name: "init",
			fields: fields{
				Init:   true,
				Parsed: true,
				Value:  999,
			},
			want: "true 999",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lv := &OptionalIntValue{
				Init:   tt.fields.Init,
				Parsed: tt.fields.Parsed,
				Value:  tt.fields.Value,
			}
			if got := lv.String(); got != tt.want {
				t.Errorf("OptionalIntValue.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptionalIntValue_Set(t *testing.T) {
	type fields struct {
		Init   bool
		Parsed bool
		Value  uint32
	}
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "no value",
			fields: fields{},
			args: args{
				input: "",
			},
			wantErr: false,
		}, {
			name:   "invalid input",
			fields: fields{},
			args: args{
				input: "foobar",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lv := &OptionalIntValue{
				Init:   tt.fields.Init,
				Parsed: tt.fields.Parsed,
				Value:  tt.fields.Value,
			}
			if err := lv.Set(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("OptionalIntValue.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
