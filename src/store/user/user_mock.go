// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package user

import (
	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
	"sync"
)

var (
	lockServiceMockCreate  sync.RWMutex
	lockServiceMockFind    sync.RWMutex
	lockServiceMockFindAll sync.RWMutex
)

// Ensure, that ServiceMock does implement Service.
// If this is not the case, regenerate this file with moq.
var _ Service = &ServiceMock{}

// ServiceMock is a mock implementation of Service.
//
//     func TestSomethingThatUsesService(t *testing.T) {
//
//         // make and configure a mocked Service
//         mockedService := &ServiceMock{
//             CreateFunc: func(p *domain.CreateUserInput) (*models.User, error) {
// 	               panic("mock out the Create method")
//             },
//             FindFunc: func(p *domain.CreateUserInput) (*models.User, error) {
// 	               panic("mock out the Find method")
//             },
//             FindAllFunc: func() ([]*models.User, error) {
// 	               panic("mock out the FindAll method")
//             },
//         }
//
//         // use mockedService in code that requires Service
//         // and then make assertions.
//
//     }
type ServiceMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(p *domain.CreateUserInput) (*models.User, error)

	// FindFunc mocks the Find method.
	FindFunc func(p *domain.CreateUserInput) (*models.User, error)

	// FindAllFunc mocks the FindAll method.
	FindAllFunc func() ([]*models.User, error)

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// P is the p argument value.
			P *domain.CreateUserInput
		}
		// Find holds details about calls to the Find method.
		Find []struct {
			// P is the p argument value.
			P *domain.CreateUserInput
		}
		// FindAll holds details about calls to the FindAll method.
		FindAll []struct {
		}
	}
}

// Create calls CreateFunc.
func (mock *ServiceMock) Create(p *domain.CreateUserInput) (*models.User, error) {
	if mock.CreateFunc == nil {
		panic("ServiceMock.CreateFunc: method is nil but Service.Create was just called")
	}
	callInfo := struct {
		P *domain.CreateUserInput
	}{
		P: p,
	}
	lockServiceMockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	lockServiceMockCreate.Unlock()
	return mock.CreateFunc(p)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//     len(mockedService.CreateCalls())
func (mock *ServiceMock) CreateCalls() []struct {
	P *domain.CreateUserInput
} {
	var calls []struct {
		P *domain.CreateUserInput
	}
	lockServiceMockCreate.RLock()
	calls = mock.calls.Create
	lockServiceMockCreate.RUnlock()
	return calls
}

// Find calls FindFunc.
func (mock *ServiceMock) Find(p *domain.CreateUserInput) (*models.User, error) {
	if mock.FindFunc == nil {
		panic("ServiceMock.FindFunc: method is nil but Service.Find was just called")
	}
	callInfo := struct {
		P *domain.CreateUserInput
	}{
		P: p,
	}
	lockServiceMockFind.Lock()
	mock.calls.Find = append(mock.calls.Find, callInfo)
	lockServiceMockFind.Unlock()
	return mock.FindFunc(p)
}

// FindCalls gets all the calls that were made to Find.
// Check the length with:
//     len(mockedService.FindCalls())
func (mock *ServiceMock) FindCalls() []struct {
	P *domain.CreateUserInput
} {
	var calls []struct {
		P *domain.CreateUserInput
	}
	lockServiceMockFind.RLock()
	calls = mock.calls.Find
	lockServiceMockFind.RUnlock()
	return calls
}

// FindAll calls FindAllFunc.
func (mock *ServiceMock) FindAll() ([]*models.User, error) {
	if mock.FindAllFunc == nil {
		panic("ServiceMock.FindAllFunc: method is nil but Service.FindAll was just called")
	}
	callInfo := struct {
	}{}
	lockServiceMockFindAll.Lock()
	mock.calls.FindAll = append(mock.calls.FindAll, callInfo)
	lockServiceMockFindAll.Unlock()
	return mock.FindAllFunc()
}

// FindAllCalls gets all the calls that were made to FindAll.
// Check the length with:
//     len(mockedService.FindAllCalls())
func (mock *ServiceMock) FindAllCalls() []struct {
} {
	var calls []struct {
	}
	lockServiceMockFindAll.RLock()
	calls = mock.calls.FindAll
	lockServiceMockFindAll.RUnlock()
	return calls
}
