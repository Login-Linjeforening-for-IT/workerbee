package api

import (
	"testing"
)

func TestRemovePrefix(t *testing.T) {
	t.Run("Positive", func(t *testing.T) {
		if removePrefix("img/testfolder/testfile.txt", "img/testfolder/") != "testfile.txt" {
			t.Error("Failed to remove prefix")
		}

		// In the following case there is nothing to remove. Would be developer error
		if removePrefix("img/testfolder/testfile.txt", "img/wrongfolder/") != "img/testfolder/testfile.txt" {
			t.Error("Failed, should not have removed anything")
		}

	})
}

func TestByteConverter(t *testing.T) {
	t.Run("Positive", func(t *testing.T) {
		if byteConverter(1024, 2) != "1.00 KiB" {
			t.Error("Not converting bytes correctly")
		}

		if byteConverter((1024+512)*1024, 3) != "1.500 MiB" {
			t.Error("Not converting bytes correctly")
		}
	})
}
