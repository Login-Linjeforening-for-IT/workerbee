package images

import (
	"image"
	"os"
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

func TestCheckFileSize(t *testing.T) {
	t.Run("SizeTest", func(t *testing.T) {
		testImageFolder := "../testdata/images/"

		testCases := []struct {
			fileName  string
			fileType  string
			maxSize   int64
			expectErr bool
		}{
			{"badratio_4867kb.gif", "image/gif", 2000 * 1024, true},
			{"goodratio_20kb.jpg", "image/jpeg", 500 * 1024, false},
			{"goodratio_1390kb.jpg", "image/jpeg", 500 * 1024, true},
			{"goodratio_23kb.png", "image/png", 500 * 1024, false},
		}

		for _, tc := range testCases {
			t.Run(tc.fileName, func(t *testing.T) {
				filePath := testImageFolder + tc.fileName
				f, err := os.Open(filePath)
				if err != nil {
					t.Fatalf("Failed to open file %s: %s", filePath, err)
				}
				defer f.Close()

				info, err := f.Stat()
				if err != nil {
					t.Fatalf("Failed to get file info for %s: %s", filePath, err)
				}

				err = checkFileSize(info.Size(), tc.fileType)

				if tc.expectErr && err == nil {
					t.Errorf("Expected error for file %s, but got none", filePath)
				}

				if !tc.expectErr && err != nil {
					t.Errorf("Unexpected error for file %s: %s", filePath, err)
				}
			})
		}
	})
}

func TestCheckFileRatio(t *testing.T) {
	t.Run("RatioTests", func(t *testing.T) {
		testImageFolder := "../testdata/images/"

		testCases := []struct {
			fileName  string
			ratioW    int
			ratioH    int
			expectErr bool
		}{
			{"badratio_4867kb.gif", 10, 4, true},
			{"goodratio_20kb.jpg", 10, 4, false},
			{"goodratio_1390kb.jpg", 10, 4, false},
			{"goodratio_23kb.png", 10, 4, false},
			{"goodratio_23kb.png", 233, 13, true}, // Random ratio
			{"goodratio32_51kb.png", 3, 2, false},
			{"goodratio32_51kb.png", 123, 33, true}, // Random ratio
		}

		for _, tc := range testCases {
			t.Run(tc.fileName, func(t *testing.T) {
				filePath := testImageFolder + tc.fileName
				f, err := os.Open(filePath)
				if err != nil {
					t.Fatalf("Failed to open file %s: %s", filePath, err)
				}
				defer f.Close()

				img, _, err := image.Decode(f)
				if err != nil {
					t.Fatalf("Failed to decode image %s: %s", filePath, err)
				}

				err = checkFileRatio(img, tc.ratioW, tc.ratioH)

				if tc.expectErr && err == nil {
					t.Errorf("Expected error for file %s, but got none", filePath)
				}

				if !tc.expectErr && err != nil {
					t.Errorf("Unexpected error for file %s: %s", filePath, err)
				}
			})
		}
	})
}
