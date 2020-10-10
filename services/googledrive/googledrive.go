package googledrive

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"gitlab.com/velvetkeyboard/go-quickbackup/utils"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

const (
	FOLDER_MIME_TYPE         = "application/vnd.google-apps.folder"
	QUICKBACKUPZIP_MIME_TYPE = "fsbackup/zip.file"
)

type GoogleDrive struct {
	Client *drive.Service
}

func (gd *GoogleDrive) Init(jsonConfigPath string) {
	// gd.Client, _ = drive.New(
	// 	getOAuth2Client(
	// 		getConfigFromJSONFile(jsonConfigPath),
	// 	),
	// )
	gd.Client = getService(jsonConfigPath)
}

func (gd *GoogleDrive) GetEntityId(entityName string, mimeType string, skip_trash bool) string {
	query := "name contains '" + entityName + "'"
	if mimeType != "" {
		query += " and mimeType = '" + mimeType + "'"
	}
	query += " and trashed = " + strconv.FormatBool(!skip_trash) + ""

	println("query", query)

	r, err := gd.Client.Files.List().
		PageSize(1).
		Q(query).
		Fields("files(id)").
		Do()

	if err != nil {
		println("Entity Id Retrieval Failed: ", err)
		return "" // TODO should we panic instead of return empty string?
	}

	if len(r.Files) == 0 {
		return ""
	}

	return r.Files[0].Id
}

func (gd *GoogleDrive) GetFolderId(folderName string) string {
	return gd.GetEntityId(folderName, FOLDER_MIME_TYPE, true)
}

func (gd *GoogleDrive) GetFileId(fileName string) string {
	// return gd.GetEntityId(fileName, QUICKBACKUPZIP_MIME_TYPE, true)
	return gd.GetEntityId(fileName, "", true)
}

func (gd *GoogleDrive) ListEntitiesFrom(
	parentEntityId string,
	mimeType string,
	skip_trash bool,
) []*drive.File {
	var nextPageToken string
	ret := []*drive.File{}
	for {
		query := "'" + parentEntityId + "' in parents"
		if mimeType != "" {
			query += " and mimeType = '" + mimeType + "'"
		}
		query += " and "
		query += "trashed = " + strconv.FormatBool(!skip_trash) + ""

		filesListCall := gd.Client.Files.List().
			PageSize(5).
			Q(query).
			Fields("nextPageToken, files(id, name)")

		if nextPageToken != "" {
			filesListCall.PageToken(nextPageToken)
		}

		r, err := filesListCall.Do()

		if err != nil {
			return []*drive.File{}
		}

		if len(r.Files) > 0 {
			ret = append(ret, r.Files...)
		}

		if nextPageToken = r.NextPageToken; nextPageToken == "" {
			break
		}
	}

	return ret
}

func (gd *GoogleDrive) ListFilesFromFolder(parentFolderName string) []*drive.File {
	parentFolderId := gd.GetFolderId(parentFolderName)
	return gd.ListEntitiesFrom(parentFolderId, "", true)
}

func (gd *GoogleDrive) DownloadFile(fileId string) {
	r, _ := gd.Client.Files.Get(fileId).Download()
	// if err != nil {
	// 	println("err:", err)
	// } else {
	// 	fmt.Printf("%v\n", r)
	// }
	fmt.Printf("%v\n", r.Body)
	bodyBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		println(err)
	}

	println(string(bodyBytes)) // TODO persist bodyBytes to a file, pls!
}

// Helpers to initialize drive.Service struct ----------------------------------

func getService(configPath string) *drive.Service {
	c, _ := drive.New(
		getOAuth2Client(
			getConfigFromJSONFile(configPath),
		),
	)
	return c
}

// Retrieve a token, saves the token, then returns the generated client.
func getOAuth2Client(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the
	// first time.
	tokFile := utils.ExpandUser("~/.config/quickbackup/token.json")
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)
	println("Paste the token:")
	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func getConfigFromJSONFile(filePath string) *oauth2.Config {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, drive.DriveScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	return config
}
