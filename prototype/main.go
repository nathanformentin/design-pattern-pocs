package main

const _fourSpacesIndentation = "    "

func main() {
	testFile1 := &file{name: "test_file1"}
	testFile2 := &file{name: "test_file2"}
	cloneTestFile3 := testFile2.clone()

	childTestFolder := &folder{
		internalInodes: []inode{testFile1},
		name:           "test_folder1",
	}

	fatherTestFolder := &folder{
		internalInodes: []inode{testFile1, testFile2, cloneTestFile3, childTestFolder},
		name:           "test_folder1",
	}

	fatherTestFolder.print(_fourSpacesIndentation)

	clonedTestFolder := fatherTestFolder.clone()

	clonedTestFolder.print(_fourSpacesIndentation)
}
