package main

import (
	"reflect"
	"testing"
)

func TestFindHousingColumns(t *testing.T) {
	headers := []string{
		"median_house_value",
		"median_income",
		"housing_median_age",
		"total_rooms",
		"total_bedrooms",
		"population",
		"households",
	}

	got, err := findHousingColumns(headers)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := map[string]int{
		"value":    0,
		"income":   1,
		"age":      2,
		"rooms":    3,
		"bedrooms": 4,
		"pop":      5,
		"hh":       6,
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestFindHousingColumnsMissingField(t *testing.T) {
	_, err := findHousingColumns([]string{"value", "income"})
	if err == nil {
		t.Fatal("expected an error for missing columns")
	}
}

func TestParseHousingRecord(t *testing.T) {
	columns := map[string]int{
		"value": 0, "income": 1, "age": 2,
		"rooms": 3, "bedrooms": 4, "pop": 5, "hh": 6,
	}

	row := []string{"452600", "8.3252", "41", "880", "129", "322", "126"}

	got, err := parseHousingRecord(row, columns)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got.Value != 452600 || got.Income != 8.3252 ||
		got.Age != 41 || got.Rooms != 880 ||
		got.Bedrooms != 129 || got.Pop != 322 || got.HH != 126 {
		t.Fatalf("unexpected record: %+v", got)
	}
}

func TestParseHousingRecordInvalidNumber(t *testing.T) {
	columns := map[string]int{
		"value": 0, "income": 1, "age": 2,
		"rooms": 3, "bedrooms": 4, "pop": 5, "hh": 6,
	}

	row := []string{"not-a-number", "8.3252", "41", "880", "129", "322", "126"}

	if _, err := parseHousingRecord(row, columns); err == nil {
		t.Fatal("expected an error for invalid numeric input")
	}
}
