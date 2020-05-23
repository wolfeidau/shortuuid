package shortuuid

import (
	"reflect"
	"testing"
)

func TestMustShorten(t *testing.T) {
	type args struct {
		uuidv4 []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should shorten without error",
			args: args{uuidv4: []string{"1c807ffe-b179-4a15-aa92-e06f5b5cf268"}},
			want: "4X8k1aRv5UMp9pdKzxaHwu",
		},
		{
			name: "should shorten more than one UUIDv4 without error",
			args: args{uuidv4: []string{"1c807ffe-b179-4a15-aa92-e06f5b5cf268", "30e300b4-9a2c-4ba7-91b4-7ca1327ffd0a"}},
			want: "2vG4a8G4TqeSFe6hmzZKGs7RRgr6fiCgVywf65Gt7iZ7",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MustShorten(tt.args.uuidv4...); got != tt.want {
				t.Errorf("MustShorten() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustUnShorten(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "should un-shorten without error",
			args: args{val: "4X8k1aRv5UMp9pdKzxaHwu"},
			want: []string{"1c807ffe-b179-4a15-aa92-e06f5b5cf268"},
		},
		{
			name: "should un-shorten more than one UUIDv4 without error",
			args: args{val: "2vG4a8G4TqeSFe6hmzZKGs7RRgr6fiCgVywf65Gt7iZ7"},
			want: []string{"1c807ffe-b179-4a15-aa92-e06f5b5cf268", "30e300b4-9a2c-4ba7-91b4-7ca1327ffd0a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MustUnShorten(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustUnShorten() = %v, want %v", got, tt.want)
			}
		})
	}
}
