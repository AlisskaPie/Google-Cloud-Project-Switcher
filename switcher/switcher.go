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

func projectSwt(projNum int64, proj map[string]interface{}, projMap map[int64]string) error {
	env := proj[projMap[projNum]].(map[string]interface{})
	err := setEnvir(env)
	if err != nil {
		return err
	}
	return nil
}

// ScannerLines sets environment variables for different programming languages and google cloud projects.
func ScannerLines() error {
	// Open our jsonFile with settings.
	jsonFile, err := os.Open("switcher/settings.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var lang map[string]interface{}
	json.Unmarshal([]byte(byteValue), &lang)
	scanner := bufio.NewScanner(os.Stdin)

	// Create map for languages and their corresponding numbers.
	langMap := make(map[int64]string, len(lang))
	fmt.Println("Choose number of the language:")
	for key := range lang {
		langNum++
		langMap[langNum] = key
		fmt.Printf("%d - %s\n", langNum, key)
	}
	// Read the number of language from console.
	for scanner.Scan() {
		langNum, err := strconv.ParseInt(scanner.Text(), 10, 0)
		if err != nil {
			return errors.New("language error: cannot parse string to int")
		}
		if langNum > int64(len(lang)) || langNum < 1 {
			return errors.New("language error: not from this list")
		}
		if len(scanner.Text()) >= 1 {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		return err
	}

	// Create map for projects and their corresponding numbers.
	projMap := make(map[int64]string, len(lang))
	fmt.Println("Choose number of the project:")
	proj := lang[langMap[langNum]].(map[string]interface{})
	for key := range proj {
		projNum++
		projMap[projNum] = key
		fmt.Printf("%d - %s\n", projNum, key)
	}

	// Read the number of projects from console for defined language.
	for scanner.Scan() {
		projNum, err = strconv.ParseInt(scanner.Text(), 10, 0)
		if err != nil {
			return errors.New("project error: cannot parse string to int")
		}
		if projNum > int64(len(projMap)) || projNum < 1 {
			return errors.New("project error: not from this list")
		}
		if len(scanner.Text()) >= 1 {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		return err
	}

	err = projectSwt(projNum, proj, projMap)
	if err != nil {
		return err
	}
	return nil
	// return errors.New("cannot switch language")
}
