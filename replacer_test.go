package replacer

import (
	"testing"
)

func TestReplacer(t *testing.T) {
	config := "./testdata/conf/replacer.yaml"
	root := "./testdata/testfile"

	testCases := []struct {
		name   string
		tag    string
		revert bool
	}{
		{"replace test", "test", false},
		{"revert test", "test", true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r, err := NewReplacer(config)
			if err != nil {
				t.Errorf("parse config error %s => %v", config, err)
				return
			}
			if tc.revert {
				err = r.Revert(root, tc.tag)
				if err != nil {
					t.Errorf("revert %s error => %v", tc.tag, err)
					return
				}
			} else {
				err = r.Replace(root, tc.tag)
				if err != nil {
					t.Errorf("replace %s error => %v", tc.tag, err)
					return
				}
			}
		})
	}
}
