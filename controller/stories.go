package controller

// import (
//     "log"
//     "time"

//     "github.com/neo4j/neo4j-go-driver/v4/neo4j"
// )

// func main() {
//     // Set up a ticker to run the job every hour
//     ticker := time.NewTicker(1 * time.Hour)
//     defer ticker.Stop()

//     for {
//         select {
//         case <-ticker.C:
//             err := deleteOldStories()
//             if err != nil {
//                 log.Printf("Failed to delete old stories: %v", err)
//             }
//         }
//     }
// }

// func deleteOldStories() error {
//     // Create a new driver for Neo4j
//     driver, err := neo4j.NewDriver("neo4j://localhost:7687", neo4j.BasicAuth("neo4j", "12345678", ""))
//     if err != nil {
//         return err
//     }
//     defer driver.Close()

//     // Create a new session
//     session := driver.NewSession(neo4j.SessionConfig{})
//     defer session.Close()

//     // Run the query to delete stories older than 24 hours
//     _, err = session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
//         result, err := transaction.Run(
//             "MATCH (s:Story) WHERE s.createdAt < datetime({hours: -24}) DETACH DELETE s",
//             nil,
//         )
//         if err != nil {
//             return nil, err
//         }
//         return result.Consume()
//     })

//     return err
// }
