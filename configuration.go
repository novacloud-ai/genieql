package genieql

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/go-yaml/yaml"
	"github.com/pkg/errors"
)

// ErrRequireHostAndPort currently require both host and port to be specified.
var ErrRequireHostAndPort = fmt.Errorf("both host and port are required")

// Configuration main configuration for genieql. usually generated by the Bootstrap
// function.
type Configuration struct {
	Location      string
	Name          string
	Dialect       string
	Driver        string
	Queryer       string
	ConnectionURL string
	Host          string
	Port          int
	Database      string
	Username      string
	Password      string
}

func (t Configuration) ReadMap(pkg, typ, name string, m *MappingConfig) error {
	return ReadMapper(t, pkg, typ, name, m)
}

func (t Configuration) WriteMap(name string, m MappingConfig) error {
	return WriteMapper(t, name, m)
}

// Bootstrap takes a db connection url and creates a genieql
// configuration from the url and writes out the configuration to the provided
// path.
func Bootstrap(options ...ConfigurationOption) error {
	var (
		err    error
		config Configuration
	)

	if config, err = NewConfiguration(options...); err != nil {
		return err
	}

	if err := os.MkdirAll(config.Location, 0755); err != nil {
		return errors.Wrap(err, "failed to make bootstrap directory")
	}

	return WriteConfiguration(config)
}

// WriteConfiguration writes the genieql configuration file to the specified path.
func WriteConfiguration(config Configuration) error {
	var (
		err error
		raw []byte
	)

	if raw, err = yaml.Marshal(config); err != nil {
		return errors.Wrap(err, "failed to serialize configuration to yaml")
	}

	return errors.Wrap(ioutil.WriteFile(filepath.Join(config.Location, config.Name), raw, 0666), "failed to persist configuration to disk")
}

// ReadConfiguration reads the genieql configuration file to the specified path.
func ReadConfiguration(config *Configuration) error {
	var (
		err error
		raw []byte
	)
	if raw, err = ioutil.ReadFile(filepath.Join(config.Location, config.Name)); err != nil {
		return errors.Wrap(err, "failed to read configuration file")
	}

	return errors.Wrap(yaml.Unmarshal(raw, config), "failed to parse configuration file")
}

// MustConfiguration builds a configuration from the provided options.
func MustConfiguration(options ...ConfigurationOption) Configuration {
	var (
		err error
		c   Configuration
	)

	if c, err = NewConfiguration(options...); err != nil {
		log.Fatalf("%+v\n", err)
	}

	return c
}

// MustReadConfiguration builds a new configuration from the provided options,
// and read's it from disk.
func MustReadConfiguration(options ...ConfigurationOption) Configuration {
	c := MustConfiguration(options...)
	if e := ReadConfiguration(&c); e != nil {
		log.Fatalf("%+v\n", e)
	}
	return c
}

// NewConfiguration builds a configuration from the provided options.
func NewConfiguration(options ...ConfigurationOption) (Configuration, error) {
	var (
		config = Configuration{
			Queryer: "*sql.DB",
		}
	)

	for _, opt := range options {
		if err := opt(&config); err != nil {
			return config, err
		}
	}

	return config, nil
}

// ConfigurationOption options for creating a Configuration
type ConfigurationOption func(*Configuration) error

// ConfigurationOptionLocation specify the absolute path of the configuration file.
func ConfigurationOptionLocation(path string) ConfigurationOption {
	return func(c *Configuration) error {
		c.Location, c.Name = filepath.Dir(path), filepath.Base(path)
		return nil
	}
}

// ConfigurationOptionQueryer specify the default queryer to use.
func ConfigurationOptionQueryer(queryer string) ConfigurationOption {
	return func(c *Configuration) error {
		c.Queryer = queryer
		return nil
	}
}

// ConfigurationOptionDatabase specify the database connection information.
func ConfigurationOptionDatabase(uri *url.URL) ConfigurationOption {
	return func(c *Configuration) error {
		splits := strings.Split(uri.Host, ":")
		if len(splits) != 2 {
			return ErrRequireHostAndPort
		}

		host, portString := splits[0], splits[1]

		port, err := strconv.Atoi(portString)
		if err != nil {
			return err
		}

		c.ConnectionURL = uri.String()
		c.Dialect = uri.Scheme
		c.Host = host
		c.Port = port
		c.Database = strings.Trim(uri.Path, "/")
		c.Username = uri.User.Username()
		c.Password, _ = uri.User.Password()
		return nil
	}
}

// ConfigurationOptionDriver specify the driver for the configuration.
func ConfigurationOptionDriver(driver string) ConfigurationOption {
	return func(c *Configuration) error {
		c.Driver = driver
		return nil
	}
}
