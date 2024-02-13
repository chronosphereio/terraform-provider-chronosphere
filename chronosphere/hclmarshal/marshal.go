package hclmarshal

import (
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
	hcl2 "github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
	"go.uber.org/multierr"
)

// Marshaler is a HCL marshaler.
type Marshaler struct {
	file *hclwrite.File
	root *hclwrite.Body
	errs error
}

const blockIndent = "  "

// New creates a HCL marshaler.
func New() *Marshaler {
	f := hclwrite.NewEmptyFile()
	return &Marshaler{
		file: f,
		root: f.Body(),
	}
}

// MarshalTo marshals the added resources to the specified writer.
func (m *Marshaler) MarshalTo(w io.Writer) error {
	if m.errs != nil {
		return m.errs
	}

	_, err := m.file.WriteTo(w)
	return err
}

func (m *Marshaler) addBlock(blockType, resourceType, id string) *Block {
	m.root.AppendNewline()
	b := m.root.AppendNewBlock(blockType, []string{resourceType, id})
	return &Block{
		body:         b.Body(),
		nesting:      1,
		errContainer: &m.errs,
	}
}

// AddData adds a top-level "data" block to the file.
func (m *Marshaler) AddData(resourceType string, id string) *Block {
	return m.addBlock("data", resourceType, id)
}

// AddResource adds a top-level "resource" block to the file.
func (m *Marshaler) AddResource(resourceType string, id string) *Block {
	return m.addBlock("resource", resourceType, id)
}

// Block represents a Terraform block, with helpers for adding fields.
type Block struct {
	nesting int
	body    *hclwrite.Body

	// prevFieldComplex holds whether the last field within this block is
	// currently complex (struct or map type), used to determine whether newlines
	// need to be added between consecutive fields.
	prevFieldComplex bool

	// Pointer to the top-level Marshaler.errs.
	errContainer *error
}

func (b *Block) addErr(err error) {
	*b.errContainer = multierr.Append(*b.errContainer, err)
}

// AddBlock adds a nested block.
func (b *Block) AddBlock(name string) *Block {
	block := b.body.AppendNewBlock(name, nil /* labels */)
	return &Block{
		body:         block.Body(),
		nesting:      b.nesting + 1,
		errContainer: b.errContainer,
	}
}

// AddNewLine adds a new line to the block, it's typically used
// before adding nested blocks, or to separate groups of fields.
func (b *Block) AddNewLine() {
	if len(b.body.Attributes()) == 0 && len(b.body.Blocks()) == 0 {
		return
	}

	b.body.AppendNewline()
}

// AddRef adds a tfid.ID, which may be a hardcoded slug, or a local TF ref.
// If the ID is empty, then no field is added.
func (b *Block) AddRef(name string, id tfid.ID) {
	switch id.Type() {
	case tfid.TypeEmpty:
		// No reference set, skip the field.
	case tfid.TypeSlug:
		b.Add(name, id.Slug())
	case tfid.TypeLocalRef:
		b.AddTFRef(name, id.LocalRef())
	default:
		panic(fmt.Errorf("failed to add tfid.ID to %q, unknown ID type %v", name, id.Type()))
	}
}

// AddRefs adds a list of tfid.ID, supporting mixed lists of hardcoded slugs
// or local TF refs.
// If an ID is empty, it is not added to the list.
// If the list is empty, then no field is added.
func (b *Block) AddRefs(name string, ids []tfid.ID) {
	if len(ids) == 0 {
		return
	}

	var tokens hclwrite.Tokens

	tokens = append(tokens, &hclwrite.Token{
		Type:  hclsyntax.TokenOBrack,
		Bytes: []byte("["),
	})

	var listHasElements bool
	for _, id := range ids {
		// Ignore empty fields, before we check for commas.
		if id.Type() == tfid.TypeEmpty {
			continue
		}

		if listHasElements {
			tokens = append(tokens, &hclwrite.Token{
				Type:  hclsyntax.TokenComma,
				Bytes: []byte(","),
			})
		}

		listHasElements = true
		switch id.Type() {
		case tfid.TypeSlug:
			v := cty.StringVal(id.Slug())
			tokens = append(tokens, hclwrite.TokensForValue(v)...)
		case tfid.TypeLocalRef:
			traversal := refToTraversal(id.LocalRef())
			expr := hclwrite.NewExpressionAbsTraversal(traversal)
			tokens = expr.BuildTokens(tokens)
		}
	}

	tokens = append(tokens, &hclwrite.Token{
		Type:  hclsyntax.TokenCBrack,
		Bytes: []byte("]"),
	})

	if !listHasElements {
		return
	}

	b.formatNewLines(false /* nextFieldComplex */)
	b.body.SetAttributeRaw(name, tokens)
}

func refToParts(ref tfid.Ref) []string {
	parts := make([]string, 0, 4)
	if ref.Datasource {
		parts = append(parts, "data")
	}
	parts = append(parts, ref.Type, ref.ID)
	if ref.Field != "" {
		parts = append(parts, ref.Field)
	} else {
		parts = append(parts, "id")
	}
	return parts
}

func refToTraversal(ref tfid.Ref) hcl2.Traversal {
	parts := refToParts(ref)
	traversal := make(hcl2.Traversal, 0, len(parts))
	for i, p := range refToParts(ref) {
		if i == 0 {
			traversal = append(traversal, hcl2.TraverseRoot{Name: p})
		} else {
			traversal = append(traversal, hcl2.TraverseAttr{Name: p})
		}
	}
	return traversal
}

// AddTFRef adds a Terraform reference.
func (b *Block) AddTFRef(name string, ref tfid.Ref) {
	b.body.SetAttributeTraversal(name, refToTraversal(ref))
}

// Add a field that is always marshalled and set to the result of a TF function call.
func (b *Block) AddFuncCall(name string, funcName string, args ...any) {
	var argTokens []hclwrite.Tokens
	for _, arg := range args {
		ctyV, _, err := ctyValue(arg)
		if err != nil {
			b.addErr(err)
			return
		}

		tokens := hclwrite.TokensForValue(ctyV)
		argTokens = append(argTokens, tokens)
	}

	b.body.SetAttributeRaw(name, hclwrite.TokensForFunctionCall(funcName, argTokens...))
}

// Add a field to the block that is always marshalled.
func (b *Block) Add(name string, v any) {
	ctyV, resInfo, err := ctyValue(v)
	if err != nil {
		b.addErr(err)
		return
	}

	b.addCtyVal(name, ctyV, resInfo)
}

// Add a field to the block that is only marshaled if the value
// is not empty.
func (b *Block) AddOptional(name string, v any) {
	ctyV, resInfo, err := ctyValue(v)
	if err != nil {
		b.addErr(err)
		return
	}

	if resInfo.zeroOrEmpty {
		return
	}

	b.addCtyVal(name, ctyV, resInfo)
}

// formatNewLines controls newline behavior between fields:
//  1. if the previous field or the next field is complex, add a preceding new line
//  2. if the next field is complex, set a flag that will ensure a succeeding newline
//     will be added if another field is marshalled
func (b *Block) formatNewLines(nextFieldComplex bool) {
	if b.prevFieldComplex || nextFieldComplex {
		b.AddNewLine()
	}

	b.prevFieldComplex = nextFieldComplex
}

func (b *Block) addCtyVal(name string, v cty.Value, resInfo ctyResultInfo) {
	// ensure we have whitespace before and after if the next field is a map type
	b.formatNewLines(resInfo.mapType)
	if v.Type() == cty.String && strings.Contains(v.AsString(), "\n") {
		b.addMultilineString(name, v.AsString())
		return
	}

	b.body.SetAttributeValue(name, v)
}

// Multi-line strings use heredoc strings where possible for better readability.
// hclwrite doesn't support this natively, though there is a TODO:
// https://github.com/krisskross/hcl2/blob/3e4b7e0e/hclwrite/generate.go#L61
// To make heredocs format currently with the appropriate indentation, we track
// the indentation (level of nested blocks) manually.
func (b *Block) addMultilineString(name string, v string) {
	const heredocMarker = "EOF"

	if strings.Contains(v, heredocMarker) {
		b.body.SetAttributeValue(name, cty.StringVal(v))
		return
	}

	tokens := hclwrite.Tokens{
		{
			Type:  hclsyntax.TokenOHeredoc,
			Bytes: []byte("<<-" + heredocMarker),
		},
		{
			Type:  hclsyntax.TokenNewline,
			Bytes: []byte("\n"),
		},
	}

	// We strip at most one trailing newline from the original string. This
	// trailing newline will be added back manually below, since heredoc strings
	// always require a trailing newline. If the original string did not have a
	// trailing newline, we'll strip the additional one via chomp.
	lines := strings.Split(strings.TrimSuffix(v, "\n"), "\n")

	// We add tokens for each line, handling indentation + newlines manually.
	for i, line := range lines {
		toAdd := line
		if i > 0 {
			// The first line is indented correctly, so we indent all other lines.
			toAdd = strings.Repeat(blockIndent, b.nesting) + line
		}
		if i < len(lines)-1 {
			// We add a new line token after the loop, so don't add it to the string.
			toAdd += "\n"
		}

		tokens = append(tokens, &hclwrite.Token{
			Type:  hclsyntax.TokenStringLit,
			Bytes: []byte(toAdd),
		})
	}

	tokens = append(
		tokens,
		&hclwrite.Token{
			Type:  hclsyntax.TokenNewline,
			Bytes: []byte("\n"),
		},
		&hclwrite.Token{
			Type:  hclsyntax.TokenOHeredoc,
			Bytes: []byte(heredocMarker),
		},
	)

	if !strings.HasSuffix(v, "\n") {
		tokens = wrapMultilineWithChomp(v, tokens)
	}

	b.body.SetAttributeRaw(name, tokens)
}

// wrapMultilineWithChomp wraps a multiline heredoc string with chomp. TF requires that heredoc markers are on their
// own line, so if the raw string didn't have a trailing newline we have to wrap everything in chomp(...) to strip
// the one we add above.
func wrapMultilineWithChomp(v string, tokens hclwrite.Tokens) hclwrite.Tokens {
	// prepend chomp(
	wrappedTokens := append(
		hclwrite.Tokens{
			&hclwrite.Token{
				// TokenStringLit does not have padding automatically prefixed, so we have to manually
				// add a space before chomp literal. We also intentionally avoid using hclsyntax.TokenOParen
				// since that causes additional indents for the rest of the multiline string.
				Type:  hclsyntax.TokenStringLit,
				Bytes: []byte(" chomp("),
			},
		},
		tokens...,
	)

	// append \n)
	wrappedTokens = append(
		wrappedTokens,
		&hclwrite.Token{
			Type:  hclsyntax.TokenNewline,
			Bytes: []byte("\n"),
		},
		&hclwrite.Token{
			Type:  hclsyntax.TokenStringLit,
			Bytes: []byte(")"),
		},
	)

	return wrappedTokens
}

type ctyResultInfo struct {
	zeroOrEmpty bool
	mapType     bool
}

func ctyValue(v any) (cty.Value, ctyResultInfo, error) {
	impliedType, err := gocty.ImpliedType(v)
	if err != nil {
		return cty.NilVal, ctyResultInfo{}, err
	}

	resInfo := ctyResultInfo{
		zeroOrEmpty: isEmpty(v),
		mapType:     impliedType.IsMapType(),
	}

	ctyV, err := gocty.ToCtyValue(v, impliedType)
	return ctyV, resInfo, err
}

func isEmpty(v any) bool {
	rv := reflect.ValueOf(v)
	if rv.IsZero() {
		return true
	}

	if k := rv.Kind(); k == reflect.Map || k == reflect.Slice {
		return rv.Len() == 0
	}

	return false
}
