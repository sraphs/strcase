package strcase

import (
	"testing"
)

func TestToPascal(t *testing.T) {
	cases := [][]string{
		{"test_case", "TestCase"},
		{"test.case", "TestCase"},
		{"test", "Test"},
		{"TestCase", "TestCase"},
		{" test  case ", "TestCase"},
		{"", ""},
		{"many_many_words", "ManyManyWords"},
		{"AnyKind of_string", "AnyKindOfString"},
		{"odd-fix", "OddFix"},
		{"numbers2And55with000", "Numbers2And55With000"},
		{"ID", "Id"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := ToPascal(in)
		if result != out {
			t.Errorf("%q (%q != %q)", in, result, out)
		}
	}
}

func BenchmarkToPascal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToPascal("test_case")
	}
}

func TestCustomAcronymsToPascal(t *testing.T) {
	tests := []struct {
		name         string
		acronymKey   string
		acronymValue string
		expected     string
	}{
		{
			name:         "API Custom Acronym",
			acronymKey:   "API",
			acronymValue: "api",
			expected:     "Api",
		},
		{
			name:         "ABCDACME Custom Acroynm",
			acronymKey:   "ABCDACME",
			acronymValue: "AbcdAcme",
			expected:     "AbcdAcme",
		},
		{
			name:         "PostgreSQL Custom Acronym",
			acronymKey:   "PostgreSQL",
			acronymValue: "PostgreSQL",
			expected:     "PostgreSQL",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ConfigureAcronym(test.acronymKey, test.acronymValue)
			if result := ToPascal(test.acronymKey); result != test.expected {
				t.Errorf("expected custom acronym result %s, got %s", test.expected, result)
			}
		})
	}
}

func BenchmarkCustomAcronymsToPascal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConfigureAcronym("API", "api")
		ToPascal("API")
	}
}
