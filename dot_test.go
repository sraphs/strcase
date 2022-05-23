package strcase

import (
	"testing"
)

func TestToDot(t *testing.T) {
	cases := [][]string{
		{"testCase", "test.case"},
		{"TestCase", "test.case"},
		{"Test Case", "test.case"},
		{" Test Case", "test.case"},
		{"Test Case ", "test.case"},
		{" Test Case ", "test.case"},
		{"test", "test"},
		{"test.case", "test.case"},
		{"Test", "test"},
		{"", ""},
		{"ManyManyWords", "many.many.words"},
		{"manyManyWords", "many.many.words"},
		{"AnyKind of.string", "any.kind.of.string"},
		{"numbers2and55with000", "numbers.2.and.55.with.000"},
		{"JSONData", "json.data"},
		{"userID", "user.id"},
		{"AAAbbb", "aa.abbb"},
		{"1A2", "1.a.2"},
		{"A1B", "a.1.b"},
		{"A1A2A3", "a.1.a.2.a.3"},
		{"A1 A2 A3", "a.1.a.2.a.3"},
		{"AB1AB2AB3", "ab.1.ab.2.ab.3"},
		{"AB1 AB2 AB3", "ab.1.ab.2.ab.3"},
		{"some string", "some.string"},
		{" some string", "some.string"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := ToDot(in)
		if result != out {
			t.Errorf("%q (%q != %q)", in, result, out)
		}
	}
}

func BenchmarkToDot(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ToDot("testCase")
	}
}

func TestToDotWithIgnore(t *testing.T) {
	cases := [][]string{
		{"testCase", "test.case"},
		{"TestCase", "test.case"},
		{"Test Case", "test.case"},
		{" Test Case", "test.case"},
		{"Test Case ", "test.case"},
		{" Test Case ", "test.case"},
		{"test", "test"},
		{"test.case", "test.case"},
		{"Test", "test"},
		{"", ""},
		{"ManyManyWords", "many.many.words"},
		{"manyManyWords", "many.many.words"},
		{"AnyKind of.string", "any.kind.of.string"},
		{"numbers2and55with000", "numbers.2.and.55.with.000"},
		{"JSONData", "json.data"},
		{"AwesomeActivity.UserID", "awesome.activity.user.id", "."},
		{"AwesomeActivity.User.Id", "awesome.activity.user.id", "."},
		{"AwesomeUsername@Awesome.Com", "awesome.username@awesome.com", ".@"},
		{"lets.ignore all.of dots.and.dashes", "lets.ignore.all.of.dots.and.dashes", ".."},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		var ignore string
		ignore = ""
		if len(i) == 3 {
			ignore = i[2]
		}
		result := ToDotWithIgnore(in, ignore)
		if result != out {
			istr := ""
			if len(i) == 3 {
				istr = " ignoring '" + i[2] + "'"
			}
			t.Errorf("%q (%q != %q%s)", in, result, out, istr)
		}
	}
}

func BenchmarkToDotWithIgnore(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ToDotWithIgnore("awesome.activity.user.id", ".")
	}
}
