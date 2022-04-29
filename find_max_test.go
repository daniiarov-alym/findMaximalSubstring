package main

import "testing"
	

//TestEmpty calls findMaximalSubstring with empty string and checks for error
func TestEmpty(t *testing.T) {
	result := findMaximalSubstring("", 2)
	if (result != 0) {
		t.Errorf("FindMaximalSubstring(\"\", 2) = %d; want 0", result)
	}
}


//TestTwo calls findMaximalSubstring with string of length one and checks for error
func TestOne(t *testing.T) {
	result := findMaximalSubstring("h", 2)
	if (result != 1) {
		t.Errorf("FindMaximalSubstring(\"h\", 2) = %d; want 1", result)
	}
}

//TestTwo calls findMaximalSubstring with strings of length two and checks for error
func TestTwo(t *testing.T) {
	result := findMaximalSubstring("ab", 2)
	if (result != 2) {
		t.Fatalf("FindMaximalSubstring(\"ab\", 2) = %d; want 2", result)
	}
	
	result = findMaximalSubstring("aa", 2)
	if (result != 2) {
		t.Errorf("FindMaximalSubstring(\"aa\", 2) = %d; want 2", result)
	}
}

//TestThree calls findMaximalSubstring with strings of length three and checks for error
func TestThree(t *testing.T) {
	result := findMaximalSubstring("abc", 2)
	if (result != 2) {
		t.Fatalf("FindMaximalSubstring(\"abc\", 2) = %d; want 2", result)
	}
	
	result = findMaximalSubstring("aab", 2)
	if (result != 3) {
		t.Errorf("FindMaximalSubstring(\"aaa\", 2) = %d; want 3", result)
	}
}

//TestThree calls findMaximalSubstring with large string and checks for error
func TestString(t *testing.T) {
	result := findMaximalSubstring("abaccab", 2)
	if (result != 4) {
		t.Fatalf("FindMaximalSubstring(\"abaccab\", 2) = %d; want 4", result)
	}
}