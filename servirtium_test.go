package servirtium

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ServirtiumTestSuite struct {
	servirtium *Impl
	suite.Suite
}

func TestServirtiumTestSuite(t *testing.T) {
	suite.Run(t, new(ServirtiumTestSuite))
}

func (s *ServirtiumTestSuite) SetupTest() {
	servirtium := NewServirtium()
	s.servirtium = servirtium
}

func (s *ServirtiumTestSuite) TestInitServerPlayback() {
	s.servirtium.initServerPlayback("mockName")
	s.NotNil(s.servirtium.ServerPlayback)
	s.servirtium.EndPlayback()
}

func (s *ServirtiumTestSuite) TestAnualAvgHandlerPlayback() {
	result := s.servirtium.playbackHandler("mockName")
	s.NotNil(result)
}

func (s *ServirtiumTestSuite) TestInitRecordServer() {
	s.servirtium.initRecordServer("https://google.com")
	s.NotNil(s.servirtium.ServerRecord)
	s.servirtium.EndRecord()
}

func (s *ServirtiumTestSuite) TestManInTheMiddleHandler() {
	result := s.servirtium.recordHandler("https://google.com")
	s.NotNil(result)
}

func (s *ServirtiumTestSuite) TestCheckMarkdownExist_True() {
	workingPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	result := s.servirtium.checkMarkdownExists(fmt.Sprintf("%s/servirtium.go", workingPath))
	s.True(result)
}

func (s *ServirtiumTestSuite) TestCheckMarkdownExist_False() {
	result := s.servirtium.checkMarkdownExists("./servitium1.go")
	s.False(result)
}

func (s *ServirtiumTestSuite) TestAppendContentInFile_WhenSequenceZero() {
	s.servirtium.interactionSequence = 0
	result := s.servirtium.appendContentInFile("currentContent", "newContent")
	s.Equal(result, "newContent")
}

func (s *ServirtiumTestSuite) TestAppendContentInFile_WhenSequenceOne() {
	s.servirtium.interactionSequence = 1
	result := s.servirtium.appendContentInFile("currentContent", "newContent")
	s.Equal("currentContent\nnewContent", result)
}

func (s *ServirtiumTestSuite) TestRecord() {
	params := recordData{
		RequestMethod:       "GET",
		RequestURLPath:      "http://google.com",
		RequestHeader:       nil,
		RequestBody:         "request",
		ResponseContentType: "application/json",
		ResponseBody:        "response",
		ResponseHeader:      nil,
		ResponseStatus:      200,
	}
	s.servirtium.interactionSequence = 0
	s.servirtium.record(params)
	s.NotNil(s.servirtium.content)
}

func (s *ServirtiumTestSuite) TestReplaceContent() {
	inputContent := "<password>your password</password>"
	passwordRegex := regexp.MustCompile(`(<password>.{0,}<\/password>)`)
	result := replaceContent(inputContent, map[*regexp.Regexp]string{passwordRegex: "<password>MASKED</password>"})
	s.Equal("<password>MASKED</password>", result)
}

func (s *ServirtiumTestSuite) TestRemoveHeaders() {
	mockHeader := http.Header{"User-Agent": []string{"123"}, "Authorization": []string{"mockAuthorization"}}
	expected := http.Header{"User-Agent": []string{"123"}}
	result := removeHeaders(mockHeader, []string{"Authorization"})
	s.Equal(expected, result)
}

func (s *ServirtiumTestSuite) TestReplaceHeaders() {
	mockHeader := http.Header{"User-Agent": []string{"123"}, "Authorization": []string{"mockAuthorization"}}
	expected := http.Header{"User-Agent": []string{"123"}, "Authorization": []string{"MASKED"}}
	authorizationRegex := regexp.MustCompile(`(Authorization: .{0,})`)
	result := replaceHeaders(mockHeader, map[*regexp.Regexp]string{authorizationRegex: "Authorization: MASKED"})
	s.Equal(expected, result)
}

func (s *ServirtiumTestSuite) TestRemoveHeadersPlayback() {
	mockHeader := map[string]string{"User-Agent": "123", "Authorization": "mockAuthorization"}
	expected := map[string]string{"User-Agent": "123"}
	result := removeHeadersPlayback(mockHeader, []string{"Authorization"})
	s.Equal(expected, result)
}

func (s *ServirtiumTestSuite) TestReplaceHeadersPlayback() {
	mockHeader := map[string]string{"User-Agent": "123", "Authorization": "mockAuthorization"}
	expected := map[string]string{"User-Agent": "123", "Authorization": "MASKED"}
	authorizationRegex := regexp.MustCompile(`(Authorization: .{0,})`)
	result := replaceHeadersPlayback(mockHeader, map[*regexp.Regexp]string{authorizationRegex: "Authorization: MASKED"})
	s.Equal(expected, result)
}

func (s *ServirtiumTestSuite) TestSetCallerRequestHeadersRemoval() {
	s.servirtium.SetCallerRequestHeadersRemoval([]string{"Authorization"})
	expected := []string{"Authorization"}
	s.Equal(expected, s.servirtium.callerRequestHeadersRemoval)
}

func (s *ServirtiumTestSuite) TestSetCallerRequestHeaderReplacements() {
	authorizationRegex := regexp.MustCompile(`(Authorization: .{0,})`)
	s.servirtium.SetCallerRequestHeaderReplacements(map[*regexp.Regexp]string{authorizationRegex: "123"})
	expected := map[*regexp.Regexp]string{authorizationRegex: "123"}
	s.Equal(expected, s.servirtium.callerRequestHeaderReplacements)
}

func (s *ServirtiumTestSuite) TestSetCallerRequestBodyReplacement() {
	inputContent := "<password>your password</password>"
	passwordRegex := regexp.MustCompile(`(<password>.{0,}<\/password>)`)
	expected := map[*regexp.Regexp]string{passwordRegex: inputContent}
	s.servirtium.SetCallerRequestBodyReplacement(map[*regexp.Regexp]string{passwordRegex: inputContent})
	s.Equal(expected, s.servirtium.callerRequestBodyReplacement)
}

func (s *ServirtiumTestSuite) TestSetRecordRequestHeadersRemoval() {
	s.servirtium.SetRecordRequestHeadersRemoval([]string{"Authorization"})
	expected := []string{"Authorization"}
	s.Equal(expected, s.servirtium.recordRequestHeadersRemoval)
}

func (s *ServirtiumTestSuite) TestSetRecordRequestHeaderReplacements() {
	authorizationRegex := regexp.MustCompile(`(Authorization: .{0,})`)
	s.servirtium.SetRecordRequestHeaderReplacements(map[*regexp.Regexp]string{authorizationRegex: "123"})
	expected := map[*regexp.Regexp]string{authorizationRegex: "123"}
	s.Equal(expected, s.servirtium.recordRequestHeaderReplacements)
}

func (s *ServirtiumTestSuite) TestSetRecordRequestBodyReplacement() {
	inputContent := "<password>your password</password>"
	passwordRegex := regexp.MustCompile(`(<password>.{0,}<\/password>)`)
	expected := map[*regexp.Regexp]string{passwordRegex: inputContent}
	s.servirtium.SetRecordRequestBodyReplacement(map[*regexp.Regexp]string{passwordRegex: inputContent})
	s.Equal(expected, s.servirtium.recordRequestBodyReplacement)
}

func (s *ServirtiumTestSuite) TestSetCallerResponseHeadersRemoval() {
	s.servirtium.SetCallerResponseHeadersRemoval([]string{"Authorization"})
	expected := []string{"Authorization"}
	s.Equal(expected, s.servirtium.callerResponseHeadersRemoval)
}

func (s *ServirtiumTestSuite) TestSetCallerResponseHeaderReplacements() {
	authorizationRegex := regexp.MustCompile(`(Authorization: .{0,})`)
	s.servirtium.SetCallerResponseHeaderReplacements(map[*regexp.Regexp]string{authorizationRegex: "123"})
	expected := map[*regexp.Regexp]string{authorizationRegex: "123"}
	s.Equal(expected, s.servirtium.callerResponseHeaderReplacements)
}

func (s *ServirtiumTestSuite) TestSetCallerResponseBodyReplacement() {
	inputContent := "<password>your password</password>"
	passwordRegex := regexp.MustCompile(`(<password>.{0,}<\/password>)`)
	expected := map[*regexp.Regexp]string{passwordRegex: inputContent}
	s.servirtium.SetCallerResponseBodyReplacement(map[*regexp.Regexp]string{passwordRegex: inputContent})
	s.Equal(expected, s.servirtium.callerResponseBodyReplacement)
}

func (s *ServirtiumTestSuite) TestSetRecordResponseHeadersRemoval() {
	s.servirtium.SetRecordResponseHeadersRemoval([]string{"Authorization"})
	expected := []string{"Authorization"}
	s.Equal(expected, s.servirtium.recordResponseHeadersRemoval)
}

func (s *ServirtiumTestSuite) TestSetRecordResponseHeaderReplacements() {
	authorizationRegex := regexp.MustCompile(`(Authorization: .{0,})`)
	s.servirtium.SetRecordResponseHeaderReplacements(map[*regexp.Regexp]string{authorizationRegex: "123"})
	expected := map[*regexp.Regexp]string{authorizationRegex: "123"}
	s.Equal(expected, s.servirtium.recordResponseHeaderReplacements)
}

func (s *ServirtiumTestSuite) TestSetRecordResponseBodyReplacement() {
	inputContent := "<password>your password</password>"
	passwordRegex := regexp.MustCompile(`(<password>.{0,}<\/password>)`)
	expected := map[*regexp.Regexp]string{passwordRegex: inputContent}
	s.servirtium.SetRecordResponseBodyReplacement(map[*regexp.Regexp]string{passwordRegex: inputContent})
	s.Equal(expected, s.servirtium.recordResponseBodyReplacement)
}
