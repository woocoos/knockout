package res

import "testing"

func TestGenEntSchemaRes(t *testing.T) {
	type args struct {
		cfg Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "gen ent schema res",
			args: args{
				cfg: Config{
					AppConfig: "../../adminx/etc/app.yaml",
					EntConfig: "../../../codegen/entgen/schema",
					AppCode:   "resource",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GenEntSchemaRes(tt.args.cfg); (err != nil) != tt.wantErr {
				t.Errorf("GenEntSchemaRes() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
