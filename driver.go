package genieql

import (
	"database/sql"
	sqldriver "database/sql/driver"
	"fmt"
	"go/ast"
	"log"
	"reflect"

	"github.com/pkg/errors"
)

// ErrMissingDriver - returned when a driver has not been registered.
type missingDriver struct {
	driver string
}

func (t missingDriver) Error() string {
	return fmt.Sprintf("requested driver is not registered: '%s'", t.driver)
}

// ErrDuplicateDriver - returned when a ddriver gets registered twice.
var ErrDuplicateDriver = fmt.Errorf("driver has already been registered")

var drivers = driverRegistry{}

// LookupTypeDefinition converts a expression into a type definition.
type LookupTypeDefinition func(typ ast.Expr) (ColumnDefinition, error)

// RegisterDriver register a database driver with genieql. usually in an init function.
func RegisterDriver(driver string, imp Driver) error {
	return drivers.RegisterDriver(driver, imp)
}

// LookupDriver lookup a registered driver.
func LookupDriver(name string) (Driver, error) {
	return drivers.LookupDriver(name)
}

// MustLookupDriver panics if the driver cannot be found, convience method.
func MustLookupDriver(name string) Driver {
	driver, err := LookupDriver(name)
	if err != nil {
		panic(err)
	}
	return driver
}

// PrintRegisteredDrivers print drivers in the registry, debugging utility.
func PrintRegisteredDrivers() {
	for key := range map[string]Driver(drivers) {
		log.Println("Driver", key)
	}
}

// Driver - driver specific details.
type Driver interface {
	LookupType(s string) (ColumnDefinition, error)
	Exported() (res map[string]reflect.Value)
}

type decoder interface {
	sql.Scanner
	sqldriver.Valuer
}

// ColumnDefinition defines a type supported by the driver.
type ColumnDefinition struct {
	Type       string // dialect type
	Native     string // golang type
	ColumnType string // sql type
	Nullable   bool   // does this type represent a pointer type.
	PrimaryKey bool   // is the column part of the primary key
	Decode     string // template function that decodes from the Driver type to Native type
	Encode     string // template function that encodes from the Native type to Driver type
}

type driverRegistry map[string]Driver

func (t driverRegistry) RegisterDriver(driver string, imp Driver) error {
	if _, exists := t[driver]; exists {
		return ErrDuplicateDriver
	}

	t[driver] = imp

	return nil
}

func (t driverRegistry) LookupDriver(name string) (Driver, error) {
	impl, exists := t[name]
	if !exists {
		return nil, missingDriver{driver: name}
	}

	return impl, nil
}

// NewDriver builds a new driver from the component parts
func NewDriver(supported ...ColumnDefinition) Driver {
	return driver{supported: supported}
}

type driver struct {
	supported []ColumnDefinition
}

func (t driver) LookupType(l string) (ColumnDefinition, error) {
	for _, s := range t.supported {
		if s.Type == l {
			return s, nil
		}
	}

	return ColumnDefinition{}, errors.New("missing type")
}

func (t driver) Exported() (res map[string]reflect.Value) {
	res = map[string]reflect.Value{}
	for _, d := range t.supported {
		_ = d
		// TODO:
		// if typ.Decoder == nil {
		// 	continue
		// }

		// switch idx := strings.IndexRune(typ.NullType, '.'); idx {
		// case -1:
		// 	res[typ.NullType] = reflect.ValueOf(typ.Decoder)
		// default:
		// 	res[typ.NullType[idx+1:]] = reflect.ValueOf(typ.Decoder)
		// }
	}

	return res
}
