package client

import (
	"github.com/go-redis/redis"
)

type IRedis interface {
	// Deployment
	AddDeploymentManifest(namespace, deploymentName, deployment string) (bool, error)
	DeleteDeploymentManifest(namespace, deployName string) (bool, error)
	// Service
	AddServiceManifest(namespace, serviceName, service string) (bool, error)
	DeleteServiceManifest(namespace, svcName string) (bool, error)
}

type Redis struct {
	*redis.Client
}

func NewRedisClient(redisClient *redis.Client) IRedis {
	return &Redis{redisClient}
}

func (r *Redis) AddDeploymentManifest(namespace, deployName, deployment string) (bool, error) {
	key := namespace + "-" + deployName
	err := r.Client.Set(key, deployment, 0).Err()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *Redis) DeleteDeploymentManifest(namespace, deployName string) (bool, error) {
	key := namespace + "-" + deployName
	err := r.Client.Del(key).Err()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *Redis) AddServiceManifest(namespace, svcName, service string) (bool, error) {
	key := namespace + "-" + svcName
	err := r.Client.Set(key, service, 0).Err()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *Redis) DeleteServiceManifest(namespace, svcName string) (bool, error) {
	key := namespace + "-" + svcName
	err := r.Client.Del(key).Err()
	if err != nil {
		return false, err
	}
	return true, nil
}
