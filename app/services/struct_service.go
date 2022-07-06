package services

type Paginate struct {
	Data interface{}
	Page uint
	PageSize uint
	TotalRows uint
}

func (page *Paginate) Paginate()  {

}