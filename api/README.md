## Usage

1. Install dependencies with go :
```bash
go mod download
```
2. Create the .env file
```bash
cp .env.example .env
```
3. And run:
```bash
go run cmd/main.go
```
You can visit http://localhost:4003/status to check that the service is running

### Build

To build the App, run

```bash
pnpm build
```