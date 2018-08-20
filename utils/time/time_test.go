package time

import (
	"testing"
	)

func TestTime_UnmarshalJSON(t *testing.T) {
	mytime := Now()

	bytesData, _ := mytime.MarshalJSON()

	var mytime2 Time

	if err := mytime2.UnmarshalJSON(bytesData); err != nil {
		t.Fatal(err)
	}

	if mytime2.String() != mytime.String() {
		t.Error("formatted time string should be equal")
	}
}

func TestTime_MarshalJSON(t *testing.T) {
	mytime := Now()

	bytesData, err := mytime.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(bytesData))
}

func TestDate_MarshalJSON(t *testing.T) {
	mydate := Today()

	bytesData, err := mydate.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(bytesData))
}

func TestDate_UnmarshalJSON(t *testing.T) {
	mydate := Today()

	bytesData, err := mydate.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}

	var mydate2 Date

	if err := mydate2.UnmarshalJSON(bytesData); err != nil {
		t.Fatal(err)
	}

	if mydate.String() != mydate2.String() {
		t.Error("formatted time string should be equal")
	}
}

func TestParseDate(t *testing.T) {
	d, err := ParseDate("2018-08-20")
	if err != nil {
		t.Fatal(err)
	}
	if d.String() != "2018-08-20" {
		t.Error("should be 2018-08-20")
	}
}

func TestParseTime(t *testing.T) {
	d, err := ParseTime("2018-08-20 12:32:01")
	if err != nil {
		t.Fatal(err)
	}
	if d.String() != "2018-08-20 12:32:01" {
		t.Error("should be 2018-08-20 12:32:01")
	}
}