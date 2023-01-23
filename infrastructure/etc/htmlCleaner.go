package etc

import "regexp"

func HTMLCleaner(html string) string {
	re, _ := regexp.Compile(`<p><br></p>`)
	cleanedHTML := re.ReplaceAllString(html, "")
	return cleanedHTML
}
