package clean

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)


func isInArray(search string, array []string) bool {
    for _, item := range array {
        if item == search {
            fmt.Printf("Found extension: %s\n", search)
            return true
        }
    }
    fmt.Printf("Extension not found: %s\n", search)
    return false
}

func mvFilesToDir(sourcePath, destinationPath string) error {
    _, err := os.Stat(sourcePath)
    if err != nil {
        return err
    }

    err = os.Rename(sourcePath, destinationPath)
    if err != nil {
        return err
    }

    fmt.Printf("Moved file %s to %s\n", sourcePath, destinationPath)
    return nil
}

func checkExtension(filePath string) {
    u, _ := user.Current()

    docsFolder := filepath.Join(u.HomeDir, "Desktop", "foldery", "downloads", "docs")
    multimediaFolder := filepath.Join(u.HomeDir, "Desktop", "foldery", "downloads", "img")
    appFolder := filepath.Join(u.HomeDir, "Desktop", "foldery", "downloads", "apps")
    archiveFolder := filepath.Join(u.HomeDir, "Desktop", "foldery   ", "archive")

    extension := filepath.Ext(strings.TrimSpace(filePath))
    extension = extension[1:]
    fmt.Println("Extension: ", extension)

    var MultimediaExtensions = []string{
        "png", "jpg", "jpeg", "gif", "bmp", "tiff", "tif", "webp", "svg", "JFIF",
        "mp4", "avi", "mkv", "wmv", "mov", "flv", "m4v",
    }

    var DocsExtensions = []string{
        "txt", "doc", "docx", "odt", "rtf", "ppt", "pptx", "odp", "xls", "xlsx", "ods", "pdf",
    }

    var AppExtensions = []string{
        "exe", "msi", "dmg", "app", "deb", "rpm", "apk",
    }

    var ArchiveExtensions = []string{
        "zip", "rar", "7z", "tar", "gz", "bz2",
    }

    switch {
    case isInArray(extension, MultimediaExtensions):
        err := mvFilesToDir(filePath, filepath.Join(multimediaFolder, filepath.Base(filePath)))
        if err != nil {
            fmt.Println("Error moving file:", err)
        }
    case isInArray(extension, DocsExtensions):
        err := mvFilesToDir(filePath, filepath.Join(docsFolder, filepath.Base(filePath)))
        if err != nil {
            fmt.Println("Error moving file:", err)
        }
    case isInArray(extension, AppExtensions):
        err := mvFilesToDir(filePath, filepath.Join(appFolder, filepath.Base(filePath)))
        if err != nil {
            fmt.Println("Error moving file:", err)
        }
    case isInArray(extension, ArchiveExtensions):
        err := mvFilesToDir(filePath, filepath.Join(archiveFolder, filepath.Base(filePath)))
        if err != nil {
            fmt.Println("Error moving file:", err)
        }
    }
}
