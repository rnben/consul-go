package consul_go

import (
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/hashicorp/consul/api"
)

// NewConsulRegister create a new consul register
func NewConsulRegister() *ConsulRegister {
	return &ConsulRegister{
		Config:                         api.DefaultConfig(),
		DeregisterCriticalServiceAfter: time.Duration(5) * time.Minute,
		Interval:                       time.Duration(5) * time.Second,
	}
}

// ConsulRegister consul service register
type ConsulRegister struct {
	Config                         *api.Config
	DeregisterCriticalServiceAfter time.Duration
	Interval                       time.Duration
}

// Register register service
func (r *ConsulRegister) Register(serviceID string, serviceName string, servicePort int) error {
	if serviceID == "" || serviceName == "" || servicePort == 0 {
		return errors.New("service info is null")
	}
	client, err := api.NewClient(r.Config)
	if err != nil {
		return err
	}
	agent := client.Agent()

	IP := localIP()
	reg := &api.AgentServiceRegistration{
		ID:      serviceID,   // 服务唯一标识
		Name:    serviceName, // 服务名称
		Port:    servicePort, // 服务端口
		Address: IP,          // 服务 IP
		Check: &api.AgentServiceCheck{
			Interval:                       r.Interval.String(),
			GRPC:                           fmt.Sprintf("%v:%v/%v", IP, servicePort, ""),
			DeregisterCriticalServiceAfter: r.DeregisterCriticalServiceAfter.String(),
		},
	}
	return agent.ServiceRegister(reg)
}

// DeregisterRegister 反注册服务
func (r *ConsulRegister) DeregisterRegister(serviceID string) {
	if serviceID == "" {
		return
	}
	consul, err := api.NewClient(r.Config)
	if err != nil {
		log.Fatalln(err)
	}
	err = consul.Agent().ServiceDeregister(serviceID)
	if err != nil {
		log.Fatalln(err)
	}
}

func localIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
