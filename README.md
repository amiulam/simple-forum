## API Endpoints

### Authentication
- `POST /memberships/sign-up` - Register new user
- `POST /memberships/login` - User login
- `POST /memberships/refresh-token` - Refresh access token

### Posts
- `GET /posts` - Get paginated list of posts
- `POST /posts/create` - Create new post
- `GET /posts/:postID` - Get post details
- `POST /posts/comment/:postID` - Add comment to post
- `PUT /posts/user_activity/:postID` - Like/unlike post

## Configuration

The application uses a YAML configuration file located at `internal/configs/config.yaml`. Configurable settings include:

- Service port
- Database connection
- JWT secret key
- Logging settings

Example configuration:

```yaml:README.md
service:
  port: ":9876"
  secretJWT: "yoursecret"

database:
  dataSourceName: "<username>:<password>@tcp(<host>:<port>)/<dbname>?parseTime=true"
```

## Development

### Prerequisites
- Go 1.23.2 or higher
- MySQL
- Make (for running migration commands)

### Database Migrations
Create new migration:
```bash
make migrate-create name=migration_name
```

Apply migrations:
```bash
make migrate-up
```

Rollback migrations:
```bash
make migrate-down
```

## Project Architecture

The project follows a clean architecture pattern with the following layers:
- Handlers (HTTP request handling)
- Services (Business logic)
- Repositories (Data access)
- Models (Data structures)

## Security

- JWT-based authentication
- Password hashing using bcrypt
- Refresh token mechanism for session management
- Middleware for protected routes
