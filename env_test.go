package env

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func assertEqual(t *testing.T, expected, actual any) {
	if expected != actual {
		t.Errorf("Not equal: \nexpected: %v\n actual  : %v",
			expected,
			actual)
	}
}

func TestGet(t *testing.T) {
	envVar := "MY_ENV"
	t.Cleanup(func() { os.Unsetenv(envVar) })

	os.Unsetenv(envVar)
	val := Get(envVar, "default_value")
	assertEqual(t, "default_value", val)
	log.Println("env=Nil, Get('default_value'), return 'default_value")

	os.Setenv(envVar, "")
	val = Get(envVar, "default_value")
	assertEqual(t, "", val)
	log.Println("env='', Get('default_value'), return ''")

	os.Unsetenv(envVar)
	os.Setenv(envVar, "defined_value")
	val = Get(envVar, "default_value")
	assertEqual(t, "defined_value", val)
	log.Println("env='defined_value', Get('default_value'), return 'defined_value")
}

func TestGetInt(t *testing.T) {
	envVar := "MY_INT_ENV"
	t.Cleanup(func() {
		os.Unsetenv(envVar)
	})

	os.Unsetenv(envVar)
	val := GetInt(envVar, 10)
	assertEqual(t, 10, val)
	log.Println("env=nil, GetInt(10), return 10")

	os.Setenv(envVar, "1")
	val = GetInt(envVar, 10)
	assertEqual(t, 1, val)
	log.Println("env=1, GetInt(10), return 1")

	os.Unsetenv(envVar)

	os.Setenv(envVar, "notAnInt")
	val = GetInt(envVar, 10)
	assertEqual(t, 10, val)
	log.Println("env=notAnInt, GetInt(10), return 10")
}

func TestGetInt64(t *testing.T) {
	envVar := "MY_INT64_ENV"
	t.Cleanup(func() { os.Unsetenv(envVar) })

	os.Unsetenv(envVar)
	val := GetInt64(envVar, 10)
	assertEqual(t, int64(10), val)
	log.Println("env=nil, GetInt64(10), return 10")

	os.Setenv(envVar, "8")
	val = GetInt64(envVar, 10)
	assertEqual(t, int64(8), val)
	log.Println("env=8, GetInt64(10), return 8")

	os.Setenv(envVar, "notAnInt64")
	val = GetInt64(envVar, 10)
	assertEqual(t, int64(10), val)
	log.Println("env=notAnInt64, GetInt64(10), return 10")
}

func TestGetBool(t *testing.T) {
	envVar := "MY_BOOL_ENV"
	t.Cleanup(func() { os.Unsetenv(envVar) })

	val := GetBool(envVar, true)
	assertEqual(t, true, val)
	log.Println("env=nil, GetBool(true) return true")

	os.Setenv(envVar, "8")
	val = GetBool(envVar, false)
	assertEqual(t, false, val)
	log.Println("env=8, GetBool(false) return false")

	os.Setenv(envVar, "8")
	val = GetBool(envVar, true)
	assertEqual(t, true, val)
	log.Println("env=8, GetBool(true) return true")

	os.Setenv(envVar, "0")
	val = GetBool(envVar, true)
	assertEqual(t, false, val)
	log.Println("env=0, GetBool(true) return false")

	os.Setenv(envVar, "")
	val = GetBool(envVar, true)
	assertEqual(t, true, val)
	log.Println("env='', GetBool(true) return true")

	os.Setenv(envVar, "")
	val = GetBool(envVar, false)
	assertEqual(t, false, val)
	log.Println("env='', GetBool(false) return false")

	os.Unsetenv(envVar)
	os.Setenv(envVar, "notAnBool")
	val = GetBool(envVar, true)
	assertEqual(t, true, val)
	log.Println("env='notAnBool', GetBool(true) return true")

	os.Unsetenv(envVar)
	os.Setenv(envVar, "notAnBool")
	val = GetBool(envVar, false)
	assertEqual(t, false, val)
	log.Println("env='notAnBool', GetBool(false) return false")

	os.Setenv(envVar, "True")
	val = GetBool(envVar, false)
	assertEqual(t, true, val)
	log.Println("env='True', GetBool(false) return true")

	os.Setenv(envVar, "TRUE")
	val = GetBool(envVar, false)
	assertEqual(t, true, val)
	log.Println("env='TRUE', GetBool(false) return true")
}

func checkEqualArray(t *testing.T, expected, actual []string) {
	if len(expected) != len(actual) {
		t.Errorf("Not equal Array Len: \nexpected: %v=%d\nactual  : %v=%d",
			expected, len(expected),
			actual, len(actual))
	}

	for i, v := range expected {
		if v != actual[i] {
			t.Errorf("Not equal Array: \nexpected: %v\nactual  : %v",
				expected,
				actual)
		}
	}
}

func TestGetArray(t *testing.T) {
	envVar := "MY_ARRAY_ENV"
	t.Cleanup(func() { os.Unsetenv(envVar) })
	DefaultValue := []string{"foo", "bar"}
	DefinedValue0 := []string{"hello", "world"}
	DefinedValue1 := []string{"helloWorld"}
	DefinedValue2 := []string{}

	val := GetArray(envVar, DefaultValue)
	checkEqualArray(t, DefaultValue, val)
	log.Println("env=nil, GetArray() is OK")

	os.Setenv(envVar, "hello,world")
	val = GetArray(envVar, DefaultValue)
	checkEqualArray(t, DefinedValue0, val)
	log.Println("env={foo 0}, GetArray() is OK")

	os.Setenv(envVar, "helloWorld")
	val = GetArray(envVar, DefaultValue)
	checkEqualArray(t, DefinedValue1, val)
	log.Println("env={foo 1}, GetArray() is OK")

	os.Setenv(envVar, "")
	val = GetArray(envVar, DefinedValue2)
	checkEqualArray(t, DefinedValue2, val)
	log.Println("env='', GetArray() is OK")
}

func TestGetRequired(t *testing.T) {
	envVar := "MY_REQ_STRING_ENV"
	t.Cleanup(func() { os.Unsetenv(envVar) })

	os.Unsetenv(envVar)
	defer func() {
		log.Println("Test Get() for Required Variable is OK")
		_ = recover()
	}()
	Get(envVar)
	t.Errorf("Required Variable did not panic")
}

func TestGetIntRequired(t *testing.T) {
	envVar := "MY_REQ_INT_ENV"
	t.Cleanup(func() { os.Unsetenv(envVar) })

	os.Unsetenv(envVar)
	defer func() {
		log.Println("Test GetInt() for Required Variable is OK")
		_ = recover()
	}()
	GetInt(envVar)
	t.Errorf("Int Required Variable did not panic")
}

func TestGetInt64Required(t *testing.T) {
	envVar := "MY_REQ_INT64_ENV"
	t.Cleanup(func() { os.Unsetenv(envVar) })

	os.Unsetenv(envVar)
	defer func() {
		log.Println("Test GetInt64() for Required Variable is OK")
		_ = recover()
	}()
	GetInt64(envVar)
	t.Errorf("int64 Required Variable did not panic")
}

func TestGetFloatRequired(t *testing.T) {
	envVar := "MY_REQ_FLOAT_ENV"
	t.Cleanup(func() { os.Unsetenv(envVar) })

	os.Unsetenv(envVar)
	defer func() {
		log.Println("Test GetFloat() for Required Variable is OK")
		_ = recover()
	}()
	GetFloat(envVar)
	t.Errorf("Float Required Variable did not panic")
}

func TestGetBoolRequired(t *testing.T) {
	envVar := "MY_REQ_BOOL_ENV"
	t.Cleanup(func() { os.Unsetenv(envVar) })

	os.Unsetenv(envVar)
	defer func() {
		log.Println("Test GetBool() for Required Variable is OK")
		_ = recover()
	}()
	GetBool(envVar)
	t.Errorf("Bool Required Variable did not panic")
}

func TestGetArrayRequired(t *testing.T) {
	envVar := "MY_ARRAY_ENV"
	t.Cleanup(func() { os.Unsetenv(envVar) })

	os.Unsetenv(envVar)
	defer func() {
		log.Println("Test GetArray() for Required Variable is OK")
		_ = recover()
	}()
	GetArray(envVar, nil)
	t.Errorf("Array Required Variable did not panic")
}

func ExampleGet() {
	fmt.Println(Get("MY_ENV_VAR", "default"))
	os.Setenv("MY_ENV_VAR", "custom")
	fmt.Println(Get("MY_ENV_VAR", "default"))
	os.Unsetenv("MY_ENV_VAR")

	// Output:
	// default
	// custom
}

func ExampleGetBool() {
	fmt.Println(GetBool("MY_ENV_VAR", true))
	os.Setenv("MY_ENV_VAR", "false")
	fmt.Println(GetBool("MY_ENV_VAR", true))
	os.Unsetenv("MY_ENV_VAR")

	// Output:
	// true
	// false
}

func ExampleGetInt() {
	fmt.Println(GetInt("MY_ENV_VAR", 1))
	os.Setenv("MY_ENV_VAR", "100")
	fmt.Println(GetInt("MY_ENV_VAR", 1))
	os.Unsetenv("MY_ENV_VAR")

	// Output:
	// 1
	// 100
}

func ExampleGetInt64() {
	fmt.Println(GetInt64("MY_ENV_VAR", int64(12345678910)))
	os.Setenv("MY_ENV_VAR", "100")
	fmt.Println(GetInt64("MY_ENV_VAR", int64(12345678910)))
	os.Unsetenv("MY_ENV_VAR")

	// Output:
	// 12345678910
	// 100
}

func ExampleGetFloat() {
	fmt.Println(GetFloat("MY_ENV_VAR", 3.14))
	os.Setenv("MY_ENV_VAR", "34.02")
	fmt.Println(GetFloat("MY_ENV_VAR", 3.14))
	os.Unsetenv("MY_ENV_VAR")

	// Output:
	// 3.14
	// 34.02
}
