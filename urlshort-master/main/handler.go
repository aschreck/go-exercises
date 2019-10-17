package main

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

type parsedYAML []struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...

	// TODO: need to handle case when URL is not in the list of URLs
	return func(w http.ResponseWriter, r *http.Request) {
		returnedURL := ""
		for registeredURL := range pathsToUrls {
			incomingURL := r.URL.String()
			if incomingURL == registeredURL {
				returnedURL = pathsToUrls[registeredURL]
			}
		}
		if !(returnedURL == "") {
			http.Redirect(w, r, returnedURL, 302)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	ymlStruct := parsedYAML{}
	parsedYaml := yaml.Unmarshal(yml, &ymlStruct)
	fmt.Println(parsedYaml)
	// this takes in YAML as a byte slice. need to parse the yaml

	// unmarshall the bytle slice into a struct which contains the fields.
	return nil, nil
}
