package main

//Node for a Tree or List
type City struct {
	Name        string
	Connections []Link
}

//A linked city
type Link struct {
	City     City  //The linked city
	Cost     int64 //The cost to get from the city to the linked city
	Distance int64 //The distance between the 2 cities
}

//Where the trip is starting from, going to, and the path to get there.
type Network struct {
	Start  City
	Target City
	Path   []City
}

func SetStartCity(city string) Network {
	startCity := City{Name: city}
	newNetwork := Network{Start: startCity}

	return newNetwork
}
