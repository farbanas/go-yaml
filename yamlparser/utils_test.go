package yamlparser

import (
	"reflect"
	"testing"
)

func TestGetValue(t *testing.T) {
	data := OpenFileRead("test.yaml")
	yamlData := ReadYaml(data)

	type args struct {
		yamlData interface{}
		query    string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"Get root value 1", args{yamlData: yamlData, query: "worry"}, "deer"},
		{"Get root value 2", args{yamlData: yamlData, query: "satellites"}, "record"},
		{"Get array value 1", args{yamlData: yamlData, query: "potatoes.0.soil.through"}, []interface{}{true, false, "station"}},
		{"Get array value 2", args{yamlData: yamlData, query: "potatoes.1"}, false},
		{"Get array value 3", args{yamlData: yamlData, query: "potatoes.2"}, true},
		{"Get array value 4", args{yamlData: yamlData, query: "potatoes.0.soil.through.2"}, "station"},
		{"Get leaf value 1", args{yamlData: yamlData, query: "potatoes.0.soil.engineer"}, true},
		{"Get leaf value 2", args{yamlData: yamlData, query: "potatoes.0.soil.tip"}, "mind"},
		{"Get leaf value 3", args{yamlData: yamlData, query: "potatoes.0.percent"}, true},
		{"Get leaf value 3", args{yamlData: yamlData, query: "potatoes.0.stiff"}, "tight"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetValue(tt.args.yamlData, tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Logf("Got type: %T, want type %T", got, tt.want)
				t.Errorf("GetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetValue(t *testing.T) {
	data := OpenFileRead("test.yaml")
	yamlData := ReadYaml(data)

	type args struct {
		yamlData interface{}
		query    string
		val      string
		path     string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"Set root value 1", args{yamlData: yamlData, query: "worry", val: "DarthVader", path: "test.yaml"}, "DarthVader"},
		{"Set root value 2", args{yamlData: yamlData, query: "satellites", val: "DarthVader", path: "test.yaml"}, "DarthVader"},
		{"Set array value 1", args{yamlData: yamlData, query: "potatoes.0.soil.through", val: "DarthVader", path: "test.yaml"}, "DarthVader"},
		{"Set array value 2", args{yamlData: yamlData, query: "potatoes.1", val: "DarthVader", path: "test.yaml"}, "DarthVader"},
		{"Set array value 3", args{yamlData: yamlData, query: "potatoes.2", val: "DarthVader", path: "test.yaml"}, "DarthVader"},
		{"Set array value 4", args{yamlData: yamlData, query: "potatoes.0.soil.through.2", val: "DarthVader", path: "test.yaml"}, "DarthVader"},
		{"Set leaf value 1", args{yamlData: yamlData, query: "potatoes.0.soil.engineer", val: "DarthVader", path: "test.yaml"}, "DarthVader"},
		{"Set leaf value 2", args{yamlData: yamlData, query: "potatoes.0.soil.tip", val: "DarthVader", path: "test.yaml"}, "DarthVader"},
		{"Set leaf value 3", args{yamlData: yamlData, query: "potatoes.0.percent", val: "DarthVader", path: "test.yaml"}, "DarthVader"},
		{"Set leaf value 3", args{yamlData: yamlData, query: "potatoes.0.stiff", val: "DarthVader", path: "test.yaml"}, "DarthVader"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldVal := GetValue(tt.args.yamlData, tt.args.query)
			SetValue(tt.args.yamlData, tt.args.query, tt.args.val, tt.args.path)
			if got := GetValue(tt.args.yamlData, tt.args.query); !reflect.DeepEqual(got, tt.want) {
				SetValue(tt.args.yamlData, tt.args.query, oldVal, tt.args.path)
				t.Logf("Got type: %T, want type %T", got, tt.want)
				t.Errorf("GetValue() = %v, want %v", got, tt.want)
			}
			SetValue(tt.args.yamlData, tt.args.query, oldVal, tt.args.path)
		})
	}
}

func BenchmarkGetValue(b *testing.B) {
	data := OpenFileRead("test.yaml")
	yamlData := ReadYaml(data)

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GetValue(yamlData, "potatoes.0.soil.through.2")
	}
}

func BenchmarkGetValueReflect(b *testing.B) {
	data := OpenFileRead("test.yaml")
	yamlData := ReadYaml(data)

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GetValueReflect(yamlData, "potatoes.0.soil.through.2")
	}
}

func BenchmarkSetValue(b *testing.B) {
	data := OpenFileRead("test.yaml")
	yamlData := ReadYaml(data)

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		SetValue(yamlData, "potatoes.0.soil.through.2", "station", "test.yaml")
	}
}

func BenchmarkSetValueReflect(b *testing.B) {
	data := OpenFileRead("test.yaml")
	yamlData := ReadYaml(data)

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		SetValueReflect(yamlData, "potatoes.0.soil.through.2", "station", "test.yaml")
	}
}
