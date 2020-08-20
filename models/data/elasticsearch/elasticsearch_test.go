package elasticsearch

import "testing"

func TestAddLog(t *testing.T) {
	MESSAGE := "泰戈尔🤩"
	if err := AddLog(MESSAGE); err != nil {
		t.Error(err)
	}else{
		t.Log(MESSAGE)
	}
}


func TestDeleteLog(t *testing.T) {
	logid := 1597917340
	err := DeleteLog(logid)
	t.Log(err)
}

func TestGetLog(t *testing.T) {
	logid := 1597917793
	t.Log(GetLog(logid))
}

func TestSearch(t *testing.T) {
	t.Log(Search("board", ""))
}