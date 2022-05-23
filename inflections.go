package strcase

import (
	"regexp"
	"strings"
	"sync"
)

type inflection struct {
	regexp  *regexp.Regexp
	replace string
}

// Regular is a regexp find replace inflection
type Regular struct {
	find    string
	replace string
}

// Irregular is a hard replace inflection,
// containing both singular and plural forms
type Irregular struct {
	singular string
	plural   string
}

// RegularSlice is a slice of Regular inflections
type RegularSlice []Regular

// IrregularSlice is a slice of Irregular inflections
type IrregularSlice []Irregular

var pluralInflections = RegularSlice{
	{"([a-z])$", "${1}s"},
	{"s$", "s"},
	{"^(ax|test)is$", "${1}es"},
	{"(octop|vir)us$", "${1}i"},
	{"(octop|vir)i$", "${1}i"},
	{"(alias|status|campus)$", "${1}es"},
	{"(bu)s$", "${1}ses"},
	{"(buffal|tomat)o$", "${1}oes"},
	{"([ti])um$", "${1}a"},
	{"([ti])a$", "${1}a"},
	{"sis$", "ses"},
	{"(?:([^f])fe|([lr])f)$", "${1}${2}ves"},
	{"(hive)$", "${1}s"},
	{"([^aeiouy]|qu)y$", "${1}ies"},
	{"(x|ch|ss|sh)$", "${1}es"},
	{"(matr|vert|ind)(?:ix|ex)$", "${1}ices"},
	{"^(m|l)ouse$", "${1}ice"},
	{"^(m|l)ice$", "${1}ice"},
	{"^(ox)$", "${1}en"},
	{"^(oxen)$", "${1}"},
	{"(quiz)$", "${1}zes"},
	{"(drive)$", "${1}s"},
}

var singularInflections = RegularSlice{
	{"s$", ""},
	{"(ss)$", "${1}"},
	{"(n)ews$", "${1}ews"},
	{"([ti])a$", "${1}um"},
	{"((a)naly|(b)a|(d)iagno|(p)arenthe|(p)rogno|(s)ynop|(t)he)(sis|ses)$", "${1}sis"},
	{"(^analy)(sis|ses)$", "${1}sis"},
	{"([^f])ves$", "${1}fe"},
	{"(hive)s$", "${1}"},
	{"(tive)s$", "${1}"},
	{"([lr])ves$", "${1}f"},
	{"([^aeiouy]|qu)ies$", "${1}y"},
	{"(s)eries$", "${1}eries"},
	{"(m)ovies$", "${1}ovie"},
	{"(c)ookies$", "${1}ookie"},
	{"(x|ch|ss|sh)es$", "${1}"},
	{"^(m|l)ice$", "${1}ouse"},
	{"(bus|campus)(es)?$", "${1}"},
	{"(o)es$", "${1}"},
	{"(shoe)s$", "${1}"},
	{"(cris|test)(is|es)$", "${1}is"},
	{"^(a)x[ie]s$", "${1}xis"},
	{"(octop|vir)(us|i)$", "${1}us"},
	{"(alias|status)(es)?$", "${1}"},
	{"^(ox)en", "${1}"},
	{"(vert|ind)ices$", "${1}ex"},
	{"(matr)ices$", "${1}ix"},
	{"(quiz)zes$", "${1}"},
	{"(database)s$", "${1}"},
	{"(drive)s$", "${1}"},
}

var irregularInflections = IrregularSlice{
	{"person", "people"},
	{"man", "men"},
	{"child", "children"},
	{"sex", "sexes"},
	{"move", "moves"},
	{"ombie", "ombies"},
	{"goose", "geese"},
	{"foot", "feet"},
	{"moose", "moose"},
	{"tooth", "teeth"},
	{"criterion", "criteria"},
}

var uncountableInflections = []string{
	"equipment",
	"information",
	"rice",
	"money",
	"species",
	"series",
	"fish",
	"sheep",
	"jeans",
	"police",
	"milk",
	"salt",
	"time",
	"water",
	"paper",
	"food",
	"art",
	"cash",
	"music",
	"help",
	"luck",
	"oil",
	"progress",
	"rain",
	"research",
	"shopping",
	"software",
	"traffic",
	"sms",
}

var compiledPluralMaps []inflection
var compiledSingularMaps []inflection

func compile() {
	compiledPluralMaps, compiledSingularMaps = nil, nil
	for _, uncountable := range uncountableInflections {
		inf := inflection{
			regexp:  regexp.MustCompile("^(?i)(" + uncountable + ")$"),
			replace: "${1}",
		}
		compiledPluralMaps = append(compiledPluralMaps, inf)
		compiledSingularMaps = append(compiledSingularMaps, inf)
	}

	for _, value := range irregularInflections {
		infs := []inflection{
			{regexp: regexp.MustCompile(strings.ToUpper(value.singular) + "$"), replace: strings.ToUpper(value.plural)},
			{regexp: regexp.MustCompile(strings.Title(value.singular) + "$"), replace: strings.Title(value.plural)},
			{regexp: regexp.MustCompile(value.singular + "$"), replace: value.plural},
		}
		compiledPluralMaps = append(compiledPluralMaps, infs...)
	}

	for _, value := range irregularInflections {
		infs := []inflection{
			{regexp: regexp.MustCompile(strings.ToUpper(value.plural) + "$"), replace: strings.ToUpper(value.singular)},
			{regexp: regexp.MustCompile(strings.Title(value.plural) + "$"), replace: strings.Title(value.singular)},
			{regexp: regexp.MustCompile(value.plural + "$"), replace: value.singular},
		}
		compiledSingularMaps = append(compiledSingularMaps, infs...)
	}

	for i := len(pluralInflections) - 1; i >= 0; i-- {
		value := pluralInflections[i]
		infs := []inflection{
			{regexp: regexp.MustCompile(strings.ToUpper(value.find)), replace: strings.ToUpper(value.replace)},
			{regexp: regexp.MustCompile(value.find), replace: value.replace},
			{regexp: regexp.MustCompile("(?i)" + value.find), replace: value.replace},
		}
		compiledPluralMaps = append(compiledPluralMaps, infs...)
	}

	for i := len(singularInflections) - 1; i >= 0; i-- {
		value := singularInflections[i]
		infs := []inflection{
			{regexp: regexp.MustCompile(strings.ToUpper(value.find)), replace: strings.ToUpper(value.replace)},
			{regexp: regexp.MustCompile(value.find), replace: value.replace},
			{regexp: regexp.MustCompile("(?i)" + value.find), replace: value.replace},
		}
		compiledSingularMaps = append(compiledSingularMaps, infs...)
	}
}

var InflectionOnce sync.Once

func Init() {
	InflectionOnce.Do(func() {
		compile()
	})
}

// AddIrregular adds an irregular inflection
func AddIrregular(singular, plural string) {
	irregularInflections = append(irregularInflections, Irregular{singular, plural})
	compile()
}

// AddUncountable adds an uncountable inflection
func AddUncountable(values ...string) {
	uncountableInflections = append(uncountableInflections, values...)
	compile()
}

// ToPlural converts a word to its plural form
func ToPlural(str string) string {
	Init()
	for _, inflection := range compiledPluralMaps {
		if inflection.regexp.MatchString(str) {
			return inflection.regexp.ReplaceAllString(str, inflection.replace)
		}
	}
	return str
}

// ToSingular converts a word to its singular form
func ToSingular(str string) string {
	Init()
	for _, inflection := range compiledSingularMaps {
		if inflection.regexp.MatchString(str) {
			return inflection.regexp.ReplaceAllString(str, inflection.replace)
		}
	}
	return str
}
