package scanning_test

import (
	scanning "newmod/scaningFileToPath"
	"testing"
)

func TestListDirByWalk(t *testing.T) {

	_, _ = scanning.ListDirByWalk("/home/anton/projects/golang-2/duplicate/scaningFileToPath/emptyFolderbyTest")

}
