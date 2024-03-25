package main

import "testing"

func TestFirstLetterUp(t *testing.T) {
	type args struct {
		value string
	}
	var (
		tests = []struct {
			name string
			args args
			want string
		}{
			{
				name: "first letter is lower, last letter isnt dot",
				args: args{"hello world"},
				want: "Hello world.",
			},
			{
				name: "firts letter is upper, last letter isnt dot",
				args: args{"Hello world"},
				want: "Hello world.",
			},
			{
				name: "firts letter is lower, last letter is dot",
				args: args{"hello world."},
				want: "Hello world.",
			},
			{
				name: "firts letter is upper, last letter is dot",
				args: args{"Hello world."},
				want: "Hello world.",
			},
			{
				name: "firts letter is lower, last letter is question mark",
				args: args{"why so serious?"},
				want: "Why so serious?",
			},
			{
				name: "firts letter is upper, last letter is exclamation mark",
				args: args{"Have a nice day!"},
				want: "Have a nice day!",
			},
			{
				name: "empty string",
				args: args{""},
				want: "",
			},
		}
	)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstLetterUp(tt.args.value); got != tt.want {
				t.Errorf("FirstLetterUp() = %v, want %v", got, tt.want)
			}
		})
	}
}
