package service

import "testing"

func TestCreateUserWithEmptyName(
	t *testing.T,
) {

	service := UserService{}

	err := service.Create("")

	if err == nil {

		t.Errorf(
			"expected error for empty name",
		)
	}
}
