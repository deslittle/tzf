package convert

import (
	"fmt"

	"github.com/deslittle/pinpoint/pb"
	"github.com/mitchellh/mapstructure"
)

const (
	MultiPolygonType = "MultiPolygon"
	PolygonType      = "Polygon"
	FeatureType      = "Feature"
)

type PolygonCoordinates [][][2]float64
type MultiPolygonCoordinates []PolygonCoordinates

type GeometryDefine struct {
	Coordinates interface{} `json:"coordinates"`
	Type        string      `json:"type"`
}

type PropertiesDefine struct {
	Name string `json:"Name"`
}

type FeatureItem struct {
	Geometry   GeometryDefine   `json:"geometry"`
	Properties PropertiesDefine `json:"properties"`
	Type       string           `json:"type"`
}

type BoundaryFile struct {
	Type     string         `json:"type"`
	Features []*FeatureItem `json:"features"`
}

func Do(input *BoundaryFile) (*pb.Locations, error) {
	output := make([]*pb.Location, 0)

	for _, item := range input.Features {
		pblocItem := &pb.Location{
			Name: item.Properties.Name,
		}

		var coordinates MultiPolygonCoordinates

		MultiPolygonTypeHandler := func() error {
			if err := mapstructure.Decode(item.Geometry.Coordinates, &coordinates); err != nil {
				return err
			}
			return nil
		}
		PolygonTypeHandler := func() error {
			var polygonCoordinates PolygonCoordinates
			if err := mapstructure.Decode(item.Geometry.Coordinates, &polygonCoordinates); err != nil {
				return err
			}
			coordinates = append(coordinates, polygonCoordinates)
			return nil
		}

		switch item.Type {
		case MultiPolygonType:
			if err := MultiPolygonTypeHandler(); err != nil {
				return nil, err
			}
		case PolygonType:
			if err := PolygonTypeHandler(); err != nil {
				return nil, err
			}
		case FeatureType:
			switch item.Geometry.Type {
			case MultiPolygonType:
				if err := MultiPolygonTypeHandler(); err != nil {
					return nil, err
				}
			case PolygonType:
				if err := PolygonTypeHandler(); err != nil {
					return nil, err
				}
			default:
				return nil, fmt.Errorf("unknown type %v", item.Type)
			}
		default:
			return nil, fmt.Errorf("unknown type %v", item.Type)
		}

		polygons := make([]*pb.Polygon, 0)

		for _, subcoordinates := range coordinates {
			newpbPoly := &pb.Polygon{
				Points: make([]*pb.Point, 0),
				Holes:  make([]*pb.Polygon, 0),
			}
			for index, geoPoly := range subcoordinates {
				if index == 0 {
					for _, rawCoords := range geoPoly {
						newpbPoly.Points = append(newpbPoly.Points, &pb.Point{
							Lng: float32(rawCoords[0]),
							Lat: float32(rawCoords[1]),
						})
					}
					continue
				}

				holePoly := &pb.Polygon{
					Points: make([]*pb.Point, 0),
				}
				for _, rawCoords := range geoPoly {
					holePoly.Points = append(holePoly.Points, &pb.Point{
						Lng: float32(rawCoords[0]),
						Lat: float32(rawCoords[1]),
					})
				}
				newpbPoly.Holes = append(newpbPoly.Holes, holePoly)

			}
			polygons = append(polygons, newpbPoly)
		}

		pblocItem.Polygons = polygons
		output = append(output, pblocItem)
	}

	return &pb.Locations{
		Locations: output,
	}, nil
}
