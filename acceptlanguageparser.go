package goacceptlanguageparser

import (
	"sort"
	"strconv"
	"strings"
)

type Language struct {
	name    string
	quality float64
}

type By func(l1, l2 *Language) bool

type languageSorter struct {
	languages []Language
	by        func(l1, l2 *Language) bool
}

func (by By) Sort(languages []Language) {
	ls := &languageSorter{
		languages: languages,
		by:        by,
	}
	sort.Sort(ls)
}

func (s *languageSorter) Len() int {
	return len(s.languages)
}

func (s *languageSorter) Swap(i, j int) {
	s.languages[i], s.languages[j] = s.languages[j], s.languages[i]
}

func (s *languageSorter) Less(i, j int) bool {
	return s.by(&s.languages[i], &s.languages[j])
}

func ParseAcceptLanguage(languages string, supported_languages []string) []string {
	langs := []Language{}
	browser_pref_langs := strings.Split(languages, ",")

	i := 0
	length := len(browser_pref_langs)

	for _, lang := range browser_pref_langs {
		// Format strings.
		lang = strings.Replace(strings.ToLower(strings.TrimSpace(lang)), "_", "-", 0)

		if lang != "" {

			// Search for quality.
			l := strings.SplitN(lang, ";", 2)
			var quality float64
			var err error

			if len(l) == 2 {
				q := l[1]

				if strings.HasPrefix(q, "q=") {
					q = strings.SplitN(q, "=", 2)[1]
					quality, err = strconv.ParseFloat(q, 64)

					if err != nil {
						// Default value (1) if quality is empty.
						quality = 1
					}
				}
			}

			// Use order of items if no quality is given.
			if quality == 0 {
				quality = float64(length - i)
			}

			language := l[0]

			// If supported languages are given, return only the langs that fit.
			if len(supported_languages) != 0 {
				for _, supported_lang := range supported_languages {
					if language == supported_lang {
						langs = append(langs, Language{language, quality})
						break
					}
				}
			} else {
				// If no supported language is given, return all langs.
				langs = append(langs, Language{language, quality})
			}

			i++
		}
	}

	// Sort in reverse order (quality descending).
	quality := func(l1, l2 *Language) bool {
		return l1.quality > l2.quality
	}

	By(quality).Sort(langs)

	// Filter quality string.
	langString := []string{}
	for _, lang := range langs {
		langString = append(langString, lang.name)
	}

	return langString

}
