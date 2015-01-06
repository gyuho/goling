package thesaurus

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/gyuho/stringx"
)

// func main() {
// 	api, err := New()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// map[verb:map[syn:[accredit account ascribe assign attribute bank calculate impute rely swear trust] ant:[debit]] noun:map[syn:[recognition credit entry deferred payment course credit citation cite acknowledgment reference mention quotation accomplishment accounting entry achievement annotation approval assets attainment commendation entry ledger entry notation note payment title] ant:[cash debit]]]
// 	rs, err := api.Get("health")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(rs)
// }

// API implements http://words.bighugelabs.com/ API client.
type API struct {
	key string
	*http.Client
}

// New returns a new API with default http.Client.
func New() (*API, error) {
	akey := os.Getenv("THESAURUS_KEY")
	if akey == "" {
		return nil, errors.New("no environment variable set THESAURUS_KEY")
	}
	api := API{
		key:    akey,
		Client: http.DefaultClient,
	}
	return &api, nil
}

// NewCustom returns a new API with customized http.Client.
func NewCustom(client *http.Client) (*API, error) {
	akey := os.Getenv("THESAURUS_KEY")
	if akey == "" {
		return nil, errors.New("no environment variable set THESAURUS_KEY")
	}
	api := API{
		key:    akey,
		Client: client,
	}
	return &api, nil
}

const endpoint = "http://words.bighugelabs.com/api/2/%s/%s/json"

// Get returns the synonyms of an input word.
func (a *API) Get(word string) ([]string, error) {
	nword := strings.TrimSpace(word)
	nword = strings.ToLower(nword)
	url := fmt.Sprintf(endpoint, a.key, nword)
	resp, err := a.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
	result := map[string]map[string][]string{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	smap := make(map[string]bool)
	for _, val := range result {
		for k, v := range val {
			if k == "syn" {
				for _, elem := range v {
					smap[elem] = true
				}
			}
		}
	}
	temSplice := []string{}
	for key := range smap {
		words := stringx.SplitToWords(strings.ToLower(key))
		temSplice = append(temSplice, words...)
	}
	found := make(map[string]bool)
	for _, elem := range temSplice {
		found[elem] = true
	}
	slice := []string{}
	for k := range found {
		slice = append(slice, k)
	}
	sort.Strings(slice)
	return slice, nil
}
