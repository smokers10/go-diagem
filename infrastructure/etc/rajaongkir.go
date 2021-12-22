package etc

const RajaOngkirAPIKey = "14b3edaa6c21f0d7e6b0a4244664e7a1"

type RajaOngkirReqBody struct {
	Origin      string
	Destination string
	Weight      int
	Curier      string
}
