// Code generated by "stringer -linecomment -type DLLStorageClass"; DO NOT EDIT.

package enum

import "strconv"

const _DLLStorageClass_name = "nonedllexportdllimport"

var _DLLStorageClass_index = [...]uint8{0, 4, 13, 22}

func (i DLLStorageClass) String() string {
	if i >= DLLStorageClass(len(_DLLStorageClass_index)-1) {
		return "DLLStorageClass(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DLLStorageClass_name[_DLLStorageClass_index[i]:_DLLStorageClass_index[i+1]]
}
