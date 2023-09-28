//go:build !windows
// +build !windows

package system

import chshare "github.com/IOTech17/neo-rport/share"

func (i Interpreter) Get() string {
	if i.InterpreterAliases != nil && i.InterpreterNameFromInput != "" {
		if mappedInterpreter, ok := i.InterpreterAliases[i.InterpreterNameFromInput]; ok {
			return mappedInterpreter
		}
	}

	if i.InterpreterNameFromInput == "" {
		return i.GetDefault()
	}

	return i.InterpreterNameFromInput
}

func (i Interpreter) GetDefault() string {
	return chshare.UnixShell
}
