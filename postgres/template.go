package postgres

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"text/template"
)

var ErrTemplate = errors.New("postgres template")
var ErrTemplateFuncNotAvail = fmt.Errorf("%w: function not available in this execution context", ErrTemplate)
var ErrJoinNotStarted = fmt.Errorf("%w: join not started", ErrTemplate)
var ErrJoinNotEnded = fmt.Errorf("%w: join started but not ended", ErrTemplate)

type Template struct {
	t *template.Template
}

func Parse(text string) (*Template, error) {
	// Generate stub functions for parsing. These will be replaced with
	// contextualized versions when the template is executed.
	funcs := templateFuncs(
		nil,   // args
		false, // counting
		nil,   // firstItemIndex
		nil,   // joinFrames
		nil,   // pageSize
	)

	// Parse the template.
	t, err := template.New("postgres").Funcs(funcs).Parse(text)
	if err != nil {
		return nil, err
	}

	// Success.
	return &Template{t}, nil
}

func MustParse(text string) *Template {
	t, err := Parse(text)
	if err != nil {
		panic(err)
	}

	return t
}

func (t *Template) Execute(data interface{}) (string, []any, error) {
	return t.execute(data, false, nil, nil)
}

func (t *Template) ExecuteCount(data interface{}) (string, []any, error) {
	return t.execute(data, true, nil, nil)
}

func (t *Template) ExecuteList(data interface{}, firstItemIndex int64, pageSize int64) (string, []any, error) {
	return t.execute(data, false, &firstItemIndex, &pageSize)
}

func (t *Template) execute(
	data interface{},
	counting bool,
	firstItemIndex *int64,
	pageSize *int64,
) (string, []any, error) {
	localTemplate, err := t.t.Clone()
	if err != nil {
		return "", nil, err
	}

	var args []any
	var buf strings.Builder
	var joinFrames []joinFrame
	if err := localTemplate.Funcs(templateFuncs(
		&args,
		counting,
		firstItemIndex,
		&joinFrames,
		pageSize,
	)).Execute(&buf, data); err != nil {
		return "", nil, err
	}

	if len(joinFrames) > 0 {
		return "", nil, ErrJoinNotEnded
	}

	return buf.String(), args, nil
}

func templateFuncs(
	args *[]any,
	counting bool,
	firstItemIndex *int64,
	joinFrames *[]joinFrame,
	pageSize *int64,
) template.FuncMap {
	return template.FuncMap{
		"arg":            templateFuncArg(args),
		"counting":       templateFuncCounting(counting),
		"endJoin":        templateFuncEndJoin(joinFrames),
		"firstItemIndex": templateFuncFirstItemIndex(firstItemIndex, args),
		"join":           templateFuncJoin(joinFrames),
		"pageSize":       templateFuncPageSize(pageSize, args),
		"sep":            templateFuncSep(joinFrames),
	}
}

type joinFrame struct {
	sep   string
	count int
}

func templateFuncArg(args *[]any) func(arg any) string {
	return func(arg any) string {
		*args = append(*args, arg)
		return "$" + strconv.Itoa(len(*args))
	}
}

func templateFuncCounting(counting bool) func() bool {
	return func() bool {
		return counting
	}
}

func templateFuncEndJoin(fames *[]joinFrame) func() (string, error) {
	return func() (string, error) {
		if len(*fames) == 0 {
			return "", ErrJoinNotStarted
		}

		*fames = (*fames)[:len(*fames)-1]
		return "", nil
	}
}

func templateFuncFirstItemIndex(firstItemIndex *int64, args *[]any) func() (string, error) {
	return func() (string, error) {
		if firstItemIndex == nil {
			return "", fmt.Errorf("%w: firstItemIndex", ErrTemplateFuncNotAvail)
		}

		*args = append(*args, *firstItemIndex)
		return "$" + strconv.Itoa(len(*args)), nil
	}
}

func templateFuncJoin(frames *[]joinFrame) func(sep string) string {
	return func(sep string) string {
		*frames = append(*frames, joinFrame{sep: sep})
		return ""
	}
}

func templateFuncPageSize(pageSize *int64, args *[]any) func() (string, error) {
	return func() (string, error) {
		if pageSize == nil {
			return "", fmt.Errorf("%w: pageSize", ErrTemplateFuncNotAvail)
		}

		*args = append(*args, *pageSize)
		return "$" + strconv.Itoa(len(*args)), nil
	}
}

func templateFuncSep(frames *[]joinFrame) func() (string, error) {
	return func() (string, error) {
		if len(*frames) == 0 {
			return "", fmt.Errorf("%w: sep", ErrTemplateFuncNotAvail)
		}

		frame := &(*frames)[len(*frames)-1]
		frame.count++
		if frame.count > 1 {
			return frame.sep, nil
		}

		return "", nil
	}
}
