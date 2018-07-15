package stringutil_test

import (
    "testing"
    "github.com/waynefoux/stringutil/palindrome"
)

func isPalindromeTest(t *testing.T) {
    if isPalindrome("foo") != true {
        t.Error("Expected true")
    }
}
