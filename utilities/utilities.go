package utilities

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"os"
	"setgo-api/structs"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/joho/godotenv"
)

func GetData(url string, cacheFile string) ([]structs.App, error) {
	if CacheIsCurrent(cacheFile) {
		apps, err := GetDataFromCache(cacheFile)
		if err != nil {
			return nil, err
		} else {
			var ret []structs.App
			for _, a := range apps {
				a.ReleaseDateString = EpochToHumanDate(a.ReleaseDate)
				ret = append(ret, a)
			}
			return ret, nil
		}

	} else {
		return GetDataFromSetapp(url, cacheFile)
	}
}

func GetDataFromCache(cacheFile string) ([]structs.App, error) {
	var appList []structs.App

	dataApps, err := os.ReadFile(cacheFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(dataApps), &appList)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	// apps = append(apps, appList...)
	return appList, nil
}

func GetDataFromSetapp(url string, cacheFile string) ([]structs.App, error) {
	var apps []structs.App
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	doc.Find("div[data-apps]").Each(func(i int, s *goquery.Selection) {
		dataApps, exists := s.Attr("data-apps")
		if !exists {
			fmt.Println("data-apps attribute not found")
			return
		}

		dataApps = html.UnescapeString(dataApps)

		var appList []structs.App
		err := json.Unmarshal([]byte(dataApps), &appList)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return
		}

		apps = append(apps, appList...)
	})

	// we have the data, put it in the cache
	appsJson, _ := json.Marshal(apps)
	err = os.WriteFile(cacheFile, appsJson, 0644)
	if err != nil {
		panic(err)
	}

	return apps, nil
}

func CacheIsCurrent(cacheFile string) bool {
	if fileinfo, err := os.Stat(cacheFile); err == nil {
		if time.Since(fileinfo.ModTime()) > 6*time.Hour {
			return false // exists but is old
		} else {
			return true // exists and is not old
		}

	} else {
		return false // does not exist
	}
}

func GetConfigValue(key string) string {
	godotenv.Load()
	return os.Getenv(key)
}

func RemoveDuplicate[T comparable](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func LoadSubCats(subCatFile string) ([]structs.Subcat, error) {
	var subCatList []structs.Subcat

	dataSubcats, err := os.ReadFile(subCatFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(dataSubcats), &subCatList)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	return subCatList, nil
}

func EpochToHumanDate(epoch int64) string {
	t := time.Unix(epoch, 0)
	return t.Format("2006-01-02")
}
