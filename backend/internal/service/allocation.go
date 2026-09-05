package service

import (
// "inventory/internal/domain"
)

// ### Scenario:

// 1. Customer orders 10 units of Laptop.
// 2. Warehouse stock available:
// 	• Dubai (Priority 1): 8 units
// 	• Sharjah (Priority 2): 20 units
// 	• Abu Dhabi (Priority 3): 0 units

// ### Business Question:

// • Rule 1 (Single Warehouse Preference): If one warehouse
// can fulfill the whole order (Sharjah has 20), should we
// prefer fulfilling entirely from Sharjah to avoid
// splitting shipments, or take 8 from Dubai (higher
// priority) and 2 from Sharjah?
// • For our first basic version: Let's say we fulfill
// based on highest warehouse priority first (Take 8 from
// Dubai, then 2 from Sharjah).

type AllocationItem struct {
	WarehouseID string
	ProductID   string
	Quantity    int
}

func Allocate() error {
	return nil
}
