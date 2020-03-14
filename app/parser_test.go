package main

/*
func TestParse(t *testing.T) {
	type args struct {
		url    string
		reader io.Reader
		chData chan stockTicker
		chErr  chan error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_closeReader(t *testing.T) {
	type args struct {
		response *http.Response
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := closeReader(tt.args.response); (err != nil) != tt.wantErr {
				t.Errorf("closeReader() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_containsVerdana(t *testing.T) {
	type args struct {
		in0       int
		selection *goquery.Selection
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsVerdana(tt.args.in0, tt.args.selection); got != tt.want {
				t.Errorf("containsVerdana() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertIfNeeded(t *testing.T) {
	type args struct {
		response *http.Response
	}
	tests := []struct {
		name    string
		args    args
		want    io.Reader
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertIfNeeded(tt.args.response)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertIfNeeded() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertIfNeeded() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getResponse(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getResponse(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("getResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getResponse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lastLevelTable(t *testing.T) {
	type args struct {
		in0       int
		selection *goquery.Selection
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lastLevelTable(tt.args.in0, tt.args.selection); got != tt.want {
				t.Errorf("lastLevelTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseOnline(t *testing.T) {
	tests := []struct {
		name  string
		want  []stockTicker
		want1 []error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseOnline()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseOnline() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parseOnline() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_parseOnlinePage(t *testing.T) {
	type args struct {
		url       string
		chData    chan stockTicker
		chErr     chan error
		chCounter chan int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
*/
