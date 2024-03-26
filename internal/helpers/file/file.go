package file

import (
	"bufio"
	"io"
	"os"
)

func FileExists(filePath string) (bool, error) {
	if _, err := os.Stat(filePath); err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}

func ReplaceFile(sourcePath string, destinationPath string) error {
	err := os.Rename(sourcePath, destinationPath)
	if err != nil {
		return err
	}

	return nil
}

func CopyFile(src, dst string) error {
	// Open the source file
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Create the destination file
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// Copy the content from source to destination
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	// Sync to flush writes to disk
	err = dstFile.Sync()
	if err != nil {
		return err
	}

	return nil
}

func WriteStringFile(filePath string, text string) error {
	// Open the file for writing (truncating it if it exists)
	file, createErr := os.Create(filePath)
	if createErr != nil {
		return createErr
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	_, writeErr := writer.WriteString(text)
	if writeErr != nil {
		return writeErr
	}

	// Flush the bufio.Writer to ensure all data is written to the file
	flushErr := writer.Flush()
	if flushErr != nil {
		return flushErr
	}

	return nil
}
