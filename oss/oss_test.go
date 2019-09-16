package oss

import "testing"

func TestPlugin_Exec(t *testing.T) {
	type fields struct {
		Config   Config
		DistList []string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:"test",
			fields:fields{
				Config:Config{
					Dist:"dist",
					Path:"xxx/test",
					EndPoint:"oss-cn-shanghai.aliyuncs.com",
					AccessKeyID:"xxxxx",
					AccessKeySecret:"xxxxxx",
				},
			},

		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Plugin{
				Config:   tt.fields.Config,
			}
			if err := p.Exec(); (err != nil) != tt.wantErr {
				t.Errorf("Plugin.Exec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
