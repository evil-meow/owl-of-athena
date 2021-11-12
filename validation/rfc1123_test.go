package validation

import "testing"

func TestTooLongStringIsNotValid(t *testing.T) {
	s := "ihavemorethan63characterssoiamnotvalidreallyiamnotvalidirepeatiamnotvalid"

	if IsRFC1123(s) {
		t.Fatalf("A string that is too long is valid")
	}
}

func TestContainsAnUppercaseLetterIsNotValid(t *testing.T) {
	s := "iamnotValid"

	if IsRFC1123(s) {
		t.Fatalf("A string that contains an uppercase letter is not valid")
	}
}

func TestContainsOnlyAlphanumericAndDashIsValid(t *testing.T) {
	s := "iam-2-valid"

	if !IsRFC1123(s) {
		t.Fatalf("A string that contains only alphanumeric and dash is valid")
	}
}
