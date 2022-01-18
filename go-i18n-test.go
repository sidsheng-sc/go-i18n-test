package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// tutorial: https://lokalise.com/blog/go-internationalization-using-go-i18n/
// package: https://github.com/nicksnyder/go-i18n

func main() {
	//partOne()
	//partTwo()
}

var localizer *i18n.Localizer
var bundle *i18n.Bundle

func init() {
	// partTwo()
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile("resources/en-US.json")
	bundle.LoadMessageFile("resources/fr-FR.json")
	localizer = i18n.NewLocalizer(bundle, language.AmericanEnglish.String(), language.French.String())

	// partThree()
	http.HandleFunc("/setlang/", SetLangPreferences)
	http.HandleFunc("/localize/", Localize)
	http.ListenAndServe(":8080", nil)

	// GET http://localhost:8080/localize?msg=hello
	// GET http://localhost:8080/setlang?lang=fr-FR
	// GET http://localhost:8080/setlang, Accept-Language header = fr-FR
}

func partOne() {
	messageEn := i18n.Message{
		ID:    "hello",
		Other: "Hello!",
	}
	messageFr := i18n.Message{
		ID:    "hello",
		Other: "Bonjour!",
	}

	bundle := i18n.NewBundle(language.English)
	bundle.AddMessages(language.English, &messageEn)
	bundle.AddMessages(language.French, &messageFr)

	localiser := i18n.NewLocalizer(bundle,
		language.French.String(),
		language.English.String())
	localiseConfig := i18n.LocalizeConfig{
		MessageID: "hello",
	}

	localisation, _ := localiser.Localize(&localiseConfig)

	fmt.Println(localisation)
}

func partTwo() {
	localizeConfigWelcome := i18n.LocalizeConfig{
		MessageID: "welcome",
	}
	localizationUsingJson, _ := localizer.Localize(&localizeConfigWelcome)
	fmt.Println(localizationUsingJson)
}

// partThree
func SetLangPreferences(_ http.ResponseWriter, request *http.Request) {
	lang := request.FormValue("lang")
	accept := request.Header.Get("Accept-Language")
	localizer = i18n.NewLocalizer(bundle, lang, accept)
}

func Localize(responseWriter http.ResponseWriter, request *http.Request) {
	valToLocalize := request.URL.Query().Get("msg")
	localizeConfig := i18n.LocalizeConfig{
		MessageID: valToLocalize,
	}
	localization, _ := localizer.Localize(&localizeConfig)
	fmt.Fprintln(responseWriter, localization)
}
