package main

//Node for a Tree or List
type City struct {
	Name          string
	NeighborCites []Link
}

//A linked city
type Link struct {
	City     City  //The linked city
	Distance int64 //The distance between the 2 cities
}

type Path struct {
	Current       City
	Target        City
	Route         []Link
	RouteDistance int64
}

func FindPath(p Path) Path {
	currentPath := p
	currentCity := p.Current

	for _, NeighborCity := range currentCity.NeighborCites {
		if &NeighborCity.City == &currentPath.Target {
			var distanceFromStart int64
			for _, link := range currentPath.Route {
				distanceFromStart += link.Distance
			}
			if distanceFromStart < currentPath.RouteDistance {
				currentPath.RouteDistance = distanceFromStart
				currentPath.Route = append(currentPath.Route, NeighborCity)
				return currentPath
			}
		} else {
			currentPath = FindPath(currentPath)
		}
	}
	return currentPath
}
