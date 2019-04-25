// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package item

import (
	"database/sql"
	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
	"sync"
)

var (
	lockServiceMockAdd                 sync.RWMutex
	lockServiceMockCheckItemExist      sync.RWMutex
	lockServiceMockDelete              sync.RWMutex
	lockServiceMockFindByID            sync.RWMutex
	lockServiceMockGetAllItemsByMenuID sync.RWMutex
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
//             AddFunc: func(tx *sql.Tx, i *domain.Item) (*models.Item, error) {
// 	               panic("mock out the Add method")
//             },
//             CheckItemExistFunc: func(tx *sql.Tx, itemID int) (bool, error) {
// 	               panic("mock out the CheckItemExist method")
//             },
//             DeleteFunc: func(tx *sql.Tx, i *models.Item) error {
// 	               panic("mock out the Delete method")
//             },
//             FindByIDFunc: func(tx *sql.Tx, itemID int) (*models.Item, error) {
// 	               panic("mock out the FindByID method")
//             },
//             GetAllItemsByMenuIDFunc: func(menuID int) ([]*models.Item, error) {
// 	               panic("mock out the GetAllItemsByMenuID method")
//             },
//         }
//
//         // use mockedService in code that requires Service
//         // and then make assertions.
//
//     }
type ServiceMock struct {
	// AddFunc mocks the Add method.
	AddFunc func(tx *sql.Tx, i *domain.Item) (*models.Item, error)

	// CheckItemExistFunc mocks the CheckItemExist method.
	CheckItemExistFunc func(tx *sql.Tx, itemID int) (bool, error)

	// DeleteFunc mocks the Delete method.
	DeleteFunc func(tx *sql.Tx, i *models.Item) error

	// FindByIDFunc mocks the FindByID method.
	FindByIDFunc func(tx *sql.Tx, itemID int) (*models.Item, error)

	// GetAllItemsByMenuIDFunc mocks the GetAllItemsByMenuID method.
	GetAllItemsByMenuIDFunc func(menuID int) ([]*models.Item, error)

	// calls tracks calls to the methods.
	calls struct {
		// Add holds details about calls to the Add method.
		Add []struct {
			// Tx is the tx argument value.
			Tx *sql.Tx
			// I is the i argument value.
			I *domain.Item
		}
		// CheckItemExist holds details about calls to the CheckItemExist method.
		CheckItemExist []struct {
			// Tx is the tx argument value.
			Tx *sql.Tx
			// ItemID is the itemID argument value.
			ItemID int
		}
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// Tx is the tx argument value.
			Tx *sql.Tx
			// I is the i argument value.
			I *models.Item
		}
		// FindByID holds details about calls to the FindByID method.
		FindByID []struct {
			// Tx is the tx argument value.
			Tx *sql.Tx
			// ItemID is the itemID argument value.
			ItemID int
		}
		// GetAllItemsByMenuID holds details about calls to the GetAllItemsByMenuID method.
		GetAllItemsByMenuID []struct {
			// MenuID is the menuID argument value.
			MenuID int
		}
	}
}

// Add calls AddFunc.
func (mock *ServiceMock) Add(tx *sql.Tx, i *domain.Item) (*models.Item, error) {
	if mock.AddFunc == nil {
		panic("ServiceMock.AddFunc: method is nil but Service.Add was just called")
	}
	callInfo := struct {
		Tx *sql.Tx
		I  *domain.Item
	}{
		Tx: tx,
		I:  i,
	}
	lockServiceMockAdd.Lock()
	mock.calls.Add = append(mock.calls.Add, callInfo)
	lockServiceMockAdd.Unlock()
	return mock.AddFunc(tx, i)
}

// AddCalls gets all the calls that were made to Add.
// Check the length with:
//     len(mockedService.AddCalls())
func (mock *ServiceMock) AddCalls() []struct {
	Tx *sql.Tx
	I  *domain.Item
} {
	var calls []struct {
		Tx *sql.Tx
		I  *domain.Item
	}
	lockServiceMockAdd.RLock()
	calls = mock.calls.Add
	lockServiceMockAdd.RUnlock()
	return calls
}

// CheckItemExist calls CheckItemExistFunc.
func (mock *ServiceMock) CheckItemExist(tx *sql.Tx, itemID int) (bool, error) {
	if mock.CheckItemExistFunc == nil {
		panic("ServiceMock.CheckItemExistFunc: method is nil but Service.CheckItemExist was just called")
	}
	callInfo := struct {
		Tx     *sql.Tx
		ItemID int
	}{
		Tx:     tx,
		ItemID: itemID,
	}
	lockServiceMockCheckItemExist.Lock()
	mock.calls.CheckItemExist = append(mock.calls.CheckItemExist, callInfo)
	lockServiceMockCheckItemExist.Unlock()
	return mock.CheckItemExistFunc(tx, itemID)
}

// CheckItemExistCalls gets all the calls that were made to CheckItemExist.
// Check the length with:
//     len(mockedService.CheckItemExistCalls())
func (mock *ServiceMock) CheckItemExistCalls() []struct {
	Tx     *sql.Tx
	ItemID int
} {
	var calls []struct {
		Tx     *sql.Tx
		ItemID int
	}
	lockServiceMockCheckItemExist.RLock()
	calls = mock.calls.CheckItemExist
	lockServiceMockCheckItemExist.RUnlock()
	return calls
}

// Delete calls DeleteFunc.
func (mock *ServiceMock) Delete(tx *sql.Tx, i *models.Item) error {
	if mock.DeleteFunc == nil {
		panic("ServiceMock.DeleteFunc: method is nil but Service.Delete was just called")
	}
	callInfo := struct {
		Tx *sql.Tx
		I  *models.Item
	}{
		Tx: tx,
		I:  i,
	}
	lockServiceMockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	lockServiceMockDelete.Unlock()
	return mock.DeleteFunc(tx, i)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//     len(mockedService.DeleteCalls())
func (mock *ServiceMock) DeleteCalls() []struct {
	Tx *sql.Tx
	I  *models.Item
} {
	var calls []struct {
		Tx *sql.Tx
		I  *models.Item
	}
	lockServiceMockDelete.RLock()
	calls = mock.calls.Delete
	lockServiceMockDelete.RUnlock()
	return calls
}

// FindByID calls FindByIDFunc.
func (mock *ServiceMock) FindByID(tx *sql.Tx, itemID int) (*models.Item, error) {
	if mock.FindByIDFunc == nil {
		panic("ServiceMock.FindByIDFunc: method is nil but Service.FindByID was just called")
	}
	callInfo := struct {
		Tx     *sql.Tx
		ItemID int
	}{
		Tx:     tx,
		ItemID: itemID,
	}
	lockServiceMockFindByID.Lock()
	mock.calls.FindByID = append(mock.calls.FindByID, callInfo)
	lockServiceMockFindByID.Unlock()
	return mock.FindByIDFunc(tx, itemID)
}

// FindByIDCalls gets all the calls that were made to FindByID.
// Check the length with:
//     len(mockedService.FindByIDCalls())
func (mock *ServiceMock) FindByIDCalls() []struct {
	Tx     *sql.Tx
	ItemID int
} {
	var calls []struct {
		Tx     *sql.Tx
		ItemID int
	}
	lockServiceMockFindByID.RLock()
	calls = mock.calls.FindByID
	lockServiceMockFindByID.RUnlock()
	return calls
}

// GetAllItemsByMenuID calls GetAllItemsByMenuIDFunc.
func (mock *ServiceMock) GetAllItemsByMenuID(menuID int) ([]*models.Item, error) {
	if mock.GetAllItemsByMenuIDFunc == nil {
		panic("ServiceMock.GetAllItemsByMenuIDFunc: method is nil but Service.GetAllItemsByMenuID was just called")
	}
	callInfo := struct {
		MenuID int
	}{
		MenuID: menuID,
	}
	lockServiceMockGetAllItemsByMenuID.Lock()
	mock.calls.GetAllItemsByMenuID = append(mock.calls.GetAllItemsByMenuID, callInfo)
	lockServiceMockGetAllItemsByMenuID.Unlock()
	return mock.GetAllItemsByMenuIDFunc(menuID)
}

// GetAllItemsByMenuIDCalls gets all the calls that were made to GetAllItemsByMenuID.
// Check the length with:
//     len(mockedService.GetAllItemsByMenuIDCalls())
func (mock *ServiceMock) GetAllItemsByMenuIDCalls() []struct {
	MenuID int
} {
	var calls []struct {
		MenuID int
	}
	lockServiceMockGetAllItemsByMenuID.RLock()
	calls = mock.calls.GetAllItemsByMenuID
	lockServiceMockGetAllItemsByMenuID.RUnlock()
	return calls
}
