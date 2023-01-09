package utils

import (
	"encoding/json"
	"log"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/dstotijn/go-notion"
	"gitlab.42paris.fr/utilities/notion_service/pkg/methods"
	"gitlab.42paris.fr/utilities/notion_service/pkg/models"
)

func FindRowWithTitleInNotionDB(DBID string, title string, cursor *string) (*notion.Page, *models.ErrorResponse) {
	page, err := methods.FetchNotionDB(DBID)
	if err != nil {
		log.Printf("Could not connect to Notion Cluster Database: %s", err)
		return nil, &models.ErrorResponse{
			Code:    500,
			Message: "Bad Notion DB ID",
		}
	}

	// iterate in current page to find value
	for _, single := range page.DB.Results {
		themap, ok := single.Properties.(notion.DatabasePageProperties)
		if ok == false {
			log.Print("String Error")
		}
		// if Hostname exists return row
		for key, value := range themap {
			if key == "hostname" && (title == value.Title[0].PlainText) {
				return &single, nil
			}
		}
	}

	// if Hostname doesnt exists
	// iterate with cursor if page.DB.HasMore
	if page.DB.HasMore != false {
		log.Printf("Cluster DB HasMore Page: %t / next cursor : %s", page.DB.HasMore, reflect.ValueOf(page.DB.NextCursor))
		FindRowWithTitleInNotionDB(os.Getenv("DB_CLUSTER"), title, page.DB.NextCursor)
	}

	// if no more pages and row not found return all nil
	return nil, nil
}

func UpdateNotionDatabaseProperties(row *notion.Page, values interface{}) {}

func UpdateNotionRowProperties(row notion.DatabasePageProperties, values interface{}) (notion.DatabasePageProperties, error) {
	p, _ := json.Marshal(values) // encode to JSON
	var m map[string]interface{}
	json.Unmarshal(p, &m) // decode to map

	page := make(map[string]notion.DatabasePageProperty)

	for newkey, newvalue := range m {
	pageloop:
		for key, value := range row {
			if key == newkey {
				// modify value
				switch value.Type {
				case "title":
					value.Title[0].Text = &notion.Text{
						Content: newvalue.(string),
					}
					page[key] = value
					break pageloop
				case "select":
					value.ID = ""
					value.Select = &notion.SelectOptions{
						Name: newvalue.(string),
					}
					page[key] = value
					break pageloop
				case "rich_text":
					if len(value.RichText) > 0 {
						value.RichText[0] = notion.RichText{
							Type:      notion.RichTextTypeText,
							PlainText: newvalue.(string),
							Text: &notion.Text{
								Content: newvalue.(string),
							},
						}
					} else {
						value.RichText = append(value.RichText, notion.RichText{
							Type:      notion.RichTextTypeText,
							PlainText: newvalue.(string),
							Text: &notion.Text{
								Content: newvalue.(string),
							},
						})
					}
					page[key] = value
					break pageloop
				case "date":
					y, err := strconv.Atoi(newvalue.(string)[0:4])
					m, err := strconv.Atoi(newvalue.(string)[5:7])
					d, err := strconv.Atoi(newvalue.(string)[8:10])
					th, err := strconv.Atoi(newvalue.(string)[11:13])
					tm, err := strconv.Atoi(newvalue.(string)[14:16])
					ts, err := strconv.Atoi(newvalue.(string)[17:19])
					if err != nil {
						// i know...
						log.Println(err)
					}
					ti := time.Date(
						y,
						time.Month(m),
						d,
						th, tm, ts, 0, time.UTC,
					)
					value.Date.Start.Time = ti
					page[key] = value
					break pageloop
				case "checkbox":
					tmp, err := strconv.ParseBool(newvalue.(string))
					if err != nil {
						// i know...
						log.Println(err)
					}
					value.Checkbox = notion.BoolPtr(tmp)
					page[key] = value
					break pageloop
				case "number":
					break pageloop
				case "multi_select":
					break pageloop
				case "people":
					break pageloop
				case "files":
					break pageloop

				case "url":
					break pageloop
				case "email":
					break pageloop
				case "phone_number":
					break pageloop
				case "status":
					break pageloop
				case "formula":
					break pageloop
				case "relation":
					break pageloop
				case "rollup":
					break pageloop
				case "created_time":
					break pageloop
				case "created_by":
					break pageloop
				case "last_edited_time":
					break pageloop
				case "last_edited_by":
				default:
					// log.Println("Inexistant type")
				}
			}
		}
	}
	return page, nil
}

func FillNotionPropertyFromTemplate(page notion.DatabasePageProperties, values interface{}) (notion.DatabasePageProperties, error) {
	p, _ := json.Marshal(values) // encode to JSON
	var m map[string]interface{}
	json.Unmarshal(p, &m) // decode to map

	//iterate in values to set in properties
	for newkey, newvalue := range m {
	pageloop:
		for key, value := range page {
			//iterate in page properties
			// if propertie key = value key to set
			if key == newkey {
				// modify value
				switch value.Type {
				case "title":
					value.Title[0].Text = &notion.Text{
						Content: newvalue.(string),
					}
					break pageloop
				case "select":
					value.ID = ""
					value.Select = &notion.SelectOptions{
						Name: newvalue.(string),
					}
					page[key] = value
					break pageloop
				case "rich_text":
					if len(value.RichText) > 0 {
						value.RichText[0] = notion.RichText{
							Type:      notion.RichTextTypeText,
							PlainText: newvalue.(string),
							Text: &notion.Text{
								Content: newvalue.(string),
							},
						}
					} else {
						value.RichText = append(value.RichText, notion.RichText{
							Type:      notion.RichTextTypeText,
							PlainText: newvalue.(string),
							Text: &notion.Text{
								Content: newvalue.(string),
							},
						})
					}
					break pageloop
				case "date":
					y, err := strconv.Atoi(newvalue.(string)[0:4])
					m, err := strconv.Atoi(newvalue.(string)[5:7])
					d, err := strconv.Atoi(newvalue.(string)[8:10])
					th, err := strconv.Atoi(newvalue.(string)[11:13])
					tm, err := strconv.Atoi(newvalue.(string)[14:16])
					ts, err := strconv.Atoi(newvalue.(string)[17:19])
					if err != nil {
						// i know...
						log.Println(err)
					}
					ti := time.Date(
						y,
						time.Month(m),
						d,
						th, tm, ts, 0, time.UTC,
					)
					value.Date.Start.Time = ti
					break pageloop
				case "number":
					break pageloop
				case "multi_select":
					break pageloop
				case "people":
					break pageloop
				case "files":
					break pageloop
				case "checkbox":
					break pageloop
				case "url":
					break pageloop
				case "email":
					break pageloop
				case "phone_number":
					break pageloop
				case "status":
					break pageloop
				case "formula":
					break pageloop
				case "relation":
					break pageloop
				case "rollup":
					break pageloop
				case "created_time":
					break pageloop
				case "created_by":
					break pageloop
				case "last_edited_time":
					break pageloop
				case "last_edited_by":
				default:
					// log.Println("Inexistant type")
				}
			}
		}
	}

	// first search for empty values
	DeleteEmptyPropertyValues(page)
	return page, nil
}

func DeleteEmptyPropertyValues(page notion.DatabasePageProperties) notion.DatabasePageProperties {
	for key, value := range page {
		switch value.Type {
		case "title":
			if len(value.Title) == 0 {
				delete(page, key)
			}
		case "select":
			if value.Select == nil {
				delete(page, key)
			}
		case "rich_text":
			if len(value.RichText) == 0 {
				delete(page, key)
			}
		case "date":
			if value.Date == nil {
				delete(page, key)
			}
		case "number":
			if value.Number == nil {
				delete(page, key)
			}
		case "multi_select":
			if len(value.MultiSelect) == 0 {
				delete(page, key)
			}
		case "people":
			if len(value.People) == 0 {
				delete(page, key)
			}
		case "files":
			if len(value.Files) == 0 {
				delete(page, key)
			}
		// case "checkbox":
		// 	if value.Checkbox == false {delete(page, key)}
		case "url":
			if value.URL == nil {
				delete(page, key)
			}
		case "email":
			if value.Email == nil {
				delete(page, key)
			}
		case "phone_number":
			if value.PhoneNumber == nil {
				delete(page, key)
			}
		case "status":
			if value.Status == nil {
				delete(page, key)
			}
		case "formula":
			delete(page, key)
		case "relation":
			if len(value.Relation) == 0 {
				delete(page, key)
			}
		case "rollup":
			if value.Rollup == nil {
				delete(page, key)
			}
		case "created_time":
			delete(page, key)
		case "created_by":
			delete(page, key)
		case "last_edited_time":
			delete(page, key)
		case "last_edited_by":
			delete(page, key)
		default:
			// log.Println("Inexistant type")
		}
	}
	return page
}
