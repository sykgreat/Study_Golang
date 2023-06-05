package server

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"log"
)

const etcdUrl = "http://localhost:2379"
const serviceName = "syk/server"
const ttl = 10

var etcdClient *clientv3.Client

func etcdRegister(addr string) error {
	log.Printf("etcdRegister %s\n", addr)

	etcdClient, err := clientv3.NewFromURL(etcdUrl)
	if err != nil {
		return err
	}

	em, err := endpoints.NewManager(etcdClient, serviceName)
	if err != nil {
		return err
	}

	lease, err := etcdClient.Grant(context.TODO(), ttl)
	if err != nil {
		return err
	}

	err = em.AddEndpoint(context.TODO(), fmt.Sprintf("%s/%s", serviceName, addr), endpoints.Endpoint{Addr: addr}, clientv3.WithLease(lease.ID))
	if err != nil {
		return err
	}

	return nil
}

func etcdUnRegister(addr string) error {
	return nil
}
