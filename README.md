# Url Shortener API
A simple and scalable URL Shortener API built with Go, Gin, and GORM. It allows users to shorten long URLs, manage their shortened links (basically just delete it in any time), and track amount of redirects.
## Features
- User registration and login
- Shorten long URLs
- Click tracking (redirect count)
- JWT-based authentification (access and refresh tokens)
- Tokens Cleaner - revoke expired tokens and delete tokens which expired twice. Recommended to run periodically (e.g. via cron)
## Techs used
- Go (Golang)
- Gin web framework
- GORM ORM
- PostgreSQL
## Getting started
### Prerequisites
- Go 1.24+
- PostgreSQL
### Installation
1. Clone the repo:
```bash
git clone https://github.com/wrtgvr2/urlshrt.git
cd urlshrt
```
2. Configure your environment variables:
Create a .env file and add the necessary variables:
```env
JWT_SECRET=your_jwt_secret
# url salt used as part of generating short url
URL_SECRET_SALT=your_url_secret_salt 

DB_HOST=your_postgres_host
DB_NAME=your_database_name
DB_PORT=your_database_port
DB_USER=your_postgres_user
DB_PASSWORD=your_database_password
```
3. Install dependencies:
```bash
go mod tidy
```
4.1 To start API:
```bash
go run ./cmd/shortener
```
4.2 To start token cleaner:
```bash
go run ./cmd/tokensCleaner
```
## API endpoints:
| Method | Endpoint          | Description                          | Response                 | Auth required |
|--------|-------------------|--------------------------------------|--------------------------|---------------|
| POST   | /auth/login       | User login                           | Response with JWT tokens | No            |
| POST   | /auth/register    | User registration                    | Response with user model | No            |
| POST   | /auth/refresh     | Refresh JWT tokens                   | Response with JWT tokens | Yes           |
| GET    | /r/{ShortURL}     | Redirect to orig URL                 | -                        | No            |
| GET    | /api/users        | Get loged user info                  | Response with user model | Yes           |
| PATCH  | /api/users        | Update loged user info               | Response with user model | Yes           |
| DELETE | /api/users        | Delete user from DB                  | Status code only         | Yes           |
| GET    | /api/urls         | Get all loged user's short URLs      | Response with url models | Yes           |
| GET    | /api/urls/{URLid} | Get loged user url with given url id | Response with url model  | Yes           |
| DELETE | /api/urls/{URLid} | Delete url with given url id         | Status code only         | Yes           |
|--------|-------------------|--------------------------------------|--------------------------|---------------|
