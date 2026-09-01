package domain

import (
	"time"
	"uuid"
)

type User struct {
	ID        uuid.UUID
	Name      string
	Role      string // admin, salesman
	CreatedAt time.Time
}

type UserCredentail struct {
	User
	Email    string
	Password string
}

type Product struct {
	ID          uuid.UUID
	Name        string
	Description string
	Price       float32
	IsSellable  bool
	AddedDate   time.Time
}

type Stock struct {
	ID              uuid.UUID
	Product         Product
	CurrentQuantity int
	LastSoldDate    *time.Time
	Location        Warehouse
}

type StockLog struct {
	ID        int
	Stock     Stock
	Action    string // sold, added, audited
	Count     int
	CreatedAt time.Time
}

type Warehouse struct {
	ID       uuid.UUID
	Key      string
	Location string
	Priority int
	IsActive bool
}

type Customer struct {
	ID       uuid.UUID
	Name     string
	Location string
}

type OrderItem struct {
	ID             int
	Product        Product
	Quantity       int
	Discount       float32
	ItemPriceTotal float32
}

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "Pending"
	OrderStatusAllocated OrderStatus = "Allocated"
	OrderStatusRejected  OrderStatus = "Rejected"
	OrderStatusSold      OrderStatus = "Sold"
)

type Order struct {
	ID                 int
	Customer           Customer
	Items              []OrderItem
	SubTotal           float32
	Tax                float32
	TotalPrice         float32
	Status             OrderStatus // sold, rejected, pending
	OrderDate          time.Time
	OrderCompletedDate *time.Time
}

type AllocationOrder struct {
	ID        int
	Product   Product
	Warehouse Warehouse
	Priority  int // runs based on higer priority
}

type AllocationStatus string

const (
	AllocationVoided    AllocationStatus = "VOIDED"
	AllocationPending   AllocationStatus = "PENDING"
	AllocationCompleted AllocationStatus = "COMPLETED"
)

type AllocationOrderLog struct {
	Allocation AllocationOrder
	CreatedAt  time.Time
	Status     string
}
