# User Relation Backend
A Go-based microservice for managing user relationships, restaurant follows, and feedback in a social dining application. This service uses MongoDB for data persistence, Neo4j for relationship graphs, and Kafka for event streaming.

## Overview
This backend service handles:

- User follow/unfollow relationships with privacy controls (public/private accounts)
- Follow requests for private accounts
- Restaurant following/unfollowing
- Restaurant feedback and ratings
- Mutual follower discovery
- Real-time updates via Kafka event streaming

## Tech Stack
- Language: Go 1.23+
- Web Framework: Gin
- Databases:
    - MongoDB (user data, restaurant data, orders, feedback)
    - Neo4j (relationship graphs)
- Message Queue: Apache Kafka
- Authentication: JWT tokens
- Deployment: Docker, AWS ECS Fargate

## Key Features
### User Relationships
- Follow/Unfollow: Users can follow/unfollow other users
- Private Accounts: Support for follow requests on private accounts
- Request Management: Accept, decline, or cancel follow requests
- Mutual Followers: Discover mutual connections between users
- Follower/Following Lists: View followers, following, and pending requests

### Restaurant Features
- Follow Restaurants: Users can follow restaurants for updates
- Feedback System: Submit ratings (1-5 stars) and written reviews
- Rating Aggregation: Automatic calculation of average ratings at restaurant and location levels
- Review Management: View all feedback with user information
- Star Distribution: Get breakdown of ratings by star count

## Installation
Prerequisites
- Go 1.23 or higher
- MongoDB
- Neo4j
- Apache Kafka

