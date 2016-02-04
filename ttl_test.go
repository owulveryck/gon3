package gon3

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestTurtle(t *testing.T) {

	currentParserTests = []string{}
	currentParserTests = positiveParserTests

	verbosity := 2

	for _, testName := range currentParserTests {
		testFile := "./tests/turtle/" + testName
		b, err := ioutil.ReadFile(testFile)
		if err != nil {
			t.Fatalf("Error reading test file %s", testFile)
		}
		if verbosity > 0 {
			fmt.Printf("\nStarting test %s\n", testName)
		}
		p := NewParser(string(b))
		g, err := p.Parse()
		if err != nil {
			t.Fatalf("Test %s failed: %s", testName, err)
		}
		if verbosity > 0 {
			fmt.Printf("Test %s passed.\n", testName)
		}
		if verbosity > 1 {
			fmt.Printf("Graph:\n%s\n", g)
		}
	}
}

var currentParserTests []string
var positiveParserTests []string = []string{
	"turtle-syntax-base-01.ttl",
	"turtle-syntax-base-02.ttl",
	"turtle-syntax-base-03.ttl",
	"turtle-syntax-base-04.ttl",
	"turtle-syntax-blank-label.ttl",
	"turtle-syntax-bnode-01.ttl",
	"turtle-syntax-bnode-02.ttl",
	"turtle-syntax-bnode-03.ttl",
	"turtle-syntax-bnode-04.ttl",
	"turtle-syntax-bnode-05.ttl",
	"turtle-syntax-bnode-06.ttl",
	"turtle-syntax-bnode-07.ttl",
	"turtle-syntax-bnode-08.ttl",
	"turtle-syntax-bnode-09.ttl",
	"turtle-syntax-bnode-10.ttl",
	"turtle-syntax-datatypes-01.ttl",
	"turtle-syntax-datatypes-02.ttl",
	"turtle-syntax-file-01.ttl",
	"turtle-syntax-file-02.ttl",
	"turtle-syntax-file-03.ttl",
	"turtle-syntax-kw-01.ttl",
	"turtle-syntax-kw-02.ttl",
	"turtle-syntax-kw-03.ttl",
	"turtle-syntax-lists-01.ttl",
	"turtle-syntax-lists-02.ttl",
	"turtle-syntax-lists-03.ttl",
	"turtle-syntax-lists-04.ttl",
	"turtle-syntax-lists-05.ttl",
	"turtle-syntax-ln-colons.ttl",
	"turtle-syntax-ln-dots.ttl",
	"turtle-syntax-ns-dots.ttl",
	"turtle-syntax-number-01.ttl",
	"turtle-syntax-number-02.ttl",
	"turtle-syntax-number-03.ttl",
	"turtle-syntax-number-04.ttl",
	"turtle-syntax-number-05.ttl",
	"turtle-syntax-number-06.ttl",
	"turtle-syntax-number-07.ttl",
	"turtle-syntax-number-08.ttl",
	"turtle-syntax-number-09.ttl",
	"turtle-syntax-number-10.ttl",
	"turtle-syntax-number-11.ttl",
	"turtle-syntax-pname-esc-01.ttl",
	"turtle-syntax-pname-esc-02.ttl",
	"turtle-syntax-pname-esc-03.ttl",
	"turtle-syntax-prefix-01.ttl",
	// TODO: allow any combination of cases on SPARQL-style prefix
	//"turtle-syntax-prefix-02.ttl",
	"turtle-syntax-prefix-03.ttl",
	"turtle-syntax-prefix-04.ttl",
	"turtle-syntax-prefix-05.ttl",
	"turtle-syntax-prefix-06.ttl",
	"turtle-syntax-prefix-07.ttl",
	"turtle-syntax-prefix-08.ttl",
	"turtle-syntax-prefix-09.ttl",
	"turtle-syntax-str-esc-01.ttl",
	"turtle-syntax-str-esc-02.ttl",
	"turtle-syntax-str-esc-03.ttl",
	"turtle-syntax-string-01.ttl",
	"turtle-syntax-string-02.ttl",
	"turtle-syntax-string-03.ttl",
	"turtle-syntax-string-04.ttl",
	"turtle-syntax-string-05.ttl",
	"turtle-syntax-string-06.ttl",
	"turtle-syntax-string-07.ttl",
	"turtle-syntax-string-08.ttl",
	"turtle-syntax-string-09.ttl",
	"turtle-syntax-string-10.ttl",
	"turtle-syntax-string-11.ttl",
	"turtle-syntax-struct-01.ttl",
	"turtle-syntax-struct-02.ttl",
	"turtle-syntax-struct-03.ttl",
	"turtle-syntax-struct-04.ttl",
	"turtle-syntax-struct-05.ttl",
	"turtle-syntax-uri-01.ttl",
	"turtle-syntax-uri-02.ttl",
	"turtle-syntax-uri-03.ttl",
	"turtle-syntax-uri-04.ttl",
	"anonymous_blank_node_object.ttl",
	"anonymous_blank_node_subject.ttl",
	"bareword_a_predicate.ttl",
	"bareword_decimal.ttl",
	"bareword_double.ttl",
	"bareword_integer.ttl",
	"blankNodePropertyList_as_object.ttl",
	"blankNodePropertyList_as_subject.ttl",
	"blankNodePropertyList_containing_collection.ttl",
	"blankNodePropertyList_with_multiple_triples.ttl",
	"collection_object.ttl",
	"collection_subject.ttl",
	"comment_following_localName.ttl",
	"comment_following_PNAME_NS.ttl",
	"default_namespace_IRI.ttl",
	"double_lower_case_e.ttl",
	"empty_collection.ttl",
	"first.ttl",
	"HYPHEN_MINUS_in_localName.ttl",
	"IRIREF_datatype.ttl",
	"IRI_subject.ttl",
	"IRI_with_all_punctuation.ttl",
	"IRI_with_eight_digit_numeric_escape.ttl",
	"IRI_with_four_digit_numeric_escape.ttl",
	"labeled_blank_node_object.ttl",
	"labeled_blank_node_subject.ttl",
	"labeled_blank_node_with_leading_digit.ttl",
	"labeled_blank_node_with_leading_underscore.ttl",
	"labeled_blank_node_with_non_leading_extras.ttl",
	"labeled_blank_node_with_PN_CHARS_BASE_character_boundaries.ttl",
	"langtagged_LONG.ttl",
	"langtagged_LONG_with_subtag.ttl",
	"langtagged_non_LONG.ttl",
	"lantag_with_subtag.ttl",
	"last.ttl",
	"LITERAL1_all_controls.ttl",
	"LITERAL1_all_punctuation.ttl",
	"LITERAL1_ascii_boundaries.ttl",
	"LITERAL1.ttl",
	"LITERAL1_with_UTF8_boundaries.ttl",
	"LITERAL2_ascii_boundaries.ttl",
	"LITERAL2.ttl",
	"LITERAL2_with_UTF8_boundaries.ttl",
	"literal_false.ttl",
	"LITERAL_LONG1_ascii_boundaries.ttl",
	"LITERAL_LONG1.ttl",
	"LITERAL_LONG1_with_1_squote.ttl",
	"LITERAL_LONG1_with_2_squotes.ttl",
	"LITERAL_LONG1_with_UTF8_boundaries.ttl",
	"LITERAL_LONG2_ascii_boundaries.ttl",
	"LITERAL_LONG2.ttl",
	"LITERAL_LONG2_with_1_squote.ttl",
	"LITERAL_LONG2_with_2_squotes.ttl",
	"LITERAL_LONG2_with_REVERSE_SOLIDUS.ttl",
	"LITERAL_LONG2_with_UTF8_boundaries.ttl",
	"literal_true.ttl",
	"literal_with_BACKSPACE.ttl",
	"literal_with_CARRIAGE_RETURN.ttl",
	"literal_with_CHARACTER_TABULATION.nt",
	"literal_with_CHARACTER_TABULATION.ttl",
	"literal_with_escaped_BACKSPACE.ttl",
	"literal_with_escaped_CARRIAGE_RETURN.ttl",
	"literal_with_escaped_CHARACTER_TABULATION.ttl",
	"literal_with_escaped_FORM_FEED.ttl",
	"literal_with_escaped_LINE_FEED.ttl",
	"literal_with_FORM_FEED.ttl",
	"literal_with_LINE_FEED.ttl",
	"literal_with_numeric_escape4.ttl",
	"literal_with_numeric_escape8.ttl",
	"literal_with_REVERSE_SOLIDUS.ttl",
	"localName_with_assigned_nfc_bmp_PN_CHARS_BASE_character_boundaries.ttl",
	"localName_with_assigned_nfc_PN_CHARS_BASE_character_boundaries.ttl",
	"localname_with_COLON.ttl",
	"localName_with_leading_digit.ttl",
	"localName_with_leading_underscore.ttl",
	"localName_with_nfc_PN_CHARS_BASE_character_boundaries.ttl",
	"localName_with_non_leading_extras.ttl",
	// TODO: proper unescaping
	//"manifest.ttl",
	"negative_numeric.ttl",
	"nested_blankNodePropertyLists.ttl",
	"nested_collection.ttl",
	"number_sign_following_localName.ttl",
	"number_sign_following_PNAME_NS.ttl",
	"numeric_with_leading_0.ttl",
	"objectList_with_two_objects.ttl",
	"old_style_base.ttl",
	"old_style_prefix.ttl",
	"percent_escaped_localName.ttl",
	"positive_numeric.ttl",
	"predicateObjectList_with_two_objectLists.ttl",
	"prefixed_IRI_object.ttl",
	"prefixed_IRI_predicate.ttl",
	"prefixed_name_datatype.ttl",
	"prefix_only_IRI.ttl",
	"prefix_reassigned_and_used.ttl",
	"prefix_with_non_leading_extras.ttl",
	"prefix_with_PN_CHARS_BASE_character_boundaries.ttl",
	"repeated_semis_at_end.ttl",
	"repeated_semis_not_at_end.ttl",
	"reserved_escaped_localName.ttl",
	"sole_blankNodePropertyList.ttl",
	"SPARQL_style_base.ttl",
	"SPARQL_style_prefix.ttl",
	"turtle-eval-struct-01.ttl",
	"turtle-eval-struct-02.ttl",
	"turtle-subm-01.ttl",
	"turtle-subm-02.ttl",
	"turtle-subm-03.ttl",
	"turtle-subm-04.ttl",
	"turtle-subm-05.ttl",
	"turtle-subm-06.ttl",
	"turtle-subm-07.ttl",
	"turtle-subm-08.ttl",
	"turtle-subm-09.ttl",
	"turtle-subm-10.ttl",
	"turtle-subm-11.ttl",
	"turtle-subm-12.ttl",
	"turtle-subm-13.ttl",
	"turtle-subm-14.ttl",
	"turtle-subm-15.ttl",
	"turtle-subm-16.ttl",
	"turtle-subm-17.ttl",
	"turtle-subm-18.ttl",
	"turtle-subm-19.ttl",
	"turtle-subm-20.ttl",
	"turtle-subm-21.ttl",
	"turtle-subm-22.ttl",
	"turtle-subm-23.ttl",
	"turtle-subm-24.ttl",
	"turtle-subm-25.ttl",
	"turtle-subm-26.ttl",
	"turtle-subm-27.ttl",
	"two_LITERAL_LONG2s.ttl",
	"underscore_in_localName.ttl",
}
