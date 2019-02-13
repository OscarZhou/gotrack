package track

import (
	"os"
	"strings"
	"testing"
)

func TestTrackExportedFile(t *testing.T) {
	expected := `Start function:	gotrack.TestTrackExportedFile 
End function:	gotrack.TestTrackExportedFile took 2.0114ms `

	path := "./output/test.txt"
	if _, err := os.Stat(path); os.IsExist(err) {
		if err := os.Remove(path); err != nil {
			t.Error(err)
		}
	}
	track := New(Config{
		Debug:   true,
		AsynLog: false,
		// Note ./ stands for the directory where the current file places
		ExportedPath: path,
	})

	if track.Error != nil {
		t.Error("error=", track.Error)
	}

	track.Start()
	track.End()

	f, err := os.Open(path)
	if err != nil {
		t.Error(err)
	}

	var actual = make([]byte, 1024)
	_, err = f.Read(actual)
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(string(actual), "Start function:	gotrack.TestTrackExportedFile") {
		t.Errorf("actual:%s\nexpected:%s\n", actual, expected)
	}

	if !strings.Contains(string(actual), "End function:	gotrack.TestTrackExportedFile") {
		t.Errorf("actual:%s\nexpected:%s\n", actual, expected)
	}
}
