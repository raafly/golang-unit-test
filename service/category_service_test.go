package service

import (
	"belajar-golang-unit-test/repository"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryServie = CategoryService{Repository: categoryRepository}

func TestCategoryService_Get( t *testing.T ) {

	// program mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryServie.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)

}