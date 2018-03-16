package model

import "testing"

const (
	_TestUserName    = "TestUserName"
	_TestMailAddress = "TestMailAddress"
	_TestAge         = 50
	_TestGender      = "TestGender"
	_TestFrom        = "TestFrom"
	_Modification    = "Modification"
)

func TestNewUser(t *testing.T) {
	t.Run("NewUser", func(t *testing.T) {
		t.Run("2つの同じUserNameを元にすると、同じhash値のIDを持ったUserが返されること", func(t *testing.T) {
			u1 := User{
				UserName: _TestUserName,
			}

			u2 := User{
				UserName: _TestUserName,
			}

			if NewUser(u1).ID != NewUser(u2).ID {
				t.Errorf("u1.ID : %s, u2.ID : %s", u1.ID, u2.ID)
			}
		})

		t.Run("2つの異なるUserNameを元にすると、異なるhash値のIDを持ったUserが返されること", func(t *testing.T) {
			u1 := User{
				UserName: _TestUserName,
			}

			u2 := User{
				UserName: _TestUserName + _Modification,
			}

			if NewUser(u1).ID == NewUser(u2).ID {
				t.Errorf("u1.ID : %s, u2.ID : %s", u1.ID, u2.ID)
			}
		})
	})
}