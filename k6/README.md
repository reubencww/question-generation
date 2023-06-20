# Test Infra

## Setup

```bash
# Spin up timescaledb and grafana container
docker-compose up -d timescaledb grafana

# Running all tests, For windows user, use the windows version.
# Windows version: Fixes Git bash adding ;c to path
./run-all-tests.sh
./run-all-tests-windows.sh

# Run example script, for windows user, use the windows version.
./run-test.sh example.js
./run-test-windows.sh example.js

./run-test.sh PERF-US-01A.js
./run-test-windows.sh PERF-US-01A.js
```

View the dashboard in http://localhost:3333/dashboards.

In General, open Test Runs List to see the runs.

---

## Playbook

| Requirement | Title                                                     | Description                                                                                                                                                                            |
| ----------- | --------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| PERF-US-01A | Web Portal File Upload Image Post Request - Response Time | When the system is unloaded, the file must be uploaded in less than 3 seconds for at least 99% of requests coming from machines connected to the local area network.                   |
| PERF-US-01B | Web Portal File Upload Image Post Request - Throughput    | In 99% of all cases, the web server shall have a throughput of 1800 image uploads per hour under maximal load without any uploads being lost.                                          |
| PERF-US-02A | Web Portal Exercise Database Get Request - Response Time  | In 99% of the cases, all completed jobs must be sent from the workers to gRPC endpoint and the data from the database must be sent to the web portal at most 5 seconds.                |
| PERF-US-02B | Web Portal Exercise Database Get Request - Throughput     | In 99% of the cases, the throughput of the request send to the database by the webserver shall be 1,800 requests per hour under maximal load without any push notification being lost. |
