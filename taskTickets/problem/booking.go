package problem

import (
	"sync"
	"time"
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
	bs.Sectors[from].Lock()
	time.Sleep(time.Millisecond)
	bs.Sectors[to].Lock()
	defer bs.Sectors[from].Unlock()
	defer bs.Sectors[to].Unlock()

	if bs.Sectors[from].Seats > 0 {
		bs.Sectors[from].Seats--
		bs.Sectors[to].Seats++
		return true
	}
	return false
}
