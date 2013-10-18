#!/usr/bin/python

# Copyright 2014 struktur AG. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

def parseAcceptLanguage(languages, supported_languages=None):

    langs = []
    browser_pref_langs = languages.split(',')

    i = 0
    length = len(browser_pref_langs)

    # parse quality strings and build a tuple
    # like ((float(quality), lang), (float(quality), lang))
    # which is sorted afterwards
    # if no quality string is given then the list order
    # is used as quality indicator
    for lang in browser_pref_langs:
        lang = lang.strip().lower().replace('_','-')
        if lang:
            l = lang.split(';', 2)
            quality = []

            if len(l) == 2:
                try:
                    q = l[1]
                    if q.startswith('q='):
                        q = q.split('=', 2)[1]
                        quality = float(q)
                except: pass

            if quality == []:
                quality = float(length-i)

            language = l[0]
            if supported_languages:
                if language in supported_languages:
                    # if allowed the add language
                    langs.append((quality, language))
            else:
                langs.append((quality, language))
            i = i + 1

    # sort and reverse it
    langs.sort()
    langs.reverse()

    # filter quality string
    langs = map(lambda x: x[1], langs)

    return langs

if __name__ == "__main__":

    actual = parseAcceptLanguage("en-US,en;q=0.8,de;q=0.6")
    print actual



