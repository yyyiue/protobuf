package retag

import "testing"

func Test_getFieldAndTag(t *testing.T) {
	type args struct {
		line    string
		msgName string
	}

	type test struct {
		name      string
		args      args
		wantField string
		wantTag   string
	}

	var tests = make([]test, 0)
	tests = append(tests, test{
		name: "case1",
		args: args{
			line:    "  int64 number_aa = 1;   //`json:\"number123,string\"`",
			msgName: "ChildOrder",
		},
		wantField: "ChildOrder.NumberAa",
		wantTag:   "",
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotField, _ := getFieldAndTag(tt.args.line, tt.args.msgName)
			if gotField != tt.wantField {
				t.Errorf("getFieldTag() gotField = %v, want %v", gotField, tt.wantField)
			}
			//if gotTag != tt.wantTag {
			//	t.Errorf("getFieldTag() gotTag = %v, want %v", gotTag, tt.wantTag)
			//}
		})
	}
}
