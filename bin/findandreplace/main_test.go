package main

import (
	"os"
	"testing"
)
var replaceYaml = `"test.txt":
  doreplace: newcontent
"test2.txt":
  doreplace: newcontent
"**/deep.txt":
  doreplace: newcontent`
var replacetxt = "testcontent: doreplace"
var noreplacetxt = "testcontent: donotreplace"
var replacedContent = "testcontent: newcontent"

func TestMain(t *testing.T) {
	createTestData()
	run("replace.yaml", "test")
	t.Run("Should replace test.txt", func(t *testing.T) {
		// test.txt # should replace
		c, _ := os.ReadFile("test/test.txt")
		if string(c) != replacedContent {
      t.Errorf(`test.txt should have replaced doreplace with 'newcontent'
got:
%s`, c)
		}
	})
  t.Run("shoud not replace testno.txt", func(t *testing.T) {
    c, _ := os.ReadFile("test/testno.txt")
    if string(c) != replacetxt {
      t.Errorf(`testno.txt should not have been replaced.
got:
%s`, c)
    }
  })
  t.Run("shoud not replace test2.txt", func(t *testing.T) {
    c, _ := os.ReadFile("test/test2.txt")
    if string(c) != noreplacetxt {
      t.Errorf(`test2.txt should not have been replaced.
got:
%s`, c)
    }
  })
  t.Run("shoud replace deep.txt in root folder", func(t *testing.T) {
    c, _ := os.ReadFile("test/deep.txt")
    if string(c) != replacedContent {
      t.Errorf(`deep.txt should have been replaced.
got:
%s`, c)
    }
  })

  t.Run("shoud replace deep.txt in deep subfolders", func(t *testing.T) {
    c, _ := os.ReadFile("test/child1/child2/cild3/deep.txt")
    if string(c) != replacedContent {
      t.Errorf(`deep.txt should have been replaced.
got:
%s`, c)
    }
  })

  // Cleanup
  os.RemoveAll("test")
  os.Remove("replace.yaml")
}

func createTestData() {
	os.MkdirAll("test/child1/child2/cild3", os.ModePerm)
	os.WriteFile("replace.yaml", []byte(replaceYaml), 0644)
	os.WriteFile("test/test.txt", []byte(replacetxt), 0644)
	os.WriteFile("test/testno.txt", []byte(replacetxt), 0644)
	os.WriteFile("test/test2.txt", []byte(noreplacetxt), 0644)
	os.WriteFile("test/child1/child2/cild3/deep.txt", []byte(replacetxt), 0644)
	os.WriteFile("test/deep.txt", []byte(replacetxt), 0644)
}
