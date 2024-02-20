//Webserver/split.go
package groupie

import (
    "strings"
    "html/template"
)

// Define a custom function to split a string
func SplitString(s string, sep string) []string {
    return strings.Split(s, sep)
}

// Add a function to initialize the template with custom functions
func LoadTemplates() (*template.Template, error) {
    funcMap := template.FuncMap{
        "split": strings.Split,
    }

    tmpl, err := template.New("groupie.html").Funcs(funcMap).ParseFiles("../Webserver/groupie.html")
    if err != nil {
        return nil, err
    }
    return tmpl, nil
}
