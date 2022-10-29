package util_test

import (
	"os"
	"testing"
	"util"
)

func TestMustGetEnv(t *testing.T) {
	cases := []string{
		"244",
		"lmao_ye",
		"l333l_55555555",
		"TEST_ENV",
	}

	for _, c := range cases {
		happyCase := "TEST_ENV_EX_" + c
		sadCase := "TEST_ENV_NEX_" + c

		os.Setenv(happyCase, "test")
		t.Run("Should not panic when passed "+happyCase, func(t *testing.T) {
			t.Parallel()

			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Should not have panicked")
				}
			}()

			func() {
				if v := util.MustGetEnv(happyCase); v != "test" {
					t.Errorf("Expected value to be 'test', got %s", v)
				}
			}()
		})

		os.Unsetenv(sadCase)
		t.Run("Should panic when passed "+sadCase, func(t *testing.T) {
			t.Parallel()

			defer func() {
				if r := recover(); r == nil {
					t.Errorf("Should  have panicked")
				}
			}()

			func() {
				util.MustGetEnv(sadCase)
			}()
		})
	}

}
