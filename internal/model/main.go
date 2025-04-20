package model

type Product struct {
	Name        string
	Price       float64
	Brand       string
	Model       string
	Description string
	Stores      []*Store
}

type Laptop struct {
	Product
	CPU             string
	RAM             int
	Storage         string
	DisplaySize     float64
	Graphics        string
	OperatingSystem string
	BatteryLife     float64
}

type Headphones struct {
	Product
	Type            string // проводные/беспроводные
	Connection      string // Bluetooth/AUX/Lightning
	Impedance       int
	Sensitivity     float64
	DriverSize      int
	NoiseCancelling bool
}

type TV struct {
	Product
	ScreenSize  float64
	Resolution  string
	PanelType   string
	RefreshRate int
	HDRSupport  bool
	SmartTV     bool
	Frequency   int
}

type Store struct {
	Name         string
	Address      string
	Phone        string
	Website      string
	WorkingHours string
	ProductIDs   []string
}
