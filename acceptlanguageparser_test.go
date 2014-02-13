// Copyright 2014 struktur AG. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goacceptlanguageparser

import (
	"reflect"
	"testing"
)

func TestAcceptLanguageParser_de_en_us(t *testing.T) {
	languages := "de,en;q=0.7,en-us;q=0.3"
	if expected, actual := []string{"de", "en", "en-us"}, ParseAcceptLanguage(languages, []string{}); !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected result '%+v', but was '%+v'", expected, actual)
	}
}

func TestAcceptLanguageParser_en_us_de(t *testing.T) {
	languages := "en-US,en;q=0.8,de;q=0.6"
	if expected, actual := []string{"en-us", "en", "de"}, ParseAcceptLanguage(languages, []string{}); !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected result '%+v', but was '%+v'", expected, actual)
	}
}

func TestAcceptLanguageParser_fr_de(t *testing.T) {
	languages := "fr-FR,de-DE;q=0.5"
	if expected, actual := []string{"fr-fr", "de-de"}, ParseAcceptLanguage(languages, []string{}); !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected result '%+v', but was '%+v'", expected, actual)
	}
}

func TestAcceptLanguageParser_rfcExample(t *testing.T) {
	languages := "da, en-gb;q=0.8, en;q=0.7"
	if expected, actual := []string{"da", "en-gb", "en"}, ParseAcceptLanguage(languages, []string{}); !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected result '%+v', but was '%+v'", expected, actual)
	}
}

func TestAcceptLanguageParser_rfcExampleTrim(t *testing.T) {
	languages := " da, en-gb;q=0.8, en;q=0.7"
	if expected, actual := []string{"da", "en-gb", "en"}, ParseAcceptLanguage(languages, []string{}); !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected result '%+v', but was '%+v'", expected, actual)
	}
}

func TestAcceptLanguageParser_rfcExampleNoQuality(t *testing.T) {
	languages := "da, en-gb;q=, en;q=0.7"
	if expected, actual := []string{"da", "en-gb", "en"}, ParseAcceptLanguage(languages, []string{}); !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected result '%+v', but was '%+v'", expected, actual)
	}
}

func TestAcceptLanguageParser_rfcExample_supportedLanguages(t *testing.T) {
	languages := "da, en-gb;q=, en;q=0.7"
	if expected, actual := []string{"da", "en-gb", "en"}, ParseAcceptLanguage(languages, []string{"en-gb", "da", "en", "de", "fr", "cz"}); !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected result '%+v', but was '%+v'", expected, actual)
	}
}

func TestAcceptLanguageParser_rfcExample_oneSupportedLanguage(t *testing.T) {
	languages := "da, en-gb;q=, en;q=0.7, de"
	if expected, actual := []string{"de"}, ParseAcceptLanguage(languages, []string{"de"}); !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected result '%+v', but was '%+v'", expected, actual)
	}
}

func Benchmark_TestAcceptLanguageParser_de_en_us(b *testing.B) {
	languages := "en-US,en;q=0.8,de;q=0.6"
	for i := 0; i < b.N; i++ { //use b.N for looping
		ParseAcceptLanguage(languages, []string{})
	}
}

func Benchmark_TestAcceptLanguageParser_supportedLanguages(b *testing.B) {
	languages := "en-US,en;q=0.8,en-UK;q=0.7;de;q=0.6;ru=0.5"
	for i := 0; i < b.N; i++ {
		ParseAcceptLanguage(languages, []string{"en", "de"})
	}
}
