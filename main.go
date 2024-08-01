package main

import "fmt"

// Observer interface that all observers must implement
type Observer interface {
	Update(ticker string, message string)
}

// Subject interface that defines methods for attaching and notifying observers
type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObservers()
}

// The subject that holds the state and notifies observer
type StockDataSubject struct {
	observers []Observer
	price     float64
	ticker    string
}

// add a new observer to the list of observers
func (s *StockDataSubject) RegisterObserver(o Observer) {
	s.observers = append(s.observers, o)
}

// remove an observer from the list of observers
func (s *StockDataSubject) RemoveObserver(o Observer) {
	var indexToRemove int
	for i, observer := range s.observers {
		if observer == o {
			indexToRemove = i
			break
		}
	}
	s.observers = append(s.observers[:indexToRemove], s.observers[indexToRemove+1:]...)
}

// notifies all registred observers of the state change
func (s *StockDataSubject) NotifyObservers() {
	for _, observer := range s.observers {
		observer.Update(s.ticker, s.ticker+" price changed to "+fmt.Sprintf("%f", s.price))
	}
}

// sets the new sticker price and notifies observers
func (s *StockDataSubject) UpdateStockPrice(ticker string, price float64) {
	s.ticker = ticker
	s.price = price
	s.NotifyObservers()
}

// An Observer that displays the stock price update
type CurrentConditionsDisplay struct {
	ID     string
	ticker string
}

// Update method that is called by the subject when the state changes
func (o *CurrentConditionsDisplay) Update(ticker string, message string) {
	if ticker == o.ticker {
		fmt.Printf("Observer %s received message: %s\n", o.ID, message)

	}
}

// Client code
func main() {
	// Create a new stock data subject
	stockDataSubject := &StockDataSubject{}

	// Create new observers
	observer1 := &CurrentConditionsDisplay{ID: "O1", ticker: "AAPL"}
	observer2 := &CurrentConditionsDisplay{ID: "O2", ticker: "MSFT"}
	observer3 := &CurrentConditionsDisplay{ID: "O3", ticker: "GOOG"}

	// Register observers
	stockDataSubject.RegisterObserver(observer1)
	stockDataSubject.RegisterObserver(observer2)
	stockDataSubject.RegisterObserver(observer3)

	// Update stock prices
	stockDataSubject.UpdateStockPrice("AAPL", 223.0)
	stockDataSubject.UpdateStockPrice("AAPL", 230.0)
	stockDataSubject.UpdateStockPrice("AAPL", 220.0)

	stockDataSubject.UpdateStockPrice("MSFT", 425.0)
	stockDataSubject.UpdateStockPrice("MSFT", 420.0)
	stockDataSubject.UpdateStockPrice("MSFT", 430.0)

	stockDataSubject.UpdateStockPrice("GOOG", 174.0)
	stockDataSubject.UpdateStockPrice("GOOG", 180.0)
	stockDataSubject.UpdateStockPrice("GOOG", 170.0)

	// Remove observer1
	stockDataSubject.RemoveObserver(observer1)

	// Update stock prices and observer1 should not receive the message
	stockDataSubject.UpdateStockPrice("AAPL", 220.0)
	stockDataSubject.UpdateStockPrice("MSFT", 440.0)
}
