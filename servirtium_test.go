package servirtium

import (
	"fmt"
	"log"
	"os"
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

func (s *ServirtiumTestSuite) TestStartPlayback() {
	s.servirtium.StartPlayback("mockName")
	s.NotNil(s.servirtium.ServerPlayback)
	s.servirtium.EndPlayback()
}

func (s *ServirtiumTestSuite) TestInitServerPlayback() {
	s.servirtium.initServerPlayback("mockName")
	s.NotNil(s.servirtium.ServerPlayback)
	s.servirtium.EndPlayback()
}

func (s *ServirtiumTestSuite) TestAnualAvgHandlerPlayback() {
	result := s.servirtium.anualAvgHandlerPlayback("mockName")
	s.NotNil(result)
}

func (s *ServirtiumTestSuite) TestStartRecord() {
	s.servirtium.StartRecord("https://google.com")
	s.NotNil(s.servirtium.ServerRecord)
	s.servirtium.EndRecord()
}

func (s *ServirtiumTestSuite) TestInitRecordServer() {
	s.servirtium.initRecordServer("https://google.com")
	s.NotNil(s.servirtium.ServerRecord)
	s.servirtium.EndRecord()
}

func (s *ServirtiumTestSuite) TestManInTheMiddleHandler() {
	result := s.servirtium.manInTheMiddleHandler("https://google.com")
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
	s.servirtium.RequestSequence = 0
	result := s.servirtium.appendContentInFile("currentContent", "newContent")
	s.Equal(result, "newContent")
}

func (s *ServirtiumTestSuite) TestAppendContentInFile_WhenSequenceOne() {
	s.servirtium.RequestSequence = 1
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
	s.servirtium.RequestSequence = 0
	s.servirtium.record(params)
	s.NotNil(s.servirtium.Content)
}
