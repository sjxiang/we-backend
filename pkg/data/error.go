package data

import "we-backend/pkg/errno"


const RecordNoFound       uint32 = 200_001
const RecordAlreadyExists uint32 = 200_002
const DuplicateEntry      uint32 = 200_003


var (
	ErrRecordNoFound       = errno.NewErrNo(RecordNoFound, "资源不存在")
	ErrRecordAlreadyExists = errno.NewErrNo(RecordAlreadyExists, "已存在")
	ErrDuplicateEntry      = errno.NewErrNo(DuplicateEntry, "字段冲突")  	
)