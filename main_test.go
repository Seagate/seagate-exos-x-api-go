package exosx

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	. "github.com/onsi/gomega"
	"k8s.io/klog"
)

var client = NewClient()

var logging = flag.Bool("logging", true, "logging")

func init() {
	var exists bool
	settingsfile := ".env"

	// mock.LoadEnv()

	fmt.Printf("Testing setup: logging (%v)\n", *logging)

	if *logging {
		klog.InitFlags(nil)
	}

	// Note, any defined environment variable is used over the ones defined in .env
	if _, err := os.Stat(settingsfile); err == nil {
		fmt.Printf("Testing setup: Loading (%s)\n", settingsfile)
		err := godotenv.Load(settingsfile)
		if err != nil {
			fmt.Printf("Error loading file (%s), error: %v\n", settingsfile, err)
			return
		}
	}

	client.Addr, exists = os.LookupEnv("TEST_STORAGEIP")
	if exists {
		fmt.Printf("Testing setup: %s=%s\n", "TEST_STORAGEIP", client.Addr)
	}

	client.Username, exists = os.LookupEnv("TEST_USERNAME")
	if exists {
		fmt.Printf("Testing setup: %s=%s\n", "TEST_USERNAME", client.Username)
	}

	client.Password, exists = os.LookupEnv("TEST_PASSWORD")
	if exists {
		fmt.Printf("Testing setup: %s=%s\n", "TEST_PASSWORD", client.Password)
	}
}

func assert(t *testing.T, cond bool, msg string) {
	if !cond {
		t.Errorf(msg)
	} else {
		fmt.Printf("%s: OK\n", msg)
	}
}

func TestLogin(t *testing.T) {
	t.Logf("args: %s", strings.Join(flag.Args(), " "))
	g := NewWithT(t)
	g.Expect(client.Login()).To(BeNil())
}

func TestLoginFailed(t *testing.T) {
	var wrongClient = NewClient()
	wrongClient.Addr = client.Addr
	wrongClient.Username = client.Username
	wrongClient.Password = "wrongpassword"

	g := NewWithT(t)
	g.Expect(wrongClient.Login()).ToNot(BeNil())
}

func TestReLoginFailed(t *testing.T) {
	var wrongClient = NewClient()
	wrongClient.Addr = client.Addr
	wrongClient.Username = client.Username
	wrongClient.Password = client.Password

	g := NewWithT(t)
	g.Expect(wrongClient.Login()).To(BeNil())

	wrongClient.Password = "wrongpassword"
	wrongClient.SessionKey = "outdated-session-key"

	_, status, err := wrongClient.FormattedRequest("/bad/request")
	g.Expect(err).NotTo(BeNil())
	g.Expect(status.ResponseType).To(Equal("Error"))
	// This test returns one of three different values based on the  API version.
	g.Expect(status.Response).Should(BeElementOf([]string{"re-login failed", "request failed", "Invalid sessionkey"}))
}

func TestInvalidURL(t *testing.T) {
	g := NewWithT(t)
	_, status, err := client.FormattedRequest("/trololol")

	g.Expect(err).NotTo(BeNil())
	g.Expect(status.ResponseType).To(Equal("Error"))
	g.Expect(status.Response).To(Equal("request failed"))
}

func TestInvalidXML(t *testing.T) {
	g := NewWithT(t)
	_, status, err := client.FormattedRequest("/invalid/xml")

	g.Expect(err).NotTo(BeNil())
	g.Expect(status.ResponseType).To(Equal("Error"))
	g.Expect(status.Response).To(Equal("request failed"))
}

func TestStatusCodeNotZero(t *testing.T) {
	g := NewWithT(t)
	_, status, err := client.FormattedRequest("/status/code/1")

	g.Expect(err).NotTo(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))
}
