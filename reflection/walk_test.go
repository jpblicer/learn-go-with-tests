package reflection

import (
	"testing"
	"reflect"
)

type Profile struct {
	Age int
	City string
}

type Person struct {
	Name string
	Profile Profile
}


func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"nested fields",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"slices",
			[]Profile {
					{33, "London"},
					{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
		{
			"arrays",
			[2]Profile {
					{33, "London"},
					{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
		{
			"maps",
			map[string]string{
					"Cow": "Moo",
					"Sheep": "Baa",
			},
			[]string{"Moo", "Baa"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
