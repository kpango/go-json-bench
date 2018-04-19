package main

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"testing"
)

type JSON struct {
	Comment  string `json:"comment"`
	Name     string `json:"name"`
	Optional bool   `json:"optional"`
	Type     string `json:"type"`
}

var (
	j = `{
   "name":"id",
   "type":"UUID",
   "optional":true,
   "comment":"unique identifier of the domain. generated on create, never reused"
}`
)

func unmarshal() error {
	var target JSON
	body, err := ioutil.ReadAll(strings.NewReader(j))
	if err != nil {
		return err
	}
	return json.Unmarshal(body, &target)
}

func decode() error {
	var target JSON
	return json.NewDecoder(strings.NewReader(j)).Decode(&target)
}

func BenchmarkJSONUnmarshal(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		err := unmarshal()
		if err != nil {
			b.Error(err)
		}
	}
	b.StopTimer()
}

func BenchmarkJSONDecode(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		err := decode()
		if err != nil {
			b.Error(err)
		}
	}
	b.StopTimer()
}

func BenchmarkJSONUnmarshalParallel(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err := unmarshal()
			if err != nil {
				b.Error(err)
			}

		}
	})
	b.StopTimer()
}

func BenchmarkJSONDecodeParallel(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err := decode()
			if err != nil {
				b.Error(err)
			}
		}
	})
	b.StopTimer()
}
