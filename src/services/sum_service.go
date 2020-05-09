package services

var SumService sumServiceInterface = &sumService{}

type sumServiceInterface interface {
	Sum(int32, int32) int32
}

type sumService struct{}

func (*sumService) Sum(a int32, b int32) int32 {
	return a + b
}
