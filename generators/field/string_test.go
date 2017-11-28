package field

import (
	"datagenerator/config"
	"regexp"
	"testing"
)

func testRegex(t *testing.T, regex string) {
	g, err := NewStringGenerator(&config.Column{
		Name:   "test_column",
		Type:   "string",
		Regexp: &regex})

	if regex == "" {
		if err == nil {
			t.Error("Expected error for empty regex")
		}
		return
	}

	_, err2 := regexp.Compile(regex)
	if err2 != nil {
		if err == nil {
			t.Error("Expected error for invalid regex")
		}
		return
	}

	for i := 0; i < 100; i++ {
		v := g.Generate()
		sv, ok := v.(string)
		if !ok {
			t.Errorf("Expected to return string but got %T", sv)
			return
		}
		if ok, _ := regexp.Match(regex, []byte(sv)); !ok {
			t.Errorf("Expected value %v to be withtin regular expression pattern %v", sv, regex)
			return
		}
	}
}

func TestStringGenerator_Generate(t *testing.T) {
	testRegex(t, "")
	testRegex(t, "[A-Z]{3}")
	testRegex(t, "[A-Z]{3}[0-9]{10}")
	testRegex(t, "^[a-z]{5,10}@[a-z]{5,10}\\.(com|net|org)$")
	testRegex(t, "^[a-z")
}
