// Code generated by "stringer -type=InstanceType instancetype.go"; DO NOT EDIT

package terraform

import "fmt"

const _InstanceType_name = "TypeInvalidTypePrimaryTypeTaintedTypeDeposed"

var _InstanceType_index = [...]uint8{0, 11, 22, 33, 44}

func (i InstanceType) String() string {
	if i < 0 || i >= InstanceType(len(_InstanceType_index)-1) {
		return fmt.Sprintf("InstanceType(%d)", i)
	}
	return _InstanceType_name[_InstanceType_index[i]:_InstanceType_index[i+1]]
}
