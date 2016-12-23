// Copyright 2016 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package slump is a simple string template implementation for Go.
package slump

import (
	"bytes"
	"errors"
	"html/template"
)

// See https://golang.org/pkg/text/template/#Template.Delims
var (
	DelimsLeft  = "{"
	DelimsRight = "}"
)

var (
	errEmptyText = errors.New("text to format was not provided")

	_ error = (*Message)(nil)
)

// A Value represents the key-value pairs in a Value.
type Value map[string]interface{}

// Add add the key-value pairs in values.
func (v Value) Add(values map[string]interface{}) {
	for key, val := range values {
		v.Set(key, val)
	}
}

// Clear clears the key-value pairs in a Value.
func (v Value) Clear() {
	for k := range v {
		delete(v, k)
	}
}

// Del deletes the values associated with key.
func (v Value) Del(key string) {
	delete(v, key)
}

// Get gets the first value associated with the given key.
// If there are no values associated with the key, Get returns nil.
func (v Value) Get(key string) interface{} {
	if value, ok := v[key]; ok {
		return value
	}
	return nil
}

// IsEmpty returns if the value has entries.
func (v Value) IsEmpty() bool {
	return len(v) == 0
}

// Keys returns the keys set in the Set values.
func (v Value) Keys() []string {
	var keys = make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	return keys
}

// Count returns the total number of value.
func (v Value) Count() int {
	return len(v)
}

// Set sets the value entries associated with key to the single element value.
// It replaces any existing values associated with key.
func (v Value) Set(key string, value interface{}) {
	v[key] = value
}

// Message is the representation of a formated text.
type Message struct {
	Text  string
	Value Value
}

// New returns a new instance of message.
//    s := slump.New("Hello, {.name}")
//    s.Value.Set("name", "Gopher")
//
//    println(s)
func New(text string) Message {
	return Message{
		Text:  text,
		Value: make(Value),
	}
}

// Str returns a formated text into string.
//    s := slump.Str("Hello, {.name}", slump.Value{"name": "Gopher"})
//
//    println(s)
func Str(text string, v Value) string {
	m := New(text)
	m.Value = v
	return m.String()
}

// Err returns a formated text into error.
//    path := "filename.txt"
//
//    err := slump.Err("no such file or directory: {.path}", slump.Value{"path": path})
//
//    println(err.Error())
func Err(text string, v Value) error {
	return errors.New(Str(text, v))
}

// Error returns the formated text into string.
func (m *Message) Error() string {
	return m.String()
}

// Render applies a parsed text to string.
func (m *Message) Render() (s string, err error) {
	if m.Text == "" {
		return "", errEmptyText
	}

	if m.Value.IsEmpty() {
		s = m.Text
		return
	}

	t := template.New("")
	t.Delims(DelimsLeft, DelimsRight)

	t, err = t.Parse(m.Text)
	if err != nil {
		return
	}

	var b bytes.Buffer
	err = t.Execute(&b, m.Value)
	s = b.String()
	return
}

// String returns the formated text into string.
func (m *Message) String() string {
	s, err := m.Render()
	if err != nil {
		return err.Error()
	}
	return s
}
