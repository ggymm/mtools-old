package config

import "testing"

func Test_configPath(t *testing.T) {
	path := configPath()
	t.Log(path)
}

func Test_currentPath(t *testing.T) {
	path := currentPath()
	t.Log(path)
}

func Test_tempPath(t *testing.T) {
	path := tempPath()
	t.Log(path)
}

func Test_currentExec(t *testing.T) {
	path := currentExec()
	t.Log(path)
}

func Test_currentCaller(t *testing.T) {
	path := currentCaller()
	t.Log(path)
}
