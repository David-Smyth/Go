package laen

import (
    "testing"
    "math"
)

func TestG( t *testing.T) {
    if G != 6.67384e-11 {
        t.Errorf("G = %v, should be %v", G, 6.67384e-11 )
    }
    if GQuinn != 6.67545e-11 {
        t.Errorf("GQuinn = %v, should be %v", GQuinn, 6.67545e-11 )
    }
    if GTino != 6.67191e-11 {
        t.Errorf("GTino = %v, should be %v", GTino, 6.67191e-11 )
    }
    if Gdelta() != GQuinn - GTino {
        t.Errorf("Gdelta() = %v, should be %v",
                 Gdelta(), 6.67545e-11 - 6.67191e-11 )
    }
}

func TestRadiusEarth( t *testing.T) {
    if RadiusEarth != 6.378e3 {
        t.Errorf("RadiusEarth = %v, should be %v", RadiusEarth, 6.378e3 )
    }
}

func TestMassEarth( t *testing.T) {
    if MassEarth != 5.974e24 {
        t.Errorf("MassEarth = %v, should be %v", MassEarth, 5.974e24 )
    } 
}

func TestGeoSyncAltitude( t *testing.T) {
    if GeoSyncAltitude != 35790 {
        t.Errorf("GeoSyncAltitude = %v, should be %v", GeoSyncAltitude, 35790 )
    } 
}


func TestOrbitalRadius( t *testing.T) {
    if OrbitalRadius(1000) != 7.378e3 {
        t.Errorf("OrbitalRadius(1000) = %v, should be %v",
                 OrbitalRadius(1000), 7.378e3 )
    } 
    if OrbitalRadius(GeoSyncAltitude) != 42.168e3 {
        t.Errorf("OrbitalRadius(GeoSyncAltitude) = %v, should be %v",
                 OrbitalRadius(GeoSyncAltitude), 42.168e3 )
    } 
}

func TestCircularOrbitEnergy( t *testing.T) {
    err := math.Abs( CircularOrbitEnergy(500,1000) - (4.235e9/2) )
    if err > 0.001e9 {
        t.Errorf("CircularOrbitEnergy(500,1000) = %v, should be %v",
                 CircularOrbitEnergy(500,1000), (4.235e9/2) )
    }
    err = math.Abs( CircularOrbitEnergy(500,GeoSyncAltitude) - (26.52e9/2) )
    if err > 0.01e9 { 
        t.Errorf("CircularOrbitEnergy(500,GeoSyncAltitude) = %v, should be %v",
                 CircularOrbitEnergy(500,GeoSyncAltitude), (26.52e9/2) )
    } 
    e1k := CircularOrbitEnergy(500,1000)
    egs := CircularOrbitEnergy(500,GeoSyncAltitude)
    de  := egs - e1k
    err  = math.Abs( de - 1.11e10 )
    if err > 0.1e9 { 
        t.Errorf("energy delta = %v, should be %v",
                 de, 1.11e10 )
    }
}


