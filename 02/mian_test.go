package main

import (
	"testing"
)

// a. обычное поведение с K меньшим, чем количество уникальных слов;
// b. передача пустого списка слов;
// c. передача списка слов, где K больше, чем число уникальных слов.
func Test_a(t *testing.T) {
	result := PrintWordFrequencies([]string{"ab", "ac", "dr", "ab", "ac", "ac", "ab", "ab", "fd"}, 3)
	expected := "ab ac dr"
	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}
func Test_b(t *testing.T) {
	result := PrintWordFrequencies([]string{}, 3)
	expected := ""
	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}
func Test_c(t *testing.T) {
	result := PrintWordFrequencies([]string{"ab", "ac", "dr", "ab", "ac", "ac", "ab", "ab", "fd"}, 20)
	expected := "ab ac dr fd"
	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}
func Test_d(t *testing.T) {
	result := PrintWordFrequencies([]string{"ab", "ac", "dr", "ab", "ac", "ac", "ab", "ab", "fd"}, 0)
	expected := ""
	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}
