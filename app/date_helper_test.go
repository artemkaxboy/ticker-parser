package main

import (
	"reflect"
	"testing"
	"time"
)

func Test_replaceMonth(t *testing.T) {
	type args struct {
		dateString string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty string returns empty string",
			args: args{
				dateString: "",
			},
			want: "",
		},
		{
			name: "jan detected",
			args: args{
				dateString: "30 янв, 12:27",
			},
			want: "30 Jan, 12:27",
		},
		{
			name: "feb detected",
			args: args{
				dateString: "04 фев 2019, 12:02",
			},
			want: "04 Feb 2019, 12:02",
		},
		{
			name: "mar detected",
			args: args{
				dateString: "04 мар 2019, 12:02",
			},
			want: "04 Mar 2019, 12:02",
		},
		{
			name: "apr detected",
			args: args{
				dateString: "04 апр 2019, 12:02",
			},
			want: "04 Apr 2019, 12:02",
		},
		{
			name: "may detected",
			args: args{
				dateString: "04 май 2019, 12:02",
			},
			want: "04 May 2019, 12:02",
		},
		{
			name: "jun detected",
			args: args{
				dateString: "04 ИЮН 2019, 12:02",
			},
			want: "04 Jun 2019, 12:02",
		},
		{
			name: "jul detected",
			args: args{
				dateString: "04 Июл 2019, 12:02",
			},
			want: "04 Jul 2019, 12:02",
		},
		{
			name: "aug detected",
			args: args{
				dateString: "04 авГ 2019, 12:02",
			},
			want: "04 Aug 2019, 12:02",
		},
		{
			name: "sep detected",
			args: args{
				dateString: "04 сен 2019, 12:02",
			},
			want: "04 Sep 2019, 12:02",
		},
		{
			name: "oct detected",
			args: args{
				dateString: "04 окт 2019, 12:02",
			},
			want: "04 Oct 2019, 12:02",
		},
		{
			name: "nov detected",
			args: args{
				dateString: "04 ноя 2019, 12:02",
			},
			want: "04 Nov 2019, 12:02",
		},
		{
			name: "dec detected",
			args: args{
				dateString: "04 дек 2019, 12:02",
			},
			want: "04 Dec 2019, 12:02",
		},
		{
			name: "doesn't change anything else",
			args: args{
				dateString: "лдфорыпдркянвуфптлофумнглиро",
			},
			want: "лдфорыпдркJanуфптлофумнглиро",
		},
		{
			name: "doesn't change anything else",
			args: args{
				dateString: "лдфорыпдркянуфптлофумнглиро",
			},
			want: "лдфорыпдркянуфптлофумнглиро",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceMonth(tt.args.dateString); got != tt.want {
				t.Errorf("replaceMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseDate(t *testing.T) {
	type args struct {
		dateString string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name: "this year detects",
			args: args{
				dateString: "30 янв, 12:27",
			},
			want: time.Date(2020, 01, 30, 12, 27, 00, 0, time.UTC),
		},
		{
			name: "last year detects",
			args: args{
				dateString: "04 фев 2019, 12:02",
			},
			want: time.Date(2019, 2, 4, 12, 2, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseTime(tt.args.dateString)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseTime() got = %v, want %v", got, tt.want)
			}
		})
	}
}
