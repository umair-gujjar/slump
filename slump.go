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

	_ error   = (*message)(nil)
	_ Handler = (*message)(nil)
)

// Handler is an interface to handle message.
type Handler interface {
	Add(Values)
	Clear()
	Del(key string)
	Get(key string) interface{}
	HasValues() bool
	Keys() []string
	Len() int
	Set(key string, value interface{})
	SetText(string)
	Text() string
	Render() (string, error)
	String() string
	Error() string
}

// Values is the type of the map defining the mapping from keys to values.
type Values map[string]interface{}

// Message is the representation of a formated text.
type message struct {
	text   string
	values Values
}

// New returns a new instance of message.
//    s := slump.New("Hello, {.name}")
//    s.Set("name", "Gopher")
//
//    println(s)
func New(text string) Handler {
	return &message{
		text:   text,
		values: make(Values),
	}
}

// Str returns a formated text into string.
//    s := slump.Str("Hello, {.name}", slump.Values{"name": "Gopher"})
//
//    println(s)
func Str(text string, v Values) string {
	m := New(text)
	m.Add(v)
	return m.String()
}

// Err returns a formated text into error.
//    path := "filename.txt"
//
//    err := slump.Err("no such file or directory: {.path}", slump.Values{"path": path})
//
//    println(err.Error())
func Err(text string, v Values) error {
	return errors.New(Str(text, v))
}

// Add adds values.
func (m *message) Add(values Values) {
	for k, v := range values {
		m.values[k] = v
	}
}

// Clear clears all values.
func (m *message) Clear() {
	m.values = make(Values)
}

// Del deletes a value.
func (m *message) Del(key string) {
	delete(m.values, key)
}

// Error returns the formated text into string.
func (m *message) Error() string {
	return m.String()
}

// Get returns a value by name.
func (m *message) Get(key string) interface{} {
	if v, ok := m.values[key]; ok {
		return v
	}
	return nil
}

// HasValues returns if the text has values.
func (m *message) HasValues() bool {
	return m.Len() > 0
}

// Keys returns the values keys.
func (m *message) Keys() []string {
	var keys = make([]string, 0, m.Len())
	for k := range m.values {
		keys = append(keys, k)
	}
	return keys
}

// Len returns the number of values.
func (m *message) Len() int {
	return len(m.values)
}

// SetValue sets a value by name.
func (m *message) Set(key string, value interface{}) {
	m.values[key] = value
}

// SetText sets the text of message.
func (m *message) SetText(t string) {
	m.text = t
}

// Text returns the text of this current message.
func (m *message) Text() string {
	return m.text
}

// Render applies a parsed text to string.
func (m *message) Render() (s string, err error) {
	if m.text == "" || !m.HasValues() {
		return "", errEmptyText
	}

	t := template.New("")
	t.Delims(DelimsLeft, DelimsRight)

	t, err = t.Parse(m.text)
	if err != nil {
		return
	}

	var b bytes.Buffer
	err = t.Execute(&b, m.values)
	s = b.String()
	return
}

// String returns the formated text into string.
func (m *message) String() string {
	s, err := m.Render()
	if err != nil {
		return err.Error()
	}
	return s
}
