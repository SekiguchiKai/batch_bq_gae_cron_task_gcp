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

func TestUpdateUser(t *testing.T) {
	t.Run("UpdateUser", func(t *testing.T) {
		t.Run("Userの値を元に、既存のUserのIDとUserName以外の値が更新されること", func(t *testing.T) {
			actual := User{
				UserName:    _TestUserName,
				MailAddress: _TestMailAddress,
				Age:         _TestAge,
				Gender:      _TestGender,
				From:        _TestFrom,
			}

			actual = NewUser(actual)

			expected1 := User{
				UserName:    _TestUserName,
				MailAddress: _TestMailAddress + _Modification,
				Age:         _TestAge + 50,
				Gender:      _TestGender + _Modification,
				From:        _TestFrom + _Modification,
			}

			actual = UpdateUser(actual, expected1)

			expected2 := NewUser(expected1)

			if !reflect.DeepEqual(actual, expected2) {
				t.Errorf("actual : %+v, expected2 : %+v", actual, expected2)
			}

		})
	})

}