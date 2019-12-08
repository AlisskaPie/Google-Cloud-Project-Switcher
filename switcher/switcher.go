package switcher

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"golang.org/x/sys/windows/registry"
)

var (
	langNum int64
	projNum int64
)

func setEnvironment(key string, value string) error {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\ControlSet001\Control\Session Manager\Environment`, registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer k.Close()

	err = k.SetStringValue(key, value)
	if err != nil {
		return err
	}
	return nil
}
func setEnvir(env map[string]interface{}) error {
	keys := env["env_variables"].([]interface{})
	var err error
	for _, i := range keys {
		l := i.(map[string]interface{})
		for val, j := range l {
			err = setEnvironment(val, j.(string))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func projectSwt(projNum int64, proj map[string]interface{}) error {
	switch projNum {
	case 1:
		env := proj["Storage"].(map[string]interface{})
		err := setEnvir(env)
		if err != nil {
			return err
		}
		return nil
	case 2:
		env := proj["Pubsub"].(map[string]interface{})
		err := setEnvir(env)
		if err != nil {
			return err
		}
		return nil
	case 3:
		env := proj["Spanner"].(map[string]interface{})
		err := setEnvir(env)
		if err != nil {
			return err
		}
		return nil
	case 4:
		env := proj["Firestore"].(map[string]interface{})
		err := setEnvir(env)
		if err != nil {
			return err
		}
		return nil
	case 5:
		env := proj["BigQuery"].(map[string]interface{})
		err := setEnvir(env)
		if err != nil {
			return err
		}
		return nil
	case 6:
		env := proj["BigTable"].(map[string]interface{})
		err := setEnvir(env)
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("projectSwt error: cannot switch to project")
	}
}

// ScannerLines sets environment variables for different programming languages and google cloud projects.
func ScannerLines() error {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Choose number of the language:\n  1 - Python\n  2 - NodeJS\n  3 - Go\n  4 - PHP\n  5 - Java\n  6 - Ruby\n  7 - C#\n  8 - C++")
	var err error
	for scanner.Scan() {
		langNum, err = strconv.ParseInt(scanner.Text(), 10, 0)
		if err != nil {
			return errors.New("language error: cannot parse string to int")
		}
		if langNum > 8 || langNum < 1 {
			return errors.New("language error: not in this list [1:8]")
		}
		if len(scanner.Text()) >= 1 {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		return err
	}
	fmt.Println("Choose number of the project:\n  1 - Storage\n  2 - Pubsub\n  3 - Spanner\n  4 - Firestore\n  5 - BigQuery\n  6 - BigTable")
	for scanner.Scan() {
		projNum, err = strconv.ParseInt(scanner.Text(), 10, 0)
		if err != nil {
			return errors.New("project error: cannot parse string to int")
		}
		if projNum > 6 || projNum < 1 {
			return errors.New("project error: not in this list [1,6]")
		}
		if len(scanner.Text()) >= 1 {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		return err
	}
	// Open our jsonFile with settings.
	jsonFile, err := os.Open("switcher/settings.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var lang map[string]interface{}
	json.Unmarshal([]byte(byteValue), &lang)

	switch langNum {
	case 1:
		err := projectSwt(projNum, lang["Python"].(map[string]interface{}))
		if err != nil {
			return err
		}
		return nil
	case 2:
		err := projectSwt(projNum, lang["NodeJS"].(map[string]interface{}))
		if err != nil {
			return err
		}
		return nil
	case 3:
		err := projectSwt(projNum, lang["Go"].(map[string]interface{}))
		if err != nil {
			return err
		}
		return nil
	case 4:
		err := projectSwt(projNum, lang["PHP"].(map[string]interface{}))
		if err != nil {
			return err
		}
		return nil
	case 5:
		err := projectSwt(projNum, lang["Java"].(map[string]interface{}))
		if err != nil {
			return err
		}
		return nil
	case 6:
		err := projectSwt(projNum, lang["Ruby"].(map[string]interface{}))
		if err != nil {
			return err
		}
		return nil
	case 7:
		err := projectSwt(projNum, lang["C#"].(map[string]interface{}))
		if err != nil {
			return err
		}
		return nil
	case 8:
		err := projectSwt(projNum, lang["C++"].(map[string]interface{}))
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("cannot switch language")
	}
}
