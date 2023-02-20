package main

//import (
//	"archive/zip"
//	"bytes"
//	"io/ioutil"
//	"net/http"
//	"os"
//	"reflect"
//	"strconv"
//	"strings"
//	"time"
//)
//
//func ExportData(writer http.ResponseWriter, req *http.Request, session sessions.Session) (int, string) {
//
//	headers := HeaderCreation{
//		OriginalRequest: req,
//		Session:         session,
//	}
//
//	qs := req.URL.Query()
//
//	if len(qs["collectionID"]) != 1 {
//		return 400, "ERROR: Must submit one collectionID in query string"
//	}
//	if len(qs["fileType"]) != 1 {
//		return 400, "ERROR: Must submit one fileType in query string"
//	}
//
//	collID := qs["collectionID"][0]
//	fileType := qs["fileType"][0]
//
//	url := "http://" + config.Data.Address + "/api/" + collID
//	response, err := httpClient.DoSystemRequest("GET", url, nil, headers)
//
//	if err != nil {
//		return 500, "ERROR: Could not resolve DataURL/api" + err.Error()
//	} else {
//		contents, err := ioutil.ReadAll(response.Body)
//		response.Body.Close()
//
//		if err != nil {
//			return 400, "ERROR: Response from Platform unreadable"
//		}
//
//		buf := new(bytes.Buffer)
//
//		w := zip.NewWriter(buf)
//
//		file, err := w.Create(collID + "." + fileType)
//		if err != nil {
//			return 400, "ERROR: Unable to create zip file with name of: " + collID + " and type of: " + fileType + "; " + err.Error()
//		}
//
//		switch fileType {
//		case "csv":
//
//			rows, err := setupCSVRows(contents)
//
//			if err != nil {
//				return 400, err.Error()
//			}
//
//			_, err = file.Write(rows)
//			if err != nil {
//				return 400, "Unable to write CSV to zip file; " + err.Error()
//			}
//		case "json":
//			_, err := file.Write(contents)
//			if err != nil {
//				return 400, err.Error()
//			}
//		} // end switch
//
//		err = w.Close()
//		if err != nil {
//			return 400, "ERROR: Unable to close zip file writer; " + err.Error()
//		}
//
//		//create fileName based on collectionID and current time
//		fileAddress := collID + strconv.FormatInt(time.Now().Unix(), 10)
//
//		//write the zipped file to the disk
//		ioutil.WriteFile(fileAddress+".zip", buf.Bytes(), 0777)
//
//		return 200, fileAddress
//	} //end else
//}
//
//func ReturnFile(writer http.ResponseWriter, req *http.Request) {
//	queries := req.URL.Query()
//	fullFileName := queries["fullFileName"][0]
//	http.ServeFile(writer, req, fullFileName)
//	//delete file from server once it has been served
//	//defer os.Remove(fullFileName)
//}
//
//func setupCSVRows(contents []byte) ([]byte, error) {
//	//unmarshal into interface because we don't know json structure in advance
//	var collArr interface{}
//	jsonErr := json.Unmarshal(contents, &collArr)
//
//	if jsonErr != nil {
//		return nil, errors.New("ERROR: Unable to parse JSON")
//	}
//
//	//had to do some weird stuff here, not sure if it's the best method
//	s := reflect.ValueOf(collArr)
//	var rows bytes.Buffer
//	var headers []string
//
//	for i := 0; i < s.Len(); i++ {
//		var row []string
//		m := s.Index(i).Interface()
//
//		m2 := m.(map[string]interface{})
//
//		for k, v := range m2 {
//
//			if i == 0 {
//				if k != "item_id" {
//					headers = append(headers, k)
//				}
//			}
//			if k != "item_id" {
//				row = append(row, v.(string))
//			}
//		}
//
//		if i == 0 {
//			headersString := strings.Join(headers, ",")
//			rows.WriteString(headersString + "\n")
//		}
//		rowsString := strings.Join(row, ",")
//		rows.WriteString(rowsString + "\n")
//	}
//
//	return rows.Bytes(), nil
//}
//
//func SecondReturnFile(writer http.ResponseWriter, req *http.Request) {
//	queries := req.URL.Query()
//	fullFileName := queries["fullFileName"][0]
//
//	writer.Header().Set("Content-type", "application/zip")
//	http.ServeFile(writer, req, fullFileName)
//	//delete file from server once it has been served
//	defer os.Remove(fullFileName)
//}
