// Code generated by "stringer -type=State -type=NextState"; DO NOT EDIT

package protocol

import "fmt"

const _NextState_name = "NextStateStatusNextStateLogin"

var _NextState_index = [...]uint8{0, 15, 29}

func (i NextState) String() string {
	i -= 1
	if i < 0 || i >= NextState(len(_NextState_index)-1) {
		return fmt.Sprintf("NextState(%d)", i+1)
	}
	return _NextState_name[_NextState_index[i]:_NextState_index[i+1]]
}
