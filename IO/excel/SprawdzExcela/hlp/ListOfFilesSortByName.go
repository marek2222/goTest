package hlp

import (
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"
)

// ListOfFiles in directory
func ListOfFiles(extension, directoryName string) ([]string, error) {
	lists := []string{}

	files, err := ioutil.ReadDir(directoryName)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}
		ext := strings.ToLower(filepath.Ext(f.Name()))
		if ext != extension {
			continue
		}
		plikPath := filepath.Join(directoryName, f.Name())
		lists = append(lists, plikPath)
	}

	return lists, nil
}

// ListOfFilesSortByName in directory
func ListOfFilesSortByName(extension, directoryName string) ([]string, error) {
	lists := []string{}

	files, err := ioutil.ReadDir(directoryName)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}
		ext := strings.ToLower(filepath.Ext(f.Name()))
		if ext != extension {
			continue
		}
		plikPath := filepath.Join(directoryName, strings.ToLower(f.Name()))
		lists = append(lists, plikPath)
	}

	sort.Strings(lists)

	return lists, nil
}

// ListOfFilesSortByDate in directory
func ListOfFilesSortByDate(extension, directoryName string) ([]string, error) {
	lists := []string{}

	files, err := ioutil.ReadDir(directoryName)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}
		ext := strings.ToLower(filepath.Ext(f.Name()))
		if ext != extension {
			continue
		}
		plikPath := filepath.Join(directoryName, f.Name())
		lists = append(lists, plikPath)
	}

	// TODO: handle the error!
	sort.Slice(files, func(i, j int) bool {
		return files[i].ModTime().Before(files[j].ModTime())
	})

	lists = []string{}
	for _, f := range files {
		plikPath := filepath.Join(directoryName, f.Name())
		lists = append(lists, plikPath)
	}

	return lists, nil
}

// func sortedFiles(files []os.FileInfo) []os.FileInfo {
//  sort.Slice(files, func(i, j int) bool {
//      return files[i].ModTime().Unix() < files[j].ModTime().Unix()
//  })
//  return files
// }

// // SortByDateOfFile sort file by date
// func SortByDateOfFile(list []string)  {
//  sort.Slice(timeSlice, func(i, j int) bool {
//      return timeSlice[i].date.Before(timeSlice[j].date)
//  })
// }
