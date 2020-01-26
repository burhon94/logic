package main

import "testing"

func Test_test_all_true_false(t *testing.T) {
	listAll := []struct {
		description string
		list        []int
		predicate   func(int) bool
		want        bool
	}{
		{"list is true", []int{1, 1, 1}, func(i int) bool {
			if i == 1 {
				return true
			}
			return false
		}, true},
		{"list is false", []int{1, 2, 3}, func(i int) bool {
			if i == 1 {
				return true
			}
			return false
		}, false},
	}
	for _, tt := range listAll {
		if got := all(tt.list, tt.predicate); got != tt.want {
			t.Errorf("Description: %v, all() = %v, want %v", tt.description, got, tt.want)
		}
	}
}

func Test_test_any_true_false(t *testing.T) {
	listAny := []struct {
		description string
		list        []int
		predicate   func(int) bool
		want        bool
	}{
		{"list is true", []int{1, 2, 3, 4, 5, 6, 7}, func(i int) bool {
			if i > 6 {
				return true
			}
			return false
		}, true},
		{"list is false", []int{1, 2, 3, 4, 5, 6}, func(i int) bool {
			if i > 7 {
				return true
			}
			return false
		}, false},
	}
	for _, tt := range listAny {
		if got := any(tt.list, tt.predicate); got != tt.want {
			t.Errorf("Descriptin: %v, any() = %v, want: %v", tt.description, got, tt.want)
		}
	}
}

func Test_test_none_true(t *testing.T) {
	listNone := []struct {
		description string
		list        []int
		predicate   func(int) bool
		want        bool
	}{
		{"list is true", []int{1, 2, 3, 4, 5}, func(i int) bool {
			if i >= 6 {
				return true
			}
			return false
		}, true},
	}
	for _, tt := range listNone {
		if got := none(tt.list, tt.predicate); got != tt.want {
			t.Errorf("Description: %v, none() = %v, want %v", tt.description, got, tt.want)
		}
	}
}

func Test_test_index_element(t *testing.T) {
	list := []struct {
		description string
		list        []int
		predicate   func(int) bool
		want        int
	}{
		{"index element is exist", []int{1, 2, 3}, func(i int) bool {
			if i == 3 {
				return true
			}
			return false
		}, 2},
		{"index element is not exist", []int{1, 2, 3}, func(i int) bool {
			if i == 4 {
				return true
			}
			return false
		}, -1},
	}
	for _, tt := range list {
		if got := indexElement(tt.list, tt.predicate); got != tt.want {
			t.Errorf("Description: %v, indexElement() = %v, want %v", tt.description, got, tt.want)
		}
	}
}

func Test_test_find_element(t *testing.T) {
	list := []struct {
		description string
		list        []int
		predicate   func(int) bool
		want        int
	}{
		{"find result", []int{1, 2, 3, 4, 5}, func(i int) bool {
			if i == 3 {
				return true
			}
			return false
		}, 3},
	}
	for _, tt := range list {
		if got := findElement(tt.list, tt.predicate); got != tt.want {
			t.Errorf("Description: %v, findElement() = %v, want %v", tt.description, got, tt.want)
		}
	}
}
