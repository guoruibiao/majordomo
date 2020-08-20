package clipboard

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestDarwin_Get(t *testing.T) {
	if content, err := NewDarwin().Get(); err != nil {
		t.Error(err)
	}else{
		t.Log(content)
		assert.Equal(t, content, "TestDarwin_Get")
	}
}


func TestDarwin_Set(t *testing.T) {
	darwin := NewDarwin()
	content := "It works."
	darwin.Set(content)
	if content, err := NewDarwin().Get(); err != nil {
		t.Error(err)
	}else{
		t.Log(content)
		assert.Equal(t, content, "It works.")
	}

}