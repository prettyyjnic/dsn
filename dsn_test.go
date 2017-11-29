// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package dsn

import (
	"reflect"
	"testing"
)

func TestParseDSN(t *testing.T) {
	type args struct {
		dsn string
	}
	tests := []struct {
		name    string
		args    args
		wantCfg *DSN
		wantErr bool
	}{
		{
			name: "test1",
			args: args{"root:123456@tcp(127.0.0.1:3306)/Test?charset=utf8"},
			wantCfg: &DSN{User: "root", Passwd: "123456", Net: "tcp", Addr: "127.0.0.1:3306", DBName: "Test", Params: map[string]string{
				"charset": "utf8",
			}},
		},
		{
			name: "test2",
			args: args{"root:123456@/Test?charset=utf8"},
			wantCfg: &DSN{User: "root", Passwd: "123456", Net: "", Addr: "", DBName: "Test", Params: map[string]string{
				"charset": "utf8",
			}},
		},
		{
			name:    "test2",
			args:    args{"root@/Test"},
			wantCfg: &DSN{User: "root", Passwd: "", Net: "", Addr: "", DBName: "Test", Params: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCfg, err := ParseDSN(tt.args.dsn)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDSN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCfg, tt.wantCfg) {
				t.Errorf("ParseDSN() = %v, want %v", gotCfg, tt.wantCfg)
			}
		})
	}
}

func TestDSN_FormatDSN(t *testing.T) {
	type fields struct {
		User   string
		Passwd string
		Net    string
		Addr   string
		DBName string
		Params map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
	 {
	 	name: "test formatDSN",
	 	fields: fields{
	 		User:"root",
			Passwd:"123456",
			Net:"tcp",
			Addr:"localhost:3306",
			DBName:"Test",
			Params: map[string]string{
				"charset": "utf8",
			},
		},
		want: "root:123456@tcp(localhost:3306)/Test?charset=utf8",
	 },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &DSN{
				User:   tt.fields.User,
				Passwd: tt.fields.Passwd,
				Net:    tt.fields.Net,
				Addr:   tt.fields.Addr,
				DBName: tt.fields.DBName,
				Params: tt.fields.Params,
			}
			if got := cfg.FormatDSN(); got != tt.want {
				t.Errorf("DSN.FormatDSN() = %v, want %v", got, tt.want)
			}
		})
	}
}
