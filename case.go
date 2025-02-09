package str

import (
	"strings"
	"unicode"

	"github.com/go-openapi/inflect"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	casingRules = setCasingRules()
	acronyms    = make(map[string]struct{})
)

func setCasingRules() *inflect.Ruleset {
	rules := inflect.NewDefaultRuleset()
	for _, w := range []string{"ACL", "API", "ASCII", "AWS", "CPU", "CSS", "DNS", "EOF", "GB", "GUID", "HCL", "HTML", "HTTP", "HTTPS", "ID", "IP", "JSON", "KB", "LHS", "MAC", "MB", "QPS", "RAM", "RHS", "RPC", "SLA", "SMTP", "SQL", "SSH", "SSO", "TCP", "TLS", "TTL", "UDP", "UI", "UID", "URI", "URL", "UTF8", "UUID", "VM", "XML", "XMPP", "XSRF", "XSS"} {
		acronyms[w] = struct{}{}
		rules.AddAcronym(w)
	}
	return rules
}

func goPascalWords(words []string) string {
	for i, w := range words {
		upper := strings.ToUpper(w)
		if _, ok := acronyms[upper]; ok {
			words[i] = upper
		} else {
			words[i] = casingRules.Capitalize(w)
		}
	}
	return strings.Join(words, "")
}

func pascalWords(words []string) string {
	caser := cases.Title(language.English)
	for i, w := range words {
		words[i] = caser.String(strings.ToLower(w))
	}
	return strings.Join(words, "")
}

func splitWords(s string) []string {
	var words []string
	var lastIndex, lastUpper int
	for i, r := range s {
		if unicode.IsPunct(r) || unicode.IsSymbol(r) || unicode.IsSpace(r) {
			if w := s[lastIndex:i]; w != "" {
				words = append(words, w)
			}
			lastIndex = i + 1
			continue
		}

		if unicode.IsUpper(r) {
			seq := i-1 == lastUpper
			lastUpper = i
			if seq {
				continue
			}

			if w := s[lastIndex:i]; w != "" {
				words = append(words, w)
			}
			lastIndex = i
		}
	}
	if w := s[lastIndex:]; w != "" {
		words = append(words, w)
	}

	return words
}
