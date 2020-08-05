# servirtium-go
A Go implementation of Servirtium, a library that helps test interactions with APIs.

# How it works
Servirtium is a server that serves as a man-in-the-middle: it processes incoming requests, forwards them to a destination API and writes the response into a Markdown file with a special format that is common across all of the implementations of the library. Later these Markdown files are used to replay the interactions that were recorded before, allowing to test interactions without making real API calls.

# Prerequisites
The library is written in the Go programming language, so to use or compile the library you need Go Compiler. You can download it on the official site.

You can get the servirtium-go package by
```
# use Git tags
go get https://github.com/servirtium/servirtium-go@v4.0.1
# or Git branch name
go get https://github.com/servirtium/servirtium-go@master
# or Git commit hash
go get https://github.com/servirtium/servirtium-go@08c92af
```

# Usage
You can use one of these attribute on test function to start the Servirtium server and config it to serve request in records and playback mode

## For testify record mode example:

```
type ClimateTestSuiteRecord struct {
	recordClient ClientImpl
	suite.Suite
	servirtium *servirtium.Impl
}

func TestClimateTestSuiteRecord(t *testing.T) {
	suite.Run(t, new(ClimateTestSuiteRecord))
}

func (s *ClimateTestSuiteRecord) BeforeTest(suiteName, testName string) {
	validate := validator.New()
	s.servirtium = servirtium.NewServirtium()

	// For delete headers no need from request
	s.servirtium.DeleteRequestHeaders([]string{"No-need-header"})

	// For over write data in request header
	s.servirtium.ReplaceRequestHeaders(map[string]string{"User-Agent": "Servirtium-Testing"})

	// For mask sensitive data in request header
	s.servirtium.MaskResponseHeaders(map[string]string{"Set-Cookie": "REPLACED-IN-RECORDING", "Date": "Tue, 04 Aug 2020 16:53:25 GMT"})

	// For mask sensitive data in the request body
	passwordRegex := regexp.MustCompile(`(<password>.{0,}<\/password>)`)
	s.servirtium.MaskRequestBody(map[*regexp.Regexp]string{passwordRegex: "<password>MASKED</password>"})


	// For delete headers no need from response
	s.servirtium.DeleteResponseHeaders([]string{"No-need-header"})

	// For over write data in request header
	s.servirtium.ReplaceResponseHeaders(map[string]string{"User-Agent": "Servirtium-Testing"})

	// For mask sensitive data in request header
	s.servirtium.MaskResponseHeaders(map[string]string{"Set-Cookie": "REPLACED-IN-RECORDING", "Date": "Tue, 04 Aug 2020 16:53:25 GMT"})

	// For mask sensitive data in the response body
	tokenRegex := regexp.MustCompile(`(<token>.{0,}<\/token>)`)
	s.servirtium.MaskResponseBody(map[*regexp.Regexp]string{tokenRegex: "<password>MASKED</password>"})

	s.servirtium.StartRecord("http://climatedataapi.worldbank.org")
	recordClient := NewClient(http.DefaultClient, validate, s.servirtium.ServerRecord.URL)
	s.recordClient = *recordClient
}

func (s *ClimateTestSuiteRecord) AfterTest(suite, testName string) {
	isMatch := s.servirtium.CheckMarkdownIsDifferentToPreviousRecording(testName)
	s.True(isMatch)
	s.servirtium.WriteRecord(testName)
	s.servirtium.EndRecord()
}

func (s *ClimateTestSuiteRecord) TestAverageRainfallForGreatBritainFrom1980to1999Exists() {
	var (
		ctx      = context.Background()
		expected = float64(988.8454972331014)
	)
	recordResult, recordErr := s.recordClient.GetAveAnnualRainfall(ctx, 1980, 1999, "gbr")
	s.Equal(expected, recordResult)
	s.Nil(recordErr)
}

```

## For testify Playback mode example:
```
type ClimateTestSuitePlayback struct {
	playbackClient ClientImpl
	suite.Suite
	servirtium *servirtium.Impl
}

func TestClimateTestSuitePlayback(t *testing.T) {
	suite.Run(t, new(ClimateTestSuitePlayback))
}

func (s *ClimateTestSuitePlayback) BeforeTest(suiteName, testName string) {
	validate := validator.New()
	servirtium := servirtium.NewServirtium()
	s.servirtium = servirtium
	s.servirtium.StartPlayback(testName)
	playbackClient := NewClient(http.DefaultClient, validate, s.servirtium.ServerPlayback.URL)
	s.playbackClient = *playbackClient
}

func (s *ClimateTestSuitePlayback) AfterTest(suite, testName string) {
	s.servirtium.EndPlayback()
}

func (s *ClimateTestSuitePlayback) TestAverageRainfallForGreatBritainFrom1980to1999Exists() {
	var (
		ctx      = context.Background()
		expected = float64(988.8454972331014)
	)
	playbackResult, playbackErr := s.playbackClient.GetAveAnnualRainfall(ctx, 1980, 1999, "gbr")
	s.Equal(expected, playbackResult)
	s.Nil(playbackErr)
}
```

## For testify direct mode example:
```
type ClimateTestSuiteDirect struct {
	directClient ClientImpl
	suite.Suite
	servirtium *servirtium.Impl
}

func TestClimateTestSuiteDirect(t *testing.T) {
	suite.Run(t, new(ClimateTestSuiteDirect))
}

func (s *ClimateTestSuiteDirect) SetupTest() {
	validate := validator.New()
	directClient := NewClient(http.DefaultClient, validate, "http://climatedataapi.worldbank.org")
	s.directClient = *directClient
}

func (s *ClimateTestSuiteDirect) TestAverageRainfallForGreatBritainFrom1980to1999Exists() {
	var (
		ctx      = context.Background()
		expected = float64(988.8454972331014)
	)
	directResult, directErr := s.directClient.GetAveAnnualRainfall(ctx, 1980, 1999, "gbr")
	s.Equal(expected, directResult)
	s.Nil(directErr)
}
```