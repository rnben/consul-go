package consul_go

import (
	"github.com/hashicorp/consul/api"
	"testing"
	"time"
)

func TestConsulRegister_Register(t *testing.T) {
	type fields struct {
		Config                         *api.Config
		DeregisterCriticalServiceAfter time.Duration
		Interval                       time.Duration
	}
	type args struct {
		serviceName string
		servicePort int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "测试注册",
			fields: fields{
				Config: &api.Config{
					Address: "127.0.0.1:8500",
					Token:   "",
				},
				DeregisterCriticalServiceAfter: 1,
				Interval:                       5,
			},
			args: args{
				serviceName: "course",
				servicePort: 13821,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ConsulRegister{
				Config:                         tt.fields.Config,
				DeregisterCriticalServiceAfter: tt.fields.DeregisterCriticalServiceAfter,
				Interval:                       tt.fields.Interval,
			}
			if err := r.Register(tt.args.serviceName, tt.args.servicePort); (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConsulRegister_DeregisterRegister(t *testing.T) {
	type fields struct {
		Config                         *api.Config
		DeregisterCriticalServiceAfter time.Duration
		Interval                       time.Duration
	}
	type args struct {
		serviceID string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "测试反注册",
			fields: fields{
				Config: &api.Config{
					Address: "127.0.0.1:8500",
					Token:   "",
				},
				DeregisterCriticalServiceAfter: 1,
				Interval:                       5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ConsulRegister{
				Config:                         tt.fields.Config,
				DeregisterCriticalServiceAfter: tt.fields.DeregisterCriticalServiceAfter,
				Interval:                       tt.fields.Interval,
			}
			r.DeregisterRegister("course-192.1.203.240-course")
		})
	}
}
