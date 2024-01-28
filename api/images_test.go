package api

import (
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
		testImageFolder := "../assets/img/"

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

				err = checkFileSize(f, tc.fileType)

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
		testImageFolder := "../assets/img/"

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

				err = checkFileRatio(f, tc.ratioW, tc.ratioH)

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

func TestCheckFileType(t *testing.T) {
	t.Run("FiletypeTest", func(t *testing.T) {
		testImageFolder := "../assets/img/"

		testCases := []struct {
			fileName  string
			fileType  string
			expectErr bool
		}{
			{"badratio_4867kb.gif", "image/gif", false},
			{"goodratio_20kb.jpg", "image/jpeg", false},
			{"goodratio_1390kb.jpg", "image/jpeg", false},
			{"goodratio_23kb.png", "image/png", false},
			{"i_like_bits.bmp", "image/bmp", true},
		}

		for _, tc := range testCases {
			t.Run(tc.fileName, func(t *testing.T) {
				filePath := testImageFolder + tc.fileName
				f, err := os.Open(filePath)
				if err != nil {
					t.Fatalf("Failed to open file %s: %s", filePath, err)
				}
				defer f.Close()

				fType, err := checkFileType(f)

				if tc.expectErr && err == nil {
					t.Errorf("Expected error for file %s, but got none", filePath)

					if fType != "" {
						t.Errorf("Expected empty string, but got %s", fType)
					}
				}

				if !tc.expectErr && err != nil {
					t.Errorf("Unexpected error for file %s: %s", filePath, err)

					if fType != tc.fileType {
						t.Errorf("Expected %s, but got %s", tc.fileType, fType)
					}
				}
			})
		}
	})
}
