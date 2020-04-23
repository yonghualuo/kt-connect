package util

import "testing"

func TestWritePidFile(t *testing.T) {
	type args struct {
		pidFile string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "should_write_file_with_zh_cn_path",
			args: args{
				pidFile: "/tmp/中文.pid",
			},
			wantErr: false,
		},
		{
			name: "should_write_file_with_en_cn_path",
			args: args{
				pidFile: "/tmp/test.pid",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := WritePidFile(tt.args.pidFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("WritePidFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
