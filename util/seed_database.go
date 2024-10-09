package util

import (
	"context"
	"log"
	"time"

	"github.com/themelancholyspirit/airline-reservation-system/storage"
	"github.com/themelancholyspirit/airline-reservation-system/types"
)

func SeedFlights(storage storage.Storage) error {
	flights := []types.Flight{
		{FlightNumber: "AA1234", DepartureAirport: "JFK", ArrivalAirport: "LAX", DepartureTime: time.Now().Add(24 * time.Hour).Format("2006-01-02T15:04:05Z"), ArrivalTime: time.Now().Add(28 * time.Hour).Format("2006-01-02T15:04:05Z"), Capacity: 200, AvailableSeats: 180, Price: 350.00, Status: "Scheduled"},
		{FlightNumber: "UA5678", DepartureAirport: "ORD", ArrivalAirport: "SFO", DepartureTime: time.Now().Add(36 * time.Hour).Format("2006-01-02T15:04:05Z"), ArrivalTime: time.Now().Add(40 * time.Hour).Format("2006-01-02T15:04:05Z"), Capacity: 180, AvailableSeats: 150, Price: 420.00, Status: "Scheduled"},
		{FlightNumber: "DL9012", DepartureAirport: "ATL", ArrivalAirport: "MIA", DepartureTime: time.Now().Add(12 * time.Hour).Format("2006-01-02T15:04:05Z"), ArrivalTime: time.Now().Add(14 * time.Hour).Format("2006-01-02T15:04:05Z"), Capacity: 150, AvailableSeats: 120, Price: 280.00, Status: "Scheduled"},
		{FlightNumber: "SW3456", DepartureAirport: "DEN", ArrivalAirport: "LAS", DepartureTime: time.Now().Add(8 * time.Hour).Format("2006-01-02T15:04:05Z"), ArrivalTime: time.Now().Add(10 * time.Hour).Format("2006-01-02T15:04:05Z"), Capacity: 140, AvailableSeats: 100, Price: 200.00, Status: "Scheduled"},
		{FlightNumber: "BA7890", DepartureAirport: "LHR", ArrivalAirport: "CDG", DepartureTime: time.Now().Add(48 * time.Hour).Format("2006-01-02T15:04:05Z"), ArrivalTime: time.Now().Add(49 * time.Hour).Format("2006-01-02T15:04:05Z"), Capacity: 220, AvailableSeats: 200, Price: 180.00, Status: "Scheduled"},
		{FlightNumber: "LH1357", DepartureAirport: "FRA", ArrivalAirport: "FCO", DepartureTime: time.Now().Add(72 * time.Hour).Format("2006-01-02T15:04:05Z"), ArrivalTime: time.Now().Add(74 * time.Hour).Format("2006-01-02T15:04:05Z"), Capacity: 190, AvailableSeats: 170, Price: 250.00, Status: "Scheduled"},
		{FlightNumber: "EK2468", DepartureAirport: "DXB", ArrivalAirport: "SIN", DepartureTime: time.Now().Add(96 * time.Hour).Format("2006-01-02T15:04:05Z"), ArrivalTime: time.Now().Add(103 * time.Hour).Format("2006-01-02T15:04:05Z"), Capacity: 300, AvailableSeats: 280, Price: 800.00, Status: "Scheduled"},
		{FlightNumber: "QF3690", DepartureAirport: "SYD", ArrivalAirport: "LAX", DepartureTime: time.Now().Add(120 * time.Hour).Format("2006-01-02T15:04:05Z"), ArrivalTime: time.Now().Add(134 * time.Hour).Format("2006-01-02T15:04:05Z"), Capacity: 280, AvailableSeats: 250, Price: 1200.00, Status: "Scheduled"},
		{FlightNumber: "AC4812", DepartureAirport: "YYZ", ArrivalAirport: "LHR", DepartureTime: time.Now().Add(60 * time.Hour).Format("2006-01-02T15:04:05Z"), ArrivalTime: time.Now().Add(67 * time.Hour).Format("2006-01-02T15:04:05Z"), Capacity: 250, AvailableSeats: 220, Price: 650.00, Status: "Scheduled"},
		{FlightNumber: "JL5934", DepartureAirport: "NRT", ArrivalAirport: "HNL", DepartureTime: time.Now().Add(84 * time.Hour).Format("2006-01-02T15:04:05Z"), ArrivalTime: time.Now().Add(93 * time.Hour).Format("2006-01-02T15:04:05Z"), Capacity: 260, AvailableSeats: 240, Price: 750.00, Status: "Scheduled"},
		{FlightNumber: "AF7056", DepartureAirport: "CDG", ArrivalAirport: "JFK", DepartureTime: time.Now().Add(108 * time.Hour).Format("2006-01-02T15:04:05Z"), ArrivalTime: time.Now().Add(116 * time.Hour).Format("2006-01-02T15:04:05Z"), Capacity: 270, AvailableSeats: 250, Price: 700.00, Status: "Scheduled"},
		{FlightNumber: "SQ8178", DepartureAirport: "SIN", ArrivalAirport: "SYD", DepartureTime: time.Now().Add(132 * time.Hour).Format("2006-01-02T15:04:05Z"), ArrivalTime: time.Now().Add(140 * time.Hour).Format("2006-01-02T15:04:05Z"), Capacity: 290, AvailableSeats: 270, Price: 550.00, Status: "Scheduled"},
		{FlightNumber: "LX9290", DepartureAirport: "ZRH", ArrivalAirport: "JFK", DepartureTime: time.Now().Add(156 * time.Hour).Format("2006-01-02T15:04:05Z"), ArrivalTime: time.Now().Add(165 * time.Hour).Format("2006-01-02T15:04:05Z"), Capacity: 230, AvailableSeats: 210, Price: 680.00, Status: "Scheduled"},
		{FlightNumber: "TK0412", DepartureAirport: "IST", ArrivalAirport: "DXB", DepartureTime: time.Now().Add(180 * time.Hour).Format("2006-01-02T15:04:05Z"), ArrivalTime: time.Now().Add(184 * time.Hour).Format("2006-01-02T15:04:05Z"), Capacity: 210, AvailableSeats: 190, Price: 320.00, Status: "Scheduled"},
		{FlightNumber: "KE1534", DepartureAirport: "ICN", ArrivalAirport: "LAX", DepartureTime: time.Now().Add(204 * time.Hour).Format("2006-01-02T15:04:05Z"), ArrivalTime: time.Now().Add(218 * time.Hour).Format("2006-01-02T15:04:05Z"), Capacity: 280, AvailableSeats: 260, Price: 900.00, Status: "Scheduled"},
		{FlightNumber: "AY2656", DepartureAirport: "HEL", ArrivalAirport: "BKK", DepartureTime: time.Now().Add(228 * time.Hour).Format("2006-01-02T15:04:05Z"), ArrivalTime: time.Now().Add(240 * time.Hour).Format("2006-01-02T15:04:05Z"), Capacity: 240, AvailableSeats: 220, Price: 620.00, Status: "Scheduled"},
		{FlightNumber: "EY3778", DepartureAirport: "AUH", ArrivalAirport: "LHR", DepartureTime: time.Now().Add(252 * time.Hour).Format("2006-01-02T15:04:05Z"), ArrivalTime: time.Now().Add(260 * time.Hour).Format("2006-01-02T15:04:05Z"), Capacity: 270, AvailableSeats: 250, Price: 580.00, Status: "Scheduled"},
		{FlightNumber: "NH4890", DepartureAirport: "HND", ArrivalAirport: "SFO", DepartureTime: time.Now().Add(276 * time.Hour).Format("2006-01-02T15:04:05Z"), ArrivalTime: time.Now().Add(288 * time.Hour).Format("2006-01-02T15:04:05Z"), Capacity: 250, AvailableSeats: 230, Price: 850.00, Status: "Scheduled"},
		{FlightNumber: "CX5912", DepartureAirport: "HKG", ArrivalAirport: "SYD", DepartureTime: time.Now().Add(300 * time.Hour).Format("2006-01-02T15:04:05Z"), ArrivalTime: time.Now().Add(309 * time.Hour).Format("2006-01-02T15:04:05Z"), Capacity: 260, AvailableSeats: 240, Price: 600.00, Status: "Scheduled"},
		{FlightNumber: "SU7034", DepartureAirport: "SVO", ArrivalAirport: "PEK", DepartureTime: time.Now().Add(324 * time.Hour).Format("2006-01-02T15:04:05Z"), ArrivalTime: time.Now().Add(332 * time.Hour).Format("2006-01-02T15:04:05Z"), Capacity: 220, AvailableSeats: 200, Price: 450.00, Status: "Scheduled"},
	}

	ctx := context.Background()

	for _, flight := range flights {
		if err := storage.CreateFlight(ctx, flight); err != nil {
			log.Printf("Failed to seed flight %s: %v", flight.FlightNumber, err)
			return err
		}
	}

	log.Println("Successfully seeded 20 flights")
	return nil
}
