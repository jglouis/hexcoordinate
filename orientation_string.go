// Code generated by "stringer -type=Orientation"; DO NOT EDIT

package hex

import "fmt"

const _Orientation_name = "NNESESSWNW"

var _Orientation_index = [...]uint8{0, 1, 3, 5, 6, 8, 10}

func (i Orientation) String() string {
	if i < 0 || i >= Orientation(len(_Orientation_index)-1) {
		return fmt.Sprintf("Orientation(%d)", i)
	}
	return _Orientation_name[_Orientation_index[i]:_Orientation_index[i+1]]
}
