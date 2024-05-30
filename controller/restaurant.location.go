package controller

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func ViewAdminArea1() gin.HandlerFunc {
	return func(c *gin.Context) {
		country := c.GetString("country")

		// Create a new driver for Neo4j
		driver, err := neo4j.NewDriverWithContext(Neo4j, neo4j.BasicAuth(Neo4j_User, Neo4j_Password, ""))
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			if err := driver.Close(context.Background()); err != nil {
				log.Printf("Error closing Neo4j driver: %v", err)
			}
		}()

		// Create a new session
		session := driver.NewSession(context.Background(), neo4j.SessionConfig{DatabaseName: "location"})
		defer func() {
			if err := session.Close(context.Background()); err != nil {
				log.Printf("Error closing Neo4j session: %v", err)
			}
		}()

		// Run the query to find areas with REQUESTED relationship
		result, err := session.ExecuteRead(context.Background(),
			func(tx neo4j.ManagedTransaction) (interface{}, error) {
				query := `
                    MATCH (r:Country {id: $country})<-[:PART_OF]-(u:AdministrativeArea1)
                    RETURN u { .id } AS area1
                `
				parameters := map[string]interface{}{
					"country": country,
				}

				result, err := tx.Run(context.Background(), query, parameters)
				if err != nil {
					return nil, err
				}

				var areas []string
				for result.NextRecord(context.Background(), nil) {
					userNode, ok := result.Record().Get("area1")
					if !ok {
						return nil, errors.New("failed to get user node")
					}

					userMap, ok := userNode.(map[string]interface{})
					if !ok {
						return nil, errors.New("failed to convert user node to map")
					}

					id, ok := userMap["id"].(string)
					if !ok {
						id = "" // or any default value
					}

					areas = append(areas, id)
				}

				if err := result.Err(); err != nil {
					return nil, err
				}

				return areas, nil
			},
		)

		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if areas, ok := result.([]string); ok && len(areas) == 0 {
			c.JSON(http.StatusOK, []string{})
			return
		}

		// Return the list of areas with matching structure
		c.JSON(http.StatusOK, result)
	}
}

func ViewAdminArea2() gin.HandlerFunc {
	return func(c *gin.Context) {
		area1 := c.Query("area1")

		// Create a new driver for Neo4j
		driver, err := neo4j.NewDriverWithContext(Neo4j, neo4j.BasicAuth(Neo4j_User, Neo4j_Password, ""))
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			if err := driver.Close(context.Background()); err != nil {
				log.Printf("Error closing Neo4j driver: %v", err)
			}
		}()

		// Create a new session
		session := driver.NewSession(context.Background(), neo4j.SessionConfig{DatabaseName: "location"})
		defer func() {
			if err := session.Close(context.Background()); err != nil {
				log.Printf("Error closing Neo4j session: %v", err)
			}
		}()

		// Run the query to find areas with REQUESTED relationship
		result, err := session.ExecuteRead(context.Background(),
			func(tx neo4j.ManagedTransaction) (interface{}, error) {
				query := `
                    MATCH (r:AdministrativeArea1 {id: $area1})<-[:PART_OF]-(u:AdministrativeArea2)
                    RETURN u { .id } AS area2
                `
				parameters := map[string]interface{}{
					"area1": area1,
				}

				result, err := tx.Run(context.Background(), query, parameters)
				if err != nil {
					return nil, err
				}

				var areas []string
				for result.NextRecord(context.Background(), nil) {
					userNode, ok := result.Record().Get("area2")
					if !ok {
						return nil, errors.New("failed to get user node")
					}

					userMap, ok := userNode.(map[string]interface{})
					if !ok {
						return nil, errors.New("failed to convert user node to map")
					}

					id, ok := userMap["id"].(string)
					if !ok {
						id = "" // or any default value
					}

					areas = append(areas, id)
				}

				if err := result.Err(); err != nil {
					return nil, err
				}

				return areas, nil
			},
		)

		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if areas, ok := result.([]string); ok && len(areas) == 0 {
			c.JSON(http.StatusOK, []string{})
			return
		}

		// Return the list of areas with matching structure
		c.JSON(http.StatusOK, result)
	}
}
