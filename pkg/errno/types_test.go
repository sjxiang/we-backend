package errno

import (
	"testing"
)
func Test_with_message(t *testing.T) {
	err := ErrDuplicateEntry.WithMessage("邮箱冲突")
	t.Log(err)
}