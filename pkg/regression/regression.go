// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package regression

import (
	"context"
	"io/ioutil"

	"github.com/go-logr/logr"
	"gopkg.in/yaml.v2"
	"k8s.io/klog/v2"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// ConfigurationYaml provides configuration credentials for a specific storage controller.
type ConfigurationYaml struct {
	Ip        string `yaml:"ip"`        // The IP Address in dot notation of the storage controller
	Username  string `yaml:"username"`  // The username used to log into the storage controller
	Password  string `yaml:"password"`  // The password used to log into the storage controller
	Initiator string `yaml:"initiator"` // The initiator IGN value to use for iSCSI regression testing
	Pool      string `yaml:"pool"`      // The storage pool (A or B) to use for testing
}

// TestConfig provides the configuration for the regression tests. It must be
// constructed with NewTestConfig to initialize it with sane defaults. The
// user of the regression package can then override values before passing
// the instance to [Ginkgo]Test and/or (when using GinkgoTest) in a BeforeEach.
type TestConfig struct {
	StorageController ConfigurationYaml // The storage controller credentials
	Ctx               context.Context   // The cfm-regression context
}

// TestContext gets initialized by the regression package before each test
// runs. It holds the variables that each test can depend on.
type TestContext struct {
	Config *TestConfig
}

// ReadConfigurationYaml: Read configuration file and return Go struct
func ReadConfigurationYaml(ctx context.Context, filename string) (*ConfigurationYaml, error) {

	logger := klog.FromContext(ctx)
	logger.V(4).Info("read configuration setting", "filename", filename)

	// If it is not possible to extract the configuration.yaml from the tar file, use defaults
	var yamlc = ConfigurationYaml{}

	yamlFile, err := ioutil.ReadFile(filename)
	logger.V(4).Info("read file", "err", err)
	if err != nil {
		return &yamlc, err
	}

	err = yaml.Unmarshal(yamlFile, &yamlc)
	logger.V(4).Info("configuration file", "filename", filename, "contents", yamlc)
	return &yamlc, err
}

// NewTestConfig returns a config instance with all values set to values read from a config file.
func NewTestConfig(ctx context.Context, filename string) (*TestConfig, error) {

	config, err := ReadConfigurationYaml(ctx, filename)

	if err != nil {
		return nil, err
	}

	return &TestConfig{
		StorageController: ConfigurationYaml{
			Ip:        config.Ip,
			Username:  config.Username,
			Password:  config.Password,
			Initiator: config.Initiator,
			Pool:      config.Pool,
		},
		Ctx: ctx,
	}, nil
}

// NewContext sets up regression testing with a config supplied by the
// user of the regression package. Ownership of that config is shared
// between the regression package and the caller.
func NewTestContext(config *TestConfig) *TestContext {
	return &TestContext{
		Config: config,
	}
}

// Test will test the application at the specified address by
// setting up a Ginkgo suite and running it.
func Test(t GinkgoTestingT, config *TestConfig, logger logr.Logger) {
	sc := GinkgoTest(config)
	RegisterFailHandler(Fail)

	RunSpecs(t, "CFM Service Regression Test Suite")
	sc.Finalize()
}

// GinkoTest is another entry point for regression testing: instead of
// directly running tests like Test does, it merely registers the
// tests. This can be used to embed regression testing in a custom Ginkgo
// test suite.  The pointer to the configuration is merely stored by
// GinkgoTest for use when the tests run. Therefore its content can
// still be modified in a BeforeEach. The regression package itself treats
// it as read-only.
func GinkgoTest(config *TestConfig) *TestContext {
	sc := NewTestContext(config)
	registerTestsInGinkgo(sc)
	return sc
}

// Setup must be invoked before each test. It initialize per-test
// variables in the context.
func (sc *TestContext) Setup() {
	// TODO
}

// Teardown must be called after each test. It frees resources
// allocated by Setup.
func (sc *TestContext) Teardown() {
	// TODO
}

// Finalize frees any resources that might be still cached in the context.
// It should be called after running all tests.
func (sc *TestContext) Finalize() {
	// TODO
}
