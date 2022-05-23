package strcase

import (
	"strings"
	"testing"
)

var inflections = map[string]string{
	"star":        "stars",
	"STAR":        "STARS",
	"Star":        "Stars",
	"bus":         "buses",
	"fish":        "fish",
	"mouse":       "mice",
	"query":       "queries",
	"ability":     "abilities",
	"agency":      "agencies",
	"movie":       "movies",
	"archive":     "archives",
	"index":       "indices",
	"wife":        "wives",
	"safe":        "saves",
	"half":        "halves",
	"move":        "moves",
	"salesperson": "salespeople",
	"person":      "people",
	"spokesman":   "spokesmen",
	"man":         "men",
	"woman":       "women",
	"basis":       "bases",
	"diagnosis":   "diagnoses",
	"diagnosis_a": "diagnosis_as",
	"datum":       "data",
	"medium":      "media",
	"stadium":     "stadia",
	"analysis":    "analyses",
	"node_child":  "node_children",
	"child":       "children",
	"experience":  "experiences",
	"day":         "days",
	"comment":     "comments",
	"foobar":      "foobars",
	"newsletter":  "newsletters",
	"old_news":    "old_news",
	"news":        "news",
	"series":      "series",
	"species":     "species",
	"quiz":        "quizzes",
	"perspective": "perspectives",
	"ox":          "oxen",
	"photo":       "photos",
	"buffalo":     "buffaloes",
	"tomato":      "tomatoes",
	"dwarf":       "dwarves",
	"elf":         "elves",
	"information": "information",
	"equipment":   "equipment",
	"criterion":   "criteria",
	"foot":        "feet",
	"goose":       "geese",
	"moose":       "moose",
	"tooth":       "teeth",
	"milk":        "milk",
	"salt":        "salt",
	"time":        "time",
	"water":       "water",
	"paper":       "paper",
	"music":       "music",
	"help":        "help",
	"luck":        "luck",
	"oil":         "oil",
	"progress":    "progress",
	"rain":        "rain",
	"research":    "research",
	"shopping":    "shopping",
	"software":    "software",
	"traffic":     "traffic",
	"zombie":      "zombies",
	"campus":      "campuses",
	"harddrive":   "harddrives",
	"drive":       "drives",
	"sms":         "sms",
	"message":     "messages",
	"book":        "books",
}

func TestPlural(t *testing.T) {
	for key, value := range inflections {
		if v := ToPlural(strings.ToUpper(key)); v != strings.ToUpper(value) {
			t.Errorf("%v's plural should be %v, but got %v", strings.ToUpper(key), strings.ToUpper(value), v)
		}

		if v := ToPlural(strings.Title(key)); v != strings.Title(value) {
			t.Errorf("%v's plural should be %v, but got %v", strings.Title(key), strings.Title(value), v)
		}

		if v := ToPlural(key); v != value {
			t.Errorf("%v's plural should be %v, but got %v", key, value, v)
		}
	}
}

func TestSingular(t *testing.T) {
	for key, value := range inflections {
		if v := ToSingular(strings.ToUpper(value)); v != strings.ToUpper(key) {
			t.Errorf("%v's singular should be %v, but got %v", strings.ToUpper(value), strings.ToUpper(key), v)
		}

		if v := ToSingular(strings.Title(value)); v != strings.Title(key) {
			t.Errorf("%v's singular should be %v, but got %v", strings.Title(value), strings.Title(key), v)
		}

		if v := ToSingular(value); v != key {
			t.Errorf("%v's singular should be %v, but got %v", value, key, v)
		}
	}
}

func BenchmarkPlural(b *testing.B) {
	Init()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ToPlural("book")
	}
}

func BenchmarkSingular(b *testing.B) {
	Init()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ToSingular("books")
	}
}
