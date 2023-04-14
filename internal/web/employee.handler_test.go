package web_test

import (
	"errors"
	"fmt"
	"golang-api/internal/core"
	"golang-api/internal/web"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	ifaker "syreclabs.com/go/faker"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type EmployeeHandlerSuite struct {
	suite.Suite
	path                string
	target              *httpexpect.Expect
	employeeServiceMock *core.MockIEmployeeService
}

// The SetupTest method will be run before every test in the suite
func (suite *EmployeeHandlerSuite) SetupTest() {
	suite.path = "/"
	suite.employeeServiceMock = core.NewMockIEmployeeService(suite.T())
	engine := gin.New()
	handler := web.NewEmployeeHandler(suite.employeeServiceMock)
	engine.GET(suite.path, handler.GetEmployees)
	engine.POST(suite.path, handler.PostEmployee)
	suite.target = httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(engine),
			Jar:       httpexpect.NewCookieJar(),
		},
		Reporter: httpexpect.NewAssertReporter(suite.T()),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(suite.T(), true),
		},
	})
}

func (suite *EmployeeHandlerSuite) TestGetEmployees() {
	tests := []struct {
		desc         string
		isError      error
		expected     []core.Employee
		actualStatus int
	}{
		{
			desc:         "success",
			actualStatus: http.StatusOK,
		},
		{
			desc:         "db error",
			isError:      errors.New("test"),
			actualStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.desc, func(t *testing.T) {
			call := suite.employeeServiceMock.On("RetrieveAllEmployees", mock.Anything).Once()
			if tt.isError != nil {
				call.Return(nil, tt.isError)
			} else {
				faker.FakeData(&tt.expected, options.WithRandomMapAndSliceMaxSize(20))
				call.Return(tt.expected, nil)
			}

			res := suite.target.GET(suite.path).Expect()
			if tt.isError == nil {
				var actual []core.Employee
				res.JSON().Decode(&actual)
				assert.Equal(suite.T(), tt.expected, actual)
			}
			res.Status(tt.actualStatus)

			suite.employeeServiceMock.AssertExpectations(suite.T())
		})
	}
}

func (suite *EmployeeHandlerSuite) TestPostEmployee() {
	tests := []struct {
		desc          string
		inputGender   string
		internalError error
		expected      int
		actualStatus  int
	}{
		{
			desc:         "success",
			inputGender:  ifaker.RandomString(20),
			actualStatus: http.StatusOK,
			expected:     ifaker.RandomInt(0, 10),
		},
		{
			desc:          "db error",
			inputGender:   ifaker.RandomString(20),
			internalError: errors.New("test"),
			actualStatus:  http.StatusInternalServerError,
		},
		{
			desc:         "invalid gender format",
			inputGender:  fmt.Sprintf("%s$", ifaker.RandomString(20)),
			actualStatus: http.StatusBadRequest,
			expected:     -1,
		},
		{
			desc:         "gender not found",
			actualStatus: http.StatusBadRequest,
			expected:     -1,
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.desc, func(t *testing.T) {
			if tt.expected > -1 {
				call := suite.employeeServiceMock.On("AddEmployee", mock.Anything, tt.inputGender)
				if tt.internalError != nil {
					call.Return(-1, tt.internalError)
				} else {
					call.Return(tt.expected, nil)
				}
			}

			res := suite.target.POST(suite.path).WithJSON(web.Employee{Gender: tt.inputGender}).Expect()
			if tt.internalError == nil && tt.expected > -1 {
				res.JSON().IsEqual(tt.expected)
			}
			res.Status(tt.actualStatus)

			if tt.expected > -1 {
				suite.employeeServiceMock.AssertExpectations(suite.T())
			}
		})
	}
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestEmployeeHandlerSuite(t *testing.T) {
	suite.Run(t, new(EmployeeHandlerSuite))
}
