package abstractFactory

import "testing"

func getMainAndDetail(factory DAOFactory) {
	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderDetailDAO().SaveOrderDetail()
}

func ExampleRdbFactory() {
	var factory DAOFactory
	factory = &RDBDAOFactory{}
	getMainAndDetail(factory)
	// Output:
	// rdb main save
	// rdb detail save
}

func TestRdbFactory(t *testing.T) {
	var factory DAOFactory
	factory = &RDBDAOFactory{}
	getMainAndDetail(factory)
}

func ExampleXmlFactory() {
	var factory DAOFactory
	factory = &XMLDAOFactory{}
	getMainAndDetail(factory)
	// Output:
	// xml main save
	// xml detail save
}

func TestXmlFactory(t *testing.T) {
	var factory DAOFactory
	factory = &XMLDAOFactory{}
	getMainAndDetail(factory)
	// Output:
	// xml main save
	// xml detail save
}
