package cyoa

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTemplate))
}

var defaultHandlerTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>Choose Your Own Adventure</title>
</head>
<body>
  <h1>{{.Title}}</h1>
  {{range .Paragraphs}}
  <p>{{.}}</p>
  {{end}}
  <ul>
    {{range .Options}}
      <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
    {{end}}
  </ul>
</body>
</html>
`

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
type handler struct {
	s Story
}

func NewHandler(s Story) http.Handler {
	return handler{s}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)

	if path == "" || path == "/" {
		path = "/intro"
	}

	path = path[1:]

	if chapter, ok := h.s[path]; ok {
		err := tpl.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found", http.StatusNotFound)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func JSONStory(r io.Reader) (Story, error) {
	decoder := json.NewDecoder(r)

	var story Story
	if err := decoder.Decode(&story); err != nil {
		panic(err)
	}
	return story, nil
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
	var story []Chapter
	data, err := ioutil.ReadFile("story.json")
	fmt.Println(string(data))
	check(err)

	json.Unmarshal(data, &story)
	fmt.Println(story)
}
func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

// the architecture here is that people are going to get a prompt, then a choice.
