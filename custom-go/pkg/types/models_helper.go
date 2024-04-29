package types

import (
	"fmt"
	"github.com/spf13/cast"
	"io"
	"math"
	"strings"
)

type UploadFile struct {
	Reader io.Reader
	Name   string
}

const (
	geometryStringPrefix = "SRID=4326;POINT("
	geometryStringSuffix = ")"
)

type Geometry struct {
	longitude float64
	latitude  float64
}

func NewGeometry(longitude, latitude float64) Geometry {
	return Geometry{longitude: longitude, latitude: latitude}
}

func (g Geometry) DistanceTo(target Geometry) float64 {
	p := math.Abs(g.longitude - target.longitude)
	q := math.Abs(g.latitude - target.latitude)
	return math.Hypot(p, q)
}

func (g Geometry) String() string {
	return fmt.Sprintf("%s%f %f%s", geometryStringPrefix, g.longitude, g.latitude, geometryStringSuffix)
}

func (g Geometry) MarshalJSON() ([]byte, error) {
	return []byte(g.String()), nil
}

func (g *Geometry) UnmarshalJSON(dataBytes []byte) error {
	if len(dataBytes) == 0 {
		return nil
	}
	dataString := strings.TrimSuffix(string(dataBytes), geometryStringSuffix)
	dataString, ok := strings.CutPrefix(dataString, geometryStringPrefix)
	if !ok {
		return fmt.Errorf("not supported geometry data [%s]", string(dataBytes))
	}
	dataArray := strings.Split(dataString, " ")
	if dataLength := len(dataArray); dataLength != 2 {
		return fmt.Errorf("expected 2 elements but found [%d]", dataLength)
	}

	*g = Geometry{longitude: cast.ToFloat64(dataArray[0]), latitude: cast.ToFloat64(dataArray[1])}
	return nil
}
