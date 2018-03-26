package model

import (
	"reflect"
	"strconv"
	"testing"
	"time"
)

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

func TestTranslateStructToSlice(t *testing.T) {
	t.Run("TranslateStructToSlice", func(t *testing.T) {
		t.Run("構造体Userを元に、そのフィールドの値を格納したstringのsliceが作成されること", func(t *testing.T) {
			u := User{
				UserName:    _TestUserName,
				MailAddress: _TestMailAddress,
				Age:         _TestAge,
				Gender:      _TestGender,
				From:        _TestFrom,
			}

			u = NewUser(u)
			u.CreatedAt = time.Now().UTC()
			u.UpdatedAt = time.Now().UTC()

			s := TranslateStructToSlice(u)

			if s[IDIndex] != u.ID {
				t.Errorf("actual : %+v, expected : %+v", s[IDIndex], u.ID)
			}
			if s[UserNameIndex] != u.ID {
				t.Errorf("actual : %+v, expected : %+v", s[UserNameIndex], u.UserName)
			}
			if s[MailAddressIndex] != u.MailAddress {
				t.Errorf("actual : %+v, expected : %+v", s[MailAddressIndex], u.MailAddress)
			}
			if s[AgeIndex] != strconv.Itoa(u.Age) {
				t.Errorf("actual : %+v, expected : %+v", s[AgeIndex], strconv.Itoa(u.Age))
			}
			if s[GenderIndex] != string(u.Gender) {
				t.Errorf("actual : %+v, expected : %+v", s[GenderIndex], string(u.Gender))
			}
			if s[FromIndex] != u.From {
				t.Errorf("actual : %+v, expected : %+v", s[FromIndex], u.From)
			}
			if s[CreatedAtIndex] != u.CreatedAt.String() {
				t.Errorf("actual : %+v, expected : %+v", s[CreatedAtIndex], u.CreatedAt.String())
			}

			if s[UpdatedAtIndex] != u.UpdatedAt.String() {
				t.Errorf("actual : %+v, expected : %+v", s[UpdatedAtIndex], u.UpdatedAt.String())
			}

		})
	})

}
