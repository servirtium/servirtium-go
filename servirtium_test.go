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
	s.servirtium.requestSequence = 0
	result := s.servirtium.appendContentInFile("currentContent", "newContent")
	s.Equal(result, "newContent")
}

func (s *ServirtiumTestSuite) TestAppendContentInFile_WhenSequenceOne() {
	s.servirtium.requestSequence = 1
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
		ResponseStatus:      "200",
	}
	s.servirtium.requestSequence = 0
	s.servirtium.record(params)
	s.NotNil(s.servirtium.content)
}

func (s *ServirtiumTestSuite) TestMaskBody() {
	inputContent := "<password>your password</password>"
	passwordRegex := regexp.MustCompile(`(<password>.{0,}<\/password>)`)
	result := maskBody(inputContent, map[*regexp.Regexp]string{passwordRegex: "<password>MASKED</password>"})
	s.Equal("<password>MASKED</password>", result)
}

func (s *ServirtiumTestSuite) TestRemoveHeader() {
	mockHeader := http.Header{"User-Agent": []string{"123"}, "Authorization": []string{"mockAuthorization"}}
	expected := http.Header{"User-Agent": []string{"123"}}
	result := removeHeader(mockHeader, []string{"Authorization"})
	s.Equal(expected, result)
}

func (s *ServirtiumTestSuite) TestMaskHeader() {
	mockHeader := http.Header{"User-Agent": []string{"123"}, "Authorization": []string{"mockAuthorization"}}
	expected := http.Header{"User-Agent": []string{"123"}, "Authorization": []string{"MASKED"}}
	result := replaceHeader(mockHeader, map[string]string{"Authorization": "MASKED"})
	s.Equal(expected, result)
}

func (s *ServirtiumTestSuite) TestDeleteRequestHeaders() {
	s.servirtium.DeleteRequestHeaders([]string{"Authorization"})
	expected := []string{"Authorization"}
	s.Equal(expected, s.servirtium.requestHeadersNeedDelete)
}

func (s *ServirtiumTestSuite) TestMaskRequestHeaders() {
	s.servirtium.MaskRequestHeaders(map[string]string{"Authorization": "123"})
	expected := map[string]string{"Authorization": "123"}
	s.Equal(expected, s.servirtium.requestHeadersNeedMask)
}

func (s *ServirtiumTestSuite) TestReplaceRequestHeaders() {
	s.servirtium.ReplaceRequestHeaders(map[string]string{"Authorization": "123"})
	expected := map[string]string{"Authorization": "123"}
	s.Equal(expected, s.servirtium.requestHeadersNeedReplace)
}

func (s *ServirtiumTestSuite) TestMaskRequestBody() {
	inputContent := "<password>your password</password>"
	passwordRegex := regexp.MustCompile(`(<password>.{0,}<\/password>)`)
	expected := map[*regexp.Regexp]string{passwordRegex: inputContent}
	s.servirtium.MaskRequestBody(map[*regexp.Regexp]string{passwordRegex: inputContent})
	s.Equal(expected, s.servirtium.requestBodyNeedMask)
}

func (s *ServirtiumTestSuite) TestDeleteResponseHeaders() {
	s.servirtium.DeleteResponseHeaders([]string{"Authorization"})
	expected := []string{"Authorization"}
	s.Equal(expected, s.servirtium.responseHeadersNeedDelete)
}

func (s *ServirtiumTestSuite) TestMaskResponseHeaders() {
	s.servirtium.MaskResponseHeaders(map[string]string{"Authorization": "123"})
	expected := map[string]string{"Authorization": "123"}
	s.Equal(expected, s.servirtium.responseHeadersNeedMask)
}

func (s *ServirtiumTestSuite) TestReplaceResponseHeaders() {
	s.servirtium.ReplaceResponseHeaders(map[string]string{"Authorization": "123"})
	expected := map[string]string{"Authorization": "123"}
	s.Equal(expected, s.servirtium.responseHeaderNeedReplace)
}

func (s *ServirtiumTestSuite) TestMaskResponseBody() {
	inputContent := "<password>your password</password>"
	passwordRegex := regexp.MustCompile(`(<password>.{0,}<\/password>)`)
	expected := map[*regexp.Regexp]string{passwordRegex: inputContent}
	s.servirtium.MaskResponseBody(map[*regexp.Regexp]string{passwordRegex: inputContent})
	s.Equal(expected, s.servirtium.responseBodyNeedMask)
}
