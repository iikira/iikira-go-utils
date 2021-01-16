package utils_test

import (
	"fmt"
	"github.com/iikira/iikira-go-utils/utils"
	"testing"
)

func TestWalkDir(t *testing.T) {
	files, err := utils.WalkDir("/Users/syy/tmp", "")
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
}
