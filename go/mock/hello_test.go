package mock

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestCompany_Meeting(t *testing.T) {
	person := NewPerson("Mack")
	company := NewCompany(person)
	t.Log(company.Meeting("Jack"))
}

func TestCompany_Meeting2(t *testing.T) {
	ctl := gomock.NewController(t)
	mockTalker := NewMockTalker(ctl)
	mockTalker.EXPECT().SayHello(gomock.Any()).Return("Any")

	company := NewCompany(mockTalker)
	t.Log(company.Meeting("Jack"))
}
