// Code generated by "stringer -type=MType"; DO NOT EDIT.

package lorawan

import "strconv"

const _MType_name = "JoinRequestJoinAcceptUnconfirmedDataUpUnconfirmedDataDownConfirmedDataUpConfirmedDataDownRejoinRequestProprietary"

var _MType_index = [...]uint8{0, 11, 21, 38, 57, 72, 89, 102, 113}

func (i MType) String() string {
	if i >= MType(len(_MType_index)-1) {
		return "MType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MType_name[_MType_index[i]:_MType_index[i+1]]
}