package fileops

import (
	"desktop_notes/core/encryption"
	"fmt"
	"io/fs"
	"log"
	"os"
)

const defaultKey = "alskdjfaljgugh83hghljdhbkjnadibuaiugrhaiuhgiauwhgalkghlekghalireubghlaiuerh"

func GetNotesLocation() (string, error) {
	homeDir, homeDirErr := os.UserHomeDir()
	fmt.Printf("Homedir %v", homeDir)
	if homeDirErr != nil {
		fmt.Printf("Failed to get user path: %v", homeDirErr)
		return "", homeDirErr
	}
	notesLocation := homeDir + "/.local/state/desktop_notes"
	mkDirErr := os.MkdirAll(notesLocation, 0700)
	if mkDirErr != nil {
		fmt.Printf("Failed to create notes location: %v", mkDirErr)
		return "", mkDirErr
	}
	return notesLocation, nil
}

func ListSavedFiles() ([]fs.DirEntry, error) {
	location, err := GetNotesLocation()
	if err != nil {
		fmt.Printf("Get location error %v", err)
		return nil, err
	}
	files, readErr := os.ReadDir(location)
	if readErr != nil {
		fmt.Printf("Read location error %v", readErr)
		return nil, readErr
	}
	for _, file := range files {
		fmt.Printf("File name %v", file)
	}
	return files, nil
}

func SaveFile(name string, contents string, password string) error {
	log.Println("Saving file")
	dirLocation, err := GetNotesLocation()
	if err != nil {
		return err
	}

	var (
		encryptedContent []byte
		encryptErr       error
	)
	if password != "" {
		encryptedContent, encryptErr = encryption.AesEncryptText([]byte(contents), password)
	} else {
		encryptedContent, encryptErr = encryption.AesEncryptText([]byte(contents), defaultKey)
	}

	if encryptErr != nil {
		return encryptErr
	}

	writeErr := os.WriteFile(dirLocation+name, encryptedContent, 0700)
	if writeErr != nil {
		fmt.Printf("Failed to write file: %v", writeErr)
		return writeErr
	}
	return nil
}

func ReadFile(fileName string, password string) string {
	locationPath, err := GetNotesLocation()
	if err != nil {
		return "Error"
	}

	encryptedContent, readErr := os.ReadFile(locationPath + "/" + fileName)
	if readErr != nil {
		fmt.Printf("Failed to read encrypted file: %v", readErr)
	}

	var (
		decryptedText []byte
		decryptErr    error
	)
	if password != "" {
		decryptedText, decryptErr = encryption.AesDecryptText([]byte(encryptedContent), password)
	} else {
		decryptedText, decryptErr = encryption.AesDecryptText([]byte(encryptedContent), defaultKey)
	}
	if decryptErr != nil {
		return "Failed to authorize"
	}
	return string(decryptedText)
}
