package tct

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/nyulibraries/dlts/enm/cache"
	"github.com/nyulibraries/dlts/enm/util"
)

const TctBaseUrl = "https://nyuapi.infoloom.nyc"

var Source string

var TctApiEndpoints = map[string]string{
	"EpubDetail": "/api/epub/document/",
	"EpubsAll": "/api/epub/document/all/",
	"IndexPatternsAll": "/api/epub/index-pattern/all/",
	"Location": "/api/epub/location/",
	"NamesAll": "/api/hit/hits/all/",
	"TopicsAll": "/api/hit/basket/all/",
	"TopicDetail": "/api/hit/basket/",
}

func GetResponseBody(params ...string) (body []byte) {
	var request = params[0]

	id := ""
	if len(params) > 1 {
		id = params[1]
	}

	cacheFile := cache.CacheFile(request, id)

	if (Source == "cache") {
		fmt.Println("Fetching data from cache")

		cachedData, err := ioutil.ReadFile(cacheFile)
		if err != nil {
			panic(err.Error())
		}
		body = cachedData
	} else if (Source == "tct-api"){
		fmt.Println("Fetching data from TCT server")

		var url = TctBaseUrl + TctApiEndpoints[request]
		if id != "" {
			url += id
		}

		res, err := http.Get(url)
		if err != nil {
			panic(err.Error())
		}

		body, err = ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err.Error())
		}

		err = ioutil.WriteFile(cacheFile, body, 0600)
	} else {
		panic("Unknown source: " + Source)
	}

	return
}

func GetEpubDetail(epubId int) (epubDetail EpubDetail) {
	epubIdString := strconv.Itoa(epubId)
	responseBody := GetResponseBody("EpubDetail", epubIdString)

	err := json.Unmarshal(responseBody, &epubDetail)
	if err != nil {
		panic(err.Error())
	}

	return
}

func GetEpubsAll() (epubsList []Epub) {
	responseBody := GetResponseBody("EpubsAll")

	err := json.Unmarshal(responseBody, &epubsList)
	if err != nil {
		panic(err.Error())
	}

	return
}

func GetIndexPatternsAll() (indexPatternsList []IndexPattern) {
	responseBody := GetResponseBody("IndexPatternsAll")

	err := json.Unmarshal(responseBody, &indexPatternsList)
	if err != nil {
		panic(err.Error())
	}

	return
}

func GetLocation(locationId int) (location Location) {
	locationIdString := strconv.Itoa(locationId)
	responseBody := GetResponseBody("Location", locationIdString)

	// For some reason this endpoint returns an array containing a single location element instead of just a single
	// location element.  Bug?
	var locations []Location
	err := json.Unmarshal(responseBody, &locations)
	if err != nil {
		panic(err.Error())
	}

	location = locations[0]

	return
}

func GetNamesAll() (namesList []Name) {
	responseBody := GetResponseBody("NamesAll")

	err := json.Unmarshal(responseBody, &namesList)
	if err != nil {
		panic(err.Error())
	}

	return
}

func GetTopicsAll() (topicsList []Topic) {
	responseBody := GetResponseBody("TopicsAll")

	err := json.Unmarshal(responseBody, &topicsList)
	if err != nil {
		panic(err.Error())
	}

	return
}

func GetTopicDetail(topicId int) (topicDetails TopicDetail) {
	topicIdString := strconv.Itoa(topicId)
	responseBody := GetResponseBody("TopicDetail", topicIdString)

	err := json.Unmarshal(responseBody, &topicDetails)
	if err != nil {
		panic(err.Error())
	}

	return
}

func GetApiRequestTypes() ([]string){
	return util.GetMapKeys(TctApiEndpoints)
}

// The field declarations for the types below were generated using gojson:
// github.com/ChimeraCoder/gojson/gojson
//
// Example: curl https://nyuapi.infoloom.nyc/api/epub/document/all/?format=json | gojson -name=EpubsAll`
//
// For now just declaring types for single objects, not collections: e.g. Epub, not EpubsAll

// /api/epub/document/all/
type Epub struct {
	Author    string `json:"author"`
	ID        int64  `json:"id"`
	Isbn      string `json:"isbn"`
	Publisher string `json:"publisher"`
	Title     string `json:"title"`
}

// /api/epub/document/DOCUMENT_ID/
type EpubDetail struct {
	Author    string `json:"author"`
	Isbn      string `json:"isbn"`
	Locations []struct {
		Content        string `json:"content"`
		ID             int64  `json:"id"`
		Localid        string `json:"localid"`
		SequenceNumber int64  `json:"sequence_number"`
	} `json:"locations"`
	Publisher string `json:"publisher"`
	Title     string `json:"title"`
}

// /api/epub/index-pattern/all/
type IndexPattern struct {
	Description string `json:"description"`
	Documents   []struct {
		ID    int64  `json:"id"`
		Title string `json:"title"`
	} `json:"documents"`
	IndicatorsOfOccurrenceRange         []string `json:"indicators_of_occurrence_range"`
	InlineSeeAlsoEnd                    string   `json:"inline_see_also_end"`
	InlineSeeAlsoStart                  string   `json:"inline_see_also_start"`
	InlineSeeEnd                        string   `json:"inline_see_end"`
	InlineSeeStart                      string   `json:"inline_see_start"`
	Name                                string   `json:"name"`
	PagenumberCSSSelectorPattern        string   `json:"pagenumber_css_selector_pattern"`
	PagenumberPreStrings                []string `json:"pagenumber_pre_strings"`
	PagenumberXpathPattern              string   `json:"pagenumber_xpath_pattern"`
	SeeAlsoSplitStrings                 []string `json:"see_also_split_strings"`
	SeeSplitStrings                     []string `json:"see_split_strings"`
	SeparatorBeforeFirstSubentry        string   `json:"separator_before_first_subentry"`
	SeparatorBetweenEntryAndOccurrences string   `json:"separator_between_entry_and_occurrences"`
	SeparatorBetweenSeealsos            string   `json:"separator_between_seealsos"`
	SeparatorBetweenSees                string   `json:"separator_between_sees"`
	SeparatorBetweenSubentries          string   `json:"separator_between_subentries"`
	SeparatorSeeSubentry                string   `json:"separator_see_subentry"`
	SubentryClasses                     []string `json:"subentry_classes"`
	XpathEntry                          string   `json:"xpath_entry"`
	XpathOccurrenceLink                 string   `json:"xpath_occurrence_link"`
	XpathSee                            string   `json:"xpath_see"`
	XpathSeealso                        string   `json:"xpath_seealso"`
}

// /api/epub/location/LOCATION_ID/
type Location struct {
	Content struct {
		ContentDescriptor      string `json:"content_descriptor"`
		ContentUniqueIndicator string `json:"content_unique_indicator"`
		Text                   string `json:"text"`
	} `json:"content"`
	Context  interface{} `json:"context"`
	Document struct {
		Author    string `json:"author"`
		ID        int64  `json:"id"`
		Isbn      string `json:"isbn"`
		Publisher string `json:"publisher"`
		Title     string `json:"title"`
	} `json:"document"`
	ID             int64         `json:"id"`
	Localid        string        `json:"localid"`
	NextLocationID int64         `json:"next_location_id"`
	Occurrences    []interface{} `json:"occurrences"`
	Pagenumber     struct {
		CSSSelector   string `json:"css_selector"`
		Filepath      string `json:"filepath"`
		PagenumberTag string `json:"pagenumber_tag"`
		Xpath         string `json:"xpath"`
	} `json:"pagenumber"`
	PreviousLocationID interface{} `json:"previous_location_id"`
}

// /api/hit/hits/all/
type Name struct {
	Basket    int64  `json:"basket"`
	Hidden    bool   `json:"hidden"`
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Preferred bool   `json:"preferred"`
	Scope     string `json:"scope"`
}

// /api/hit/basket/all/
type Topic struct {
	DisplayName string `json:"display_name"`
	ID          int64  `json:"id"`
}

// /api/hit/basket/BASKET_ID/
type TopicDetail struct {
	Basket struct {
		DisplayName string `json:"display_name"`
		ID          int64  `json:"id"`
		Occurs      []struct {
			Basket   int64 `json:"basket"`
			ID       int64 `json:"id"`
			Location struct {
				Document struct {
					Author string `json:"author"`
					Title  string `json:"title"`
				} `json:"document"`
				ID             int64  `json:"id"`
				Localid        string `json:"localid"`
				SequenceNumber int64  `json:"sequence_number"`
			} `json:"location"`
		} `json:"occurs"`
		TopicHits []struct {
			Bypass    bool   `json:"bypass"`
			Hidden    bool   `json:"hidden"`
			ID        int64  `json:"id"`
			Name      string `json:"name"`
			Preferred bool   `json:"preferred"`
			Scope     struct {
				ID    int64  `json:"id"`
				Scope string `json:"scope"`
			} `json:"scope"`
		} `json:"topic_hits"`
	} `json:"basket"`
	Relations []struct {
		Basket struct {
			DisplayName string `json:"display_name"`
			ID          int64  `json:"id"`
		} `json:"basket"`
		Direction    string `json:"direction"`
		ID           int64  `json:"id"`
		Relationtype struct {
			ID          int64  `json:"id"`
			RoleFrom    string `json:"role_from"`
			RoleTo      string `json:"role_to"`
			Rtype       string `json:"rtype"`
			Symmetrical bool   `json:"symmetrical"`
		} `json:"relationtype"`
	} `json:"relations"`
}
