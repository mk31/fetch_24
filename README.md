# RUN INSTRUCTIONS
Step 0  
Make sure you have nothing running on localhost port 80  
If you do, you will see an error similar to:
```[GIN-debug] [ERROR] listen tcp :80: bind: address already in use```

<!-- this needs update - once we add packages -->
Local (requires Go is installed)
1. Clone Repo
2. From root directory of repo `go run .`

Docker
1. Clone repo 
2. Go to root directory of repo
3. docker build -t fetch_24 .
4. docker run -p 80:80 fetch_24

process receipt route
localhost/api/v1/receipts/process

get points route
localhost/api/v1/receipts/:id/points

### TODOs
Standardize error handling 
Real api would have auth
Log errors to some trackable location - slack for example
Create a common library/package for code re-use
CI/CD
Instrument APM for tracing
Create DAL/Repository for real data storage
