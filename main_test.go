package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/bep/helpers/envhelpers"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestScripts(t *testing.T) {
	setup := testSetupFunc()
	testscript.Run(t, testscript.Params{
		Dir: "testscripts",
		// UpdateScripts: true, // Uncomment to rewrite the test scripts with
		// TestWork: true, // Uncomment to keep the test work dir.
		Setup: func(env *testscript.Env) error {
			return setup(env)
		},
	})
}

func TestMain(m *testing.M) {
	os.Exit(
		testscript.RunMain(m, map[string]func() int{
			"gomaintemplate": func() int {
				fmt.Println(strings.Join(os.Args[1:], " "))
				time.Sleep(10 * time.Millisecond)
				return 0
			},
		}),
	)
}

func testSetupFunc() func(env *testscript.Env) error {
	sourceDir, _ := os.Getwd()
	return func(env *testscript.Env) error {
		var keyVals []string
		// Add some environment variables to the test script.
		keyVals = append(keyVals, "SOURCE", sourceDir)
		envhelpers.SetEnvVars(&env.Vars, keyVals...)

		return nil
	}
}
