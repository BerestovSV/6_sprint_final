package service_test

import (
	"testing"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func Test_AutoConvert(t *testing.T) {
	_, err := service.AutoConvert("ВСЧШИЕСФТУЫРЛЭЩРЛЩВЭЛВЗТ")
	if err != nil {
		t.Error(err)
	}
}
