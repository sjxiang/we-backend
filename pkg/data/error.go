package data

import "we-backend/pkg/errno"


const RecordNoFoundError       uint32 = 200_001
const RecordAlreadyExistsError uint32 = 200_002
const DuplicateEntryError      uint32 = 200_003


var (
	ErrRecordNoFound       = errno.NewErrNo(RecordNoFoundError, "no matching record found")
	ErrRecordAlreadyExists = errno.NewErrNo(RecordAlreadyExistsError, "record already exists")
	ErrDuplicateEntry      = errno.NewErrNo(DuplicateEntryError, "duplicate entry")  	
)