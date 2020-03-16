package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test_getRequestUrl_fails_with_wrong_url(t *testing.T) {
	backupUrl := catalogBaseUrl
	defer func() {
		catalogBaseUrl = backupUrl
	}()

	catalogBaseUrl = ":no scheme"

	_, err := getCatalogPageUrl(0)
	if err == nil {
		t.Errorf("getCatalogPageUrl() error = %v, wantErr %v", err, true)
	}
}

func Test_getRequestUrl(t *testing.T) {
	//https://www?sort=leaders&offset=0&limit=250&type=share

	type args struct {
		page int
	}
	type want struct {
		startsWith string
		contains   string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "starts with Parser.Catalog.BaseUrl",
			want: want{
				startsWith: getProperties().Parser.Catalog.BaseUrl,
			},
		},
		{
			name: "contains limit Parser.Catalog.PageSize",
			want: want{
				startsWith: getProperties().Parser.Catalog.BaseUrl,
				contains:   fmt.Sprintf("limit=%d", getProperties().Parser.Catalog.PageSize),
			},
		},
		{
			name: "contains offset 0 on first page",
			want: want{
				startsWith: getProperties().Parser.Catalog.BaseUrl,
				contains:   "offset=0",
			},
		},
		{
			name: "contains offset Parser.Catalog.PageSize * 3 on 4th page",
			args: args{3},
			want: want{
				startsWith: getProperties().Parser.Catalog.BaseUrl,
				contains:   fmt.Sprintf("offset=%d", getProperties().Parser.Catalog.PageSize*3),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getCatalogPageUrl(tt.args.page)
			if err != nil {
				t.Errorf("getCatalogPageUrl() error = %v", err)
				return
			}
			if !strings.HasPrefix(got, tt.want.startsWith) {
				t.Errorf("getCatalogPageUrl() got = %v, want starts with %v", got, tt.want)
				return
			}
			if !strings.Contains(got, tt.want.contains) {
				t.Errorf("getCatalogPageUrl() got = %v, want contains %v", got, tt.want)
				return
			}
		})
	}
}
