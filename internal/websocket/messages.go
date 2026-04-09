package websocket

import "encoding/json"

type WSMessage struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type SensorDataMsg struct {
	IntersectionID string `json:"intersection_id"`
	LadoA          struct {
		Ocupado   bool `json:"ocupado"`
		Distancia int  `json:"distancia"`
		Energia   int  `json:"energia"`
	} `json:"lado_a"`
	LadoB struct {
		Ocupado   bool `json:"ocupado"`
		Distancia int  `json:"distancia"`
	} `json:"lado_b"`
	Fase       string `json:"fase"`
	VerdeA     int    `json:"verde_a"`
	VerdeB     int    `json:"verde_b"`
	Emergencia bool   `json:"emergencia"`
	Timestamp  int64  `json:"timestamp"`
}

type EmergencyMsg struct {
	IntersectionID string  `json:"intersection_id"`
	Active         bool    `json:"active"`
	Confidence     float64 `json:"confidence"`
	Method         string  `json:"method"`
}

type IntersectionStatusMsg struct {
	IntersectionID string `json:"intersection_id"`
	Name           string `json:"name"`
	Status         string `json:"status"`
	OperationMode  string `json:"operation_mode"`
	Phase          string `json:"phase"`
	LadoA          struct {
		Ocupado   bool `json:"ocupado"`
		Distancia int  `json:"distancia"`
		Energia   int  `json:"energia"`
	} `json:"lado_a"`
	LadoB struct {
		Ocupado   bool `json:"ocupado"`
		Distancia int  `json:"distancia"`
	} `json:"lado_b"`
	Emergency bool  `json:"emergency"`
	Timestamp int64 `json:"timestamp"`
}
