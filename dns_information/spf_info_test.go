package dnsinformation

import (
	"testing"
)

func TestIsValidSPF(t *testing.T) {
	t.Run("Valid SPF Record", func(t *testing.T) {
		record := "v=spf1 include:_spf.example.com ~all"
		isValid := IsValidSPF(record)

		if !isValid {
			t.Error("Expected a valid SPF record, but got an invalid result")
		}
	})

	t.Run("Valid SPF Record (with additional content)", func(t *testing.T) {
		record := "v=spf1 include:_spf.example.com ~all - extra content"
		isValid := IsValidSPF(record)

		if !isValid {
			t.Error("Expected a valid SPF record, but got an invalid result")
		}
	})

	t.Run("Invalid SPF Record (does not start with v=spf1)", func(t *testing.T) {
		record := "include:_spf.example.com ~all"
		isValid := IsValidSPF(record)

		if isValid {
			t.Error("Expected an invalid SPF record, but got a valid result")
		}
	})

	t.Run("Invalid SPF Record (malformed v=spf1)", func(t *testing.T) {
		record := "v=spf234 include:_spf.example.com ~all"
		isValid := IsValidSPF(record)

		if isValid {
			t.Error("Expected an invalid SPF record, but got a valid result")
		}
	})
}

func TestHasSPFRecord(t *testing.T) {
	t.Run("Valid SPF Record (Starts with 'v=spf1')", func(t *testing.T) {
		record := "v=spf1 include:_spf.example.com ~all"
		hasSPF := HasSPFRecord(record)

		if !hasSPF {
			t.Error("Expected a valid SPF record, but got a false result")
		}
	})

	t.Run("Valid SPF Record (Contains 'spf')", func(t *testing.T) {
		record := "This is an spf record"
		hasSPF := HasSPFRecord(record)

		if !hasSPF {
			t.Error("Expected a valid SPF record, but got a false result")
		}
	})

	t.Run("Valid SPF Record (Contains '-all')", func(t *testing.T) {
		record := "v=spf1 -all"
		hasSPF := HasSPFRecord(record)

		if !hasSPF {
			t.Error("Expected a valid SPF record, but got a false result")
		}
	})

	t.Run("Valid SPF Record (Contains '~all')", func(t *testing.T) {
		record := "v=spf1 ~all"
		hasSPF := HasSPFRecord(record)

		if !hasSPF {
			t.Error("Expected a valid SPF record, but got a false result")
		}
	})

	t.Run("Non-SPF Record", func(t *testing.T) {
		record := "This is a regular text record"
		hasSPF := HasSPFRecord(record)

		if hasSPF {
			t.Error("Expected a non-SPF record, but got a true result")
		}
	})

	// Add more test cases as needed.
}
