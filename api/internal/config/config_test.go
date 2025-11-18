package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		env     string
		want    *Config
		wantErr bool
	}{
		{
			name: "success",
			path: "../../.conf",
			env:  "local",
			want: &Config{
				Server: Server{
					Port: ":3000",
				},
				Database: Database{
					Region:    "eu-west-1",
					TableName: "picusv3",
				},
			},
			wantErr: false,
		},
		{
			name: "error - config file not found",
			path: "../.conf",
			env:  "local",
			want: &Config{
				Server: Server{
					Port: ":3000",
				},
				Database: Database{
					Region:    "eu-west-1",
					TableName: "picusv3",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadConfig(tt.path, tt.env)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, got)
			assert.Equal(t, tt.want.Server.Port, got.Server.Port)
			assert.Equal(t, tt.want.Database.Region, got.Database.Region)
			assert.Equal(t, tt.want.Database.TableName, got.Database.TableName)
		})
	}
}
