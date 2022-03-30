import (
	"math/big"
	"regexp"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/madlambda/spells/assert"
	"github.com/zclconf/go-cty/cty"
)

func FuzzPartialEval(f *testing.F) {
	// START SEED OMIT

	seedCorpus := []string{
		"attr",
		"attr.value",
		"attr.*.value",
		"global.str",
		`"a ${global.str}"`,
		`"${global.obj}"`,
		`"${global.list} fail`,
		`"domain is ${tm_replace(global.str, "io", "com")}"`,
		`{}`,
		`10`,
		`"test"`,
		`[1, 2, 3]`,
		`a()`,
		`föo("föo") + föo`,
		`${var.name}`,
		`{ for k in var.val : k => k }`,
		`[ for k in var.val : k => k ]`,
	}

	for _, seed := range seedCorpus {
		f.Add(seed)
	}

	// END SEED OMIT

	globals := map[string]cty.Value{
		"str":  cty.StringVal("mineiros.io"),
		"bool": cty.BoolVal(true),
		"list": cty.ListVal([]cty.Value{
			cty.NumberVal(big.NewFloat(1)),
			cty.NumberVal(big.NewFloat(2)),
			cty.NumberVal(big.NewFloat(3)),
		}),
		"obj": cty.ObjectVal(map[string]cty.Value{
			"a": cty.StringVal("b"),
			"b": cty.StringVal("c"),
			"c": cty.StringVal("d"),
		}),
	}

	terramate := map[string]cty.Value{
		"path": cty.StringVal("/my/project"),
		"name": cty.StringVal("happy stack"),
	}

	f.Fuzz(func(t *testing.T, str string) {
		// WHY? because HCL uses the big.Float library for numbers and then
		// fuzzer can generate huge number strings like 100E101000000 that will
		// hang the process and eat all the memory....
		const bigNumRegex = "[\\d]+[\\s]*[.]?[\\s]*[\\d]*[EepP]{1}[\\s]*[+-]?[\\s]*[\\d]+"
		hasBigNumbers, _ := regexp.MatchString(bigNumRegex, str)
		if hasBigNumbers {
			return
		}

		const testattr = "attr"

		cfg := hcldoc(
			expr(testattr, str),
		)

		cfgString := cfg.String()
		parser := hclparse.NewParser()
		file, diags := parser.ParseHCL([]byte(cfgString), "fuzz")
		if diags.HasErrors() {
			return
		}

		body := file.Body.(*hclsyntax.Body)
		attr := body.Attributes[testattr]
		parsedExpr := attr.Expr

		exprRange := parsedExpr.Range()
		exprBytes := cfgString[exprRange.Start.Byte:exprRange.End.Byte]

		parsedTokens, diags := hclsyntax.LexExpression([]byte(exprBytes), "fuzz", hcl.Pos{})
		if diags.HasErrors() {
			return
		}

		ctx := NewContext("")
		assert.NoError(t, ctx.SetNamespace("globals", globals))
		assert.NoError(t, ctx.SetNamespace("terramate", terramate))

		want := toWriteTokens(parsedTokens)
		engine := newPartialEvalEngine(want, ctx)
		got, err := engine.Eval()

		if strings.Contains(cfgString, "global") ||
			strings.Contains(cfgString, "terramate") ||
			strings.Contains(cfgString, "tm_") {
			// TODO(katcipis): Validate generated code properties when
			// substitution is in play.
			return
		}

		assert.NoError(t, err)

		// Since we dont fuzz substitution/evaluation the tokens should be the same
		assert.EqualInts(t, len(want), len(got), "got %s != want %s", tokensStr(got), tokensStr(want))

		for i, gotToken := range got {
			wantToken := want[i]
			if diff := cmp.Diff(*gotToken, *wantToken); diff != "" {
				t.Errorf("got: %v", *gotToken)
				t.Errorf("want: %v", *wantToken)
				t.Error("diff:")
				t.Fatal(diff)
			}
		}
	})
}
