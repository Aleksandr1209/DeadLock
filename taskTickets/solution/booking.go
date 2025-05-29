package solution

import (
	"sync"
)

type Sector struct {
	sync.Mutex
	ID    int
	Seats int
}

type BookingSystem struct {
	Sectors []*Sector
}

func NewBookingSystem(numSectors, seatsPerSector int) *BookingSystem {
	sectors := make([]*Sector, numSectors)
	for i := 0; i < numSectors; i++ {
		sectors[i] = &Sector{ID: i, Seats: seatsPerSector}
	}
	return &BookingSystem{Sectors: sectors}
}

func (bs *BookingSystem) Book(from, to int) bool {
	if from == to {
		return false
	}

	s1 := bs.Sectors[from]
	s2 := bs.Sectors[to]

	// Гарантированный порядок блокировок
	if s1.ID < s2.ID {
		s1.Lock()
		s2.Lock()
	} else {
		s2.Lock()
		s1.Lock()
	}
	defer s1.Unlock()
	defer s2.Unlock()

	if s1.Seats > 0 {
		s1.Seats--
		s2.Seats++
		return true
	}
	return false
}
