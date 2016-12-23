// Copyright 2016 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slump

import (
	"reflect"
	"sort"
	"testing"
)

type messageTest struct {
	text  string
	value Value
	keys  []string
	want  string
}

var getMessagesTests = []messageTest{
	{
		"hello, {.name}",
		Value{"name": "Gophers"},
		[]string{"name"},
		"hello, Gophers",
	},
	{
		"{.lang} {.version} is released",
		Value{"lang": "Go", "version": "1.7"},
		[]string{"lang", "version"},
		"Go 1.7 is released",
	},
	{
		"the type of {.value} is float64",
		Value{"value": 3.14159265359},
		[]string{"value"},
		"the type of 3.14159265359 is float64",
	},
	{
		"the type of {printf \"%.2f\" .value} is float64",
		Value{"value": 3.14},
		[]string{"value"},
		"the type of 3.14 is float64",
	},
}

func TestMessage(t *testing.T) {
	for _, tt := range getMessagesTests {
		m := New(tt.text)
		m.Value.Add(tt.value)

		if got := m.String(); got != tt.want {
			t.Errorf("New() got %q; want %q", got, tt.want)
		}
		if got := m.Error(); got != tt.want {
			t.Errorf("New() got %q; want %q", got, tt.want)
		}
	}
}

func TestClear(t *testing.T) {
	for _, tt := range getMessagesTests {
		m := New(tt.text)
		m.Value.Add(tt.value)
		m.Value.Clear()

		if got := m.Value.IsEmpty(); !got {
			t.Errorf("Clear() got %t; want false", got)
		}
	}
}

func TestDel(t *testing.T) {
	for _, tt := range getMessagesTests {
		m := New(tt.text)
		m.Value.Add(tt.value)

		for k := range tt.value {
			m.Value.Del(k)
		}
		if got, want := m.Value.Count(), 0; got != want {
			t.Errorf("Count() got %d; want %d", got, want)
		}
	}
}

func TestErr(t *testing.T) {
	want := "interface is string, not []string"
	err := Err("interface is {.type}, not []string", Value{"type": reflect.TypeOf("test")})

	if err == nil {
		t.Errorf("Err() got nil; want %q", want)
	}
	if got := err.Error(); got != want {
		t.Errorf("Err() got %q; want %q", got, want)
	}
}

func TestGetSet(t *testing.T) {
	want := "John Doe"
	m := New("")
	m.Value.Set("author", want)

	if got := m.Value.Get("author"); got != want {
		t.Errorf("Get() got %q; want %q", got, want)
	}
}

func TestGetHasNilValue(t *testing.T) {
	m := New("")
	if got := m.Value.Get("nothing"); got != nil {
		t.Errorf("Get() got %v; want nil", got)
	}
}

func TestIsEmptyValue(t *testing.T) {
	for _, tt := range getMessagesTests {
		m := New(tt.text)
		m.Value.Add(tt.value)

		if got, want := m.Value.IsEmpty(), true; got == want {
			t.Errorf("IsEmpty() got %t; want %t", got, want)
		}
	}
}

func TestKeys(t *testing.T) {
	for _, tt := range getMessagesTests {
		m := New(tt.text)
		m.Value.Add(tt.value)

		if got, want := m.Value.Count(), len(tt.keys); got != want {
			t.Errorf("Keys() got len %d; want %d", got, want)
		}

		keys := m.Value.Keys()
		sort.Strings(keys)
		sort.Strings(tt.keys)

		for i := range keys {
			if tt.keys[i] != keys[i] {
				t.Errorf("Keys() got %v; want %v", keys[i], tt.keys[i])
			}
		}
	}
}

func TestCountValues(t *testing.T) {
	for _, tt := range getMessagesTests {
		m := New(tt.text)
		m.Value.Add(tt.value)

		if got, want := m.Value.Count(), len(tt.value); got != want {
			t.Errorf("Count() got %d; want %d", got, want)
		}
	}
}

func TestSetText(t *testing.T) {
	for _, tt := range getMessagesTests {
		m := New("")
		m.Text = tt.text
		m.Value.Add(tt.value)

		if got := m.Text; got != tt.text {
			t.Errorf("Text() got %q; want %q", got, tt.text)
		}
	}
}

func TestRenderWithEmptyValue(t *testing.T) {
	for _, tt := range getMessagesTests {
		m := New(tt.text)
		s, err := m.Render()

		if err != nil {
			t.Error(err)
		}
		if s != tt.text {
			t.Errorf("Render() got %q; want %q", s, tt.text)
		}
	}
}

func TestRender(t *testing.T) {
	for _, tt := range getMessagesTests {
		m := New(tt.text)
		m.Value.Add(tt.value)

		s, err := m.Render()
		if err != nil {
			t.Error(err)
		}
		if s != tt.want {
			t.Errorf("Render() got %q; want %q", s, tt.text)
		}
	}
}

func TestEmptyTextFail(t *testing.T) {
	m := New("")
	if got, want := m.String(), errEmptyText.Error(); got != want {
		t.Errorf("String() got %q; want %q", got, want)
	}
}

func TestParseTextErrorFail(t *testing.T) {
	m := New("Bad {{variable}}")
	m.Value.Set("variable", "trip")

	if got, want := m.String(), "Bad trip"; got == want {
		t.Errorf("String() got %q; want %q", got, want)
	}
}

var tmplBenchmark = "{.prefix}}, {.name}"

func BenchmarkStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Str(tmplBenchmark, Value{"prefix": "Hello", "name": "Gopher"})
	}
}

func BenchmarkErr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Err(tmplBenchmark, Value{"prefix": "Hello", "name": "Gopher"})
	}
}
