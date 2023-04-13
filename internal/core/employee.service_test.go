package core_test

import (
	"context"
	"errors"
	"golang-api/internal/core"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type EmployeeServiceSuite struct {
	suite.Suite
	target                 *core.EmployeeService
	employeeRepositoryMock *core.MockIEmployeeRepository
}

// The SetupTest method will be run before every test in the suite
func (suite *EmployeeServiceSuite) SetupTest() {
	suite.employeeRepositoryMock = core.NewMockIEmployeeRepository(suite.T())
	suite.target = core.NewEmployeeService(suite.employeeRepositoryMock)
}

func (suite *EmployeeServiceSuite) TestRetrieveAllEmployees() {
	tests := []struct {
		desc              string
		getEmployeesError error
		expected          []core.Employee
	}{
		{
			desc: "success",
		},
		{
			desc:              "db error",
			getEmployeesError: errors.New("test"),
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.desc, func(t *testing.T) {
			call := suite.employeeRepositoryMock.On("GetEmployees", mock.Anything).Once()
			if tt.getEmployeesError != nil {
				call.Return(nil, tt.getEmployeesError)
			} else {
				faker.FakeData(&tt.expected, options.WithRandomMapAndSliceMaxSize(20))
				call.Return(tt.expected, nil)
			}

			actual, err := suite.target.RetrieveAllEmployees(context.Background())
			if tt.getEmployeesError != nil {
				assert.Empty(suite.T(), actual)
				assert.Equal(suite.T(), tt.getEmployeesError, err)
			} else {
				assert.Equal(suite.T(), tt.expected, actual)
				assert.NoError(suite.T(), err)
			}

			suite.employeeRepositoryMock.AssertExpectations(suite.T())
		})
	}
}

func (suite *EmployeeServiceSuite) TestAddEmployee() {
	tests := []struct {
		desc             string
		inputGender      string
		addEmployeeError error
		expected         int
	}{
		{
			desc:        "success",
			inputGender: faker.Word(),
			expected:    int(faker.RandomUnixTime()),
		},
		{
			desc:             "db error",
			addEmployeeError: errors.New("test"),
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.desc, func(t *testing.T) {
			call := suite.employeeRepositoryMock.On("AddEmployee", mock.Anything, tt.inputGender).Once()
			if tt.addEmployeeError != nil {
				call.Return(-1, tt.addEmployeeError)
			} else {
				call.Return(tt.expected, nil)
			}

			actual, err := suite.target.AddEmployee(context.Background(), string(tt.inputGender))
			if tt.addEmployeeError != nil {
				assert.Equal(suite.T(), -1, actual)
				assert.Equal(suite.T(), tt.addEmployeeError, err)
			} else {
				assert.Equal(suite.T(), tt.expected, actual)
				assert.NoError(suite.T(), err)
			}

			suite.employeeRepositoryMock.AssertExpectations(suite.T())
		})
	}
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestEmployeeServiceSuite(t *testing.T) {
	suite.Run(t, new(EmployeeServiceSuite))
}
