package config_test

import (
	"fmt"
	"testing"

	"github.com/mahiro72/todo-task-app/config"
)

func TestNew(t *testing.T) {
	wantPort := 3333
	t.Setenv("PORT",fmt.Sprint(wantPort))
	
	got,err := config.New()
	if err != nil {
		t.Fatalf("cannot create config: %v",err)
	}
	if got.Port != wantPort {
		t.Errorf("want %d, but %d",wantPort,got.Port)
	}
	wantEnv:="dev"
	if got.Env != wantEnv {
		t.Errorf("want %s, but %s",wantEnv,got.Env)
	}
}
