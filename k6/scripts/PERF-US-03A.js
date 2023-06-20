import { check } from "k6";
import http from "k6/http";

export const options = {
  stages: [
    { target: 20, duration: "1m" },
    { target: 15, duration: "1m" },
    { target: 1, duration: "1m" },
  ],
  thresholds: {
    http_req_failed: ["rate<0.01"], // http errors should be less than 1%
    //http_req_duration: ["p(99)<3000"], // 99% of requests should be below 3s
  },
};

export default function () {
  const url = `http://${__ENV.CONTROL_PLANE}/api/v1/challenge/question/1`;

  const question = "What kind of may is on a plate?";
  const answer = "twooo";

  const data = {
    question,
    answer,
  };
  const res = http.patch(url, data);
  //console.log(res);
  check(res, {
    "status is 200": (r) => r.status === 200,
  });
}
