import { check } from "k6";
import http from "k6/http";

export const options = {
    stages: [
         { duration: '2m', target: 300 }, // below normal load
         { duration: '5m', target: 300 },
         { duration: '2m', target: 400 }, // normal load
         { duration: '5m', target: 400 },
         { duration: '2m', target: 600 }, // around the breaking point
         { duration: '5m', target: 600 },
         { duration: '2m', target: 400 }, // beyond the breaking point
         { duration: '5m', target: 300 },
         { duration: '10m', target: 0 }, // scale down. Recovery stage.
         ],
  thresholds: {
    http_reqs: ["count > 1800"],
    http_req_failed: ["rate<0.01"], // http errors should be less than 1%
  },
};

export default function () {
  const url = `http://${__ENV.CONTROL_PLANE}/api/v1/challenge/1`;
  const res = http.get(url);

  check(res, {
    "status is 200": (r) => r.status === 200,
  });
}
