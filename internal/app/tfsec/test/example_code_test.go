package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/aquasecurity/tfsec/internal/app/tfsec/testutil"

	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
)

func TestExampleCode(t *testing.T) {
	for _, check := range scanner.GetRegisteredRules() {

		t.Run(fmt.Sprintf("Rule explanation for %s", check.ID()), func(t *testing.T) {
			if strings.TrimSpace(check.Documentation.Explanation) == "" {
				t.Fatalf("No explanation found for %s", check.ID())
			}
		})

		t.Run(fmt.Sprintf("Rule impact for %s", check.ID()), func(t *testing.T) {
			if strings.TrimSpace(check.Documentation.Impact) == "" {
				t.Fatalf("No impact found for %s", check.ID())
			}
		})

		t.Run(fmt.Sprintf("Rule resolution for %s", check.ID()), func(t *testing.T) {
			if strings.TrimSpace(check.Documentation.Resolution) == "" {
				t.Fatalf("No resolution found for %s", check.ID())
			}
		})

		t.Run(fmt.Sprintf("Rule 'good' example code for %s", check.ID()), func(t *testing.T) {
			if strings.TrimSpace(check.Documentation.GoodExample) == "" {
				t.Fatalf("good example code not provided for %s", check.ID())
			}
			defer func() {
				if err := recover(); err != nil {
					t.Fatalf("Scan (good) failed: %s", err)
				}
			}()
			results := testutil.ScanHCL(check.Documentation.GoodExample, t)
			testutil.AssertCheckCode(t, "", check.ID(), results)
		})

		t.Run(fmt.Sprintf("Rule 'bad' example code for %s", check.ID()), func(t *testing.T) {
			if strings.TrimSpace(check.Documentation.BadExample) == "" {
				t.Fatalf("bad example code not provided for %s", check.ID())
			}
			defer func() {
				if err := recover(); err != nil {
					t.Fatalf("Scan (bad) failed: %s", err)
				}
			}()
			results := testutil.ScanHCL(check.Documentation.BadExample, t)
			testutil.AssertCheckCode(t, check.ID(), "", results)
		})

	}
}
