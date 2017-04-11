package gimei

import "testing"

func TestNewNameByLastNameEmpty(t *testing.T) {
	if _, err := NewNameByLastName(""); err == nil {
		t.Error("NewNameByLastName should raise error when given empty last name.")
	}
}

func TestNewNameByLastName(t *testing.T) {
	for i := 0; i < 100; i++ {
		name, err := NewNameByLastName("田中")
		if err != nil {
			t.Error(err)
		}
		if name.Last.Kanji() != "田中" {
			t.Error("cannot create by last name")
		}
	}
}

func TestNewNameByLastNameHiragana(t *testing.T) {
	for i := 0; i < 100; i++ {
		name, err := NewNameByLastName("たなか")
		if err != nil {
			t.Error(err)
		}
		if name.Last.Kanji() != "田中" {
			t.Error("cannot create by last name")
		}
	}
}

func TestNewNameByLastNameKataakana(t *testing.T) {
	for i := 0; i < 100; i++ {
		name, err := NewNameByLastName("タナカ")
		if err != nil {
			t.Error(err)
		}
		if name.Last.Kanji() != "田中" {
			t.Error("cannot create by last name")
		}
	}
}

func TestNewMaleByLastNameEmpty(t *testing.T) {
	if _, err := NewMaleByLastName(""); err == nil {
		t.Error("NewNameByLastName should raise error when given empty last name.")
	}
}

func TestNewMaleByLastName(t *testing.T) {
	for i := 0; i < 100; i++ {
		name, err := NewMaleByLastName("田中")
		if err != nil {
			t.Error(err)
		}
		if name.IsFemale() {
			t.Error("not male")
		}
		if name.Last.Kanji() != "田中" {
			t.Error("cannot create by last name")
		}
	}
}

func TestNewMaleByLastNameHiragana(t *testing.T) {
	for i := 0; i < 100; i++ {
		name, err := NewMaleByLastName("たなか")
		if err != nil {
			t.Error(err)
		}
		if name.IsFemale() {
			t.Error("not male")
		}
		if name.Last.Kanji() != "田中" {
			t.Error("cannot create by last name")
		}
	}
}

func TestNewMaleByLastNameKataakana(t *testing.T) {
	for i := 0; i < 100; i++ {
		name, err := NewMaleByLastName("タナカ")
		if err != nil {
			t.Error(err)
		}
		if name.IsFemale() {
			t.Error("not male")
		}
		if name.Last.Kanji() != "田中" {
			t.Error("cannot create by last name")
		}
	}
}

func TestNewFemaleByLastNameEmpty(t *testing.T) {
	if _, err := NewFemaleByLastName(""); err == nil {
		t.Error("NewNameByLastName should raise error when given empty last name.")
	}
}

func TestNewFemaleByLastName(t *testing.T) {
	for i := 0; i < 100; i++ {
		name, err := NewFemaleByLastName("田中")
		if err != nil {
			t.Error(err)
		}
		if name.IsMale() {
			t.Error("not female")
		}
		if name.Last.Kanji() != "田中" {
			t.Error("cannot create by last name")
		}
	}
}

func TestNewFemaleByLastNameHiragana(t *testing.T) {
	for i := 0; i < 100; i++ {
		name, err := NewFemaleByLastName("たなか")
		if err != nil {
			t.Error(err)
		}
		if name.IsMale() {
			t.Error("not female")
		}
		if name.Last.Kanji() != "田中" {
			t.Error("cannot create by last name")
		}
	}
}

func TestNewFemaleByLastNameKataakana(t *testing.T) {
	for i := 0; i < 100; i++ {
		name, err := NewFemaleByLastName("タナカ")
		if err != nil {
			t.Error(err)
		}
		if name.IsMale() {
			t.Error("not female")
		}
		if name.Last.Kanji() != "田中" {
			t.Error("cannot create by last name")
		}
	}
}
