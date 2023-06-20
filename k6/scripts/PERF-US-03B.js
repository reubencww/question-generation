import { check } from "k6";
import http from "k6/http";

export const options = {
  scenarios: {
    contacts: {
      executor: "constant-arrival-rate",
      // Our test should last 30 seconds in total
      duration: "5m",
      // It should start 6 iterations per `timeUnit`. Note that iterations starting points
      // will be evenly spread across the `timeUnit` period.
      rate: 40,
      // It should start `rate` iterations per second
      timeUnit: "1s",
      // It should preallocate 2 VUs before starting the test
      preAllocatedVUs: 2,
      // It is allowed to spin up to 50 maximum VUs to sustain the defined
      // constant arrival rate.
      maxVUs: 50,
    },
  },
  thresholds: {
    http_reqs: ["count > 2400"],
    http_req_failed: ["rate<0.01"], // http errors should be less than 1%
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
