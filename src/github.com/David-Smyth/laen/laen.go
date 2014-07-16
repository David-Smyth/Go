package laen

// Launch Energy


// Universal Gravitations Constant -- so weak, its tough to measure well
const G      = 6.67384e-11 // CODATA 2010
const Gold   = 6.6726e-11  // a common, older value
const GQuinn = 6.67545e-11 // by Terry Quinn, 9/2013
const GTino  = 6.67191e-11 // by Guglielmo Tino, 6/2014

func Gdelta() float64 {
    return GQuinn - GTino
}

const RadiusEarth     =  6.378e3  // kilometers
const MassEarth       =  5.974e24 // kg
const GeoSyncAltitude = 35.790e3  // kilometers

func OrbitalRadius( altitude float64 ) float64 {
    return altitude + RadiusEarth
}

func CircularOrbitEnergy( satelliteKg, altitudeKm float64 ) float64 {
    GMm := Gold * MassEarth * satelliteKg
    Re  := RadiusEarth * 1000 // convert to meters
    Rs  := OrbitalRadius(altitudeKm) * 1000 // convert to meters
    return GMm * ((1/(2*Re)) - (1/(2*Rs)))
}
