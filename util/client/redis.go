package client

import "github.com/go-redis/redis"

type IRedis interface {
	// Deployment
	AddDeploymentManifest(namespace, yaml string) (bool, error)
	DeleteDeploymentManifest(namespace string) (bool, error)
	// Service
	AddServiceManifest(namespace, yaml string) (bool, error)
	DeleteServiceManifest(namespace, yaml string) (bool, error)
}

type Redis struct {
	*redis.Client
}

func NewRedisClient(redisClient *redis.Client) IRedis {
	return &Redis{redisClient}
}

func (r *Redis) AddDeploymentManifest(namespace, yaml string) (bool, error) {
	return false, nil
}

func (r *Redis) DeleteDeploymentManifest(namespace string) (bool, error) {
	return false, nil
}

func (r *Redis) AddServiceManifest(namespace, yaml string) (bool, error) {
	return false, nil
}

func (r *Redis) DeleteServiceManifest(namespace, yaml string) (bool, error) {
	return false, nil
}
