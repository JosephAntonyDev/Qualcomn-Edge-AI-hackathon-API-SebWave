package infra

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/JosephAntonyDev/Qualcomn-Edge-AI-hackathon-API-SebWave/internal/websocket"
)

var sensorMsgCounter = 0

// HandleSensorData procesa la información de sensores y actualiza la DB
func HandleSensorData(db *sql.DB) func(websocket.SensorDataMsg) {
	return func(data websocket.SensorDataMsg) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		var uuidStr string
		err := db.QueryRowContext(ctx, "SELECT id::text FROM intersections WHERE serial_number = $1 OR id::text = $1", data.IntersectionID).Scan(&uuidStr)
		if err != nil {
			log.Printf("Error encontrando interseccion %s: %v", data.IntersectionID, err)
			return
		}

		// Calcular densidad basica basada en ocupacion de lados
		// Esto es de ejemplo: Lado A -> NS, Lado B -> EO
		densNs := 0
		if data.LadoA.Ocupado {
			densNs = 80 // un ejemplo de %
		}
		densEo := 0
		if data.LadoB.Ocupado {
			densEo = 80
		}

		// 1a) Actualizar interseccion
		_, err = db.ExecContext(ctx, `
			UPDATE intersections 
			SET current_phase = $1, 
				current_density_ns = $2, 
				current_density_eo = $3, 
				last_heartbeat_at = NOW(), 
				status = 'connected' 
			WHERE id = $4
		`, data.Fase, densNs, densEo, uuidStr)
		if err != nil {
			log.Printf("Error actualizando interseccion: %v", err)
		}

		// 1b) Insertar en sensor_readings de tanto en tanto
		sensorMsgCounter++
		if sensorMsgCounter%10 == 0 { // Guardar 1 de cada 10 para no saturar
			// Insertamos como ejemplo para el Lado A (NS)
			_, _ = db.ExecContext(ctx, `
				INSERT INTO sensor_readings (intersection_id, is_occupied, distance_mm, energy_pct, timestamp)
				VALUES ($1, $2, $3, $4, NOW())
			`, uuidStr, data.LadoA.Ocupado, data.LadoA.Distancia, data.LadoA.Energia)
			sensorMsgCounter = 0
		}
	}
}

// HandleEmergency procesa alertas de emergencia
func HandleEmergency(db *sql.DB) func(websocket.EmergencyMsg) {
	return func(data websocket.EmergencyMsg) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		var uuidStr string
		err := db.QueryRowContext(ctx, "SELECT id::text FROM intersections WHERE serial_number = $1 OR id::text = $1", data.IntersectionID).Scan(&uuidStr)
		if err != nil {
			log.Printf("Error encontrando interseccion para emergencia %s: %v", data.IntersectionID, err)
			return
		}

		// Insertar evento de emergencia
		_, err = db.ExecContext(ctx, `
			INSERT INTO emergency_events (intersection_id, confidence_score, detection_method, timestamp)
			VALUES ($1, $2, $3, NOW())
		`, uuidStr, data.Confidence, data.Method)
		if err != nil {
			log.Printf("Error insertando evento de emergencia: %v", err)
		}

		// Registrar Alerta vinculada
		title := fmt.Sprintf("Emergencia Activa por %s", data.Method)
		description := "Vehículo de emergencia detectado en la intersección"
		_, err = db.ExecContext(ctx, `
			INSERT INTO alerts (intersection_id, type, severity, status, title, description, created_at)
			VALUES ($1, 'siren_detected', 'critical', 'active', $2, $3, NOW())
		`, uuidStr, title, description)
		if err != nil {
			log.Printf("Error insertando alerta de emergencia: %v", err)
		}
	}
}
