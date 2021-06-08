package domain

import (
	"testing"
)

func getValidColumns() []string {
	return []string{"E-mail", "Name", "Id", "Wage"}
}

func getInvalidColumns() []string {
	return []string{"Number", "Size"}
}

func getFirstAndLastName() []string {
	return []string{"f.", "l."}
}
func TestValidEmail(t *testing.T) {
	emailPosition := GetEmailPosition(getValidColumns())
	if emailPosition != 0 {
		t.Errorf("Fail: The Email Position expected was 0, got %d", emailPosition)
	} else {
		t.Logf("Success: The Email Position expected was 0, got %d", emailPosition)
	}

}

func TestInvalidEmail(t *testing.T) {

	emailPosition := GetEmailPosition(getInvalidColumns())
	if emailPosition != -1 {
		t.Errorf("Fail: The Email Position expected was -1, got %d", emailPosition)
	} else {
		t.Logf("Success: The Email Position expected was -1, got %d", emailPosition)
	}

}
func TestValidId(t *testing.T) {

	idPosition := GetIdPosition(getValidColumns())
	if idPosition != 2 {
		t.Errorf("Fail: The Id Position expected was 2, got %d", idPosition)
	} else {
		t.Logf("Success: The Id Position expected was 2, got %d", idPosition)
	}

}
func TestInvalidId(t *testing.T) {

	idPosition := GetIdPosition(getInvalidColumns())
	if idPosition != -1 {
		t.Errorf("Fail: The Id Position expected was -1, got %d", idPosition)
	} else {
		t.Logf("Success: The Id Position expected was -1, got %d", idPosition)
	}

}
func TestValidSalary(t *testing.T) {
	salaryPosition := GetSalaryPosition(getValidColumns())
	if salaryPosition != 3 {
		t.Errorf("Fail: The Salary Position expected was 3, got %d", salaryPosition)
	} else {
		t.Logf("Success: The Salary Position expected was 3, got %d", salaryPosition)
	}

}
func TestInvalidSalary(t *testing.T) {
	salaryPosition := GetSalaryPosition(getInvalidColumns())
	if salaryPosition != -1 {
		t.Errorf("Fail: The Salary Position expected was -1, got %d", salaryPosition)
	} else {
		t.Logf("Success: The Salary Position expected was -1, got %d", salaryPosition)
	}

}

func TestFullName(t *testing.T) {
	fullNamePosition, _, _ := GetNamePosition(getValidColumns())
	if fullNamePosition != 1 {
		t.Errorf("Fail: The Full Name Position expected was 1, got %d", fullNamePosition)
	} else {
		t.Logf("Success: The Full Name Position expected was 1, got %d", fullNamePosition)
	}

}

func TestFirstAndLastName(t *testing.T) {
	_, first, last := GetNamePosition(getFirstAndLastName())
	if first == 0 && last == 1 {
		t.Logf("Success: The First and Last Name Position expected was 0 and 1, got %d and %d", first, last)
	} else {
		t.Errorf("Fail: The First and Last Name Position expected was 0 and 1, got %d and %d", first, last)
	}
}

func TestInvalidName(t *testing.T) {
	full, first, last := GetNamePosition(getInvalidColumns())
	if full == 0 && first == 0 && last == 0 {
		t.Logf("Success: The Full, First and Last Name Position expected was 0, 0 and 0, got %d, %d and %d", full, first, last)
	} else {
		t.Errorf("Fail: The Full, First and Last Name Position expected was 0, 0 and 0, got %d, %d and %d", full, first, last)
	}
}
