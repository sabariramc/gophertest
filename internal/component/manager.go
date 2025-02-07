package component

import (
	"context"
	"runtime/debug"

	"gitlab.com/engineering/products/api_security/go-common/log"
	"gitlab.com/engineering/products/api_security/go-common/log/factory"
	"gitlab.com/engineering/products/api_security/go-common/tester/dependency"
)

type TestDependencyManager struct {
	dependency.TestDependencyManager
	log log.Log
}

func NewTestDependencyManager(setupNeeded bool) *TestDependencyManager {
	manager := new(TestDependencyManager)
	manager.SetupNeeded = setupNeeded
	manager.log = factory.NewLog(context.Background(), "TestDependencyManager")
	return manager
}

func (m *TestDependencyManager) Setup() error {
	err := m.SetupDockerNetwork("gophertest")
	if err != nil {
		return err
	}

	if !m.SetupNeeded {
		return nil
	}

	redisConfig := dependency.RedisConfig{
		Host: "redis01",
		Port: "6379" + "/tcp",
	}
	m.AddDependency(dependency.NewRedisDependency(m.DockerTestPool, &redisConfig))
	m.log.Infof(context.Background(), "Creating a fresh setup of Docker based dependencies.")
	err = m.CreateDependencies()
	if err != nil {
		return err
	}

	return nil
}

func (m *TestDependencyManager) Teardown() error {
	defer func() {
		if r := recover(); r != nil {
			m.log.Debugf(context.Background(), string(debug.Stack()))
		}
	}()
	if !m.SetupNeeded {
		return nil
	}
	m.log.Infof(context.Background(), "Destroying the created setup of Docker based dependencies")
	err := m.DestroyDependencies()
	if err != nil {
		m.log.Errorf(context.Background(), "error destroying the created setup of Docker based dependencies: %v", err.Error())
		return err
	}
	m.log.Infof(context.Background(), "Destroying docker network")
	err = m.TeardownDockerNetwork()
	if err != nil {
		m.log.Errorf(context.Background(), "error destroying docker network: %v", err.Error())
		return err
	}
	return nil
}
