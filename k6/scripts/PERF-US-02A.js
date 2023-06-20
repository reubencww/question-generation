import { check } from "k6";
import http from "k6/http";

export const options = {
  stages: [
    { target: 20, duration: "1m" },
    { target: 15, duration: "1m" },
    { target: 0, duration: "1m" },
  ],
  thresholds: {
    http_reqs: ["count > 100"],
    http_req_failed: ["rate<0.01"], // http errors should be less than 1%
    checks: ["p(99)<10 000"], // 99% of requests should be below 3s
  },
};

export default function () {
  const url = `http://${__ENV.CONTROL_PLANE}/api/v1/challenge/1`;
  const res = http.get(url);

  check(res, {
    "status is 200": (r) => r.status === 200,
  });

  
}