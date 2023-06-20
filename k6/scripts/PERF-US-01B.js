import { check } from "k6";
import http from "k6/http";

/**
 * 1800 image uploads per hour
 * 30 image uploads per minute
 *
 */
export const options = {
  scenarios: {
    challenge: {
      executor: "constant-arrival-rate",
      // Our test should last 30 seconds in total
      duration: "60s",
      // It should start 6 iterations per `timeUnit`. Note that iterations starting points
      // will be evenly spread across the `timeUnit` period.
      rate: 30,
      // It should start `rate` iterations per second
      timeUnit: "5s",
      // It should preallocate 2 VUs before starting the test
      preAllocatedVUs: 5,
      // It is allowed to spin up to 50 maximum VUs to sustain the defined
      // constant arrival rate.
      maxVUs: 50,
    },
  },
};

const binFile = open("/scripts/assets/lemon.jpg", "b");

export default function () {
  const url = `http://${__ENV.CONTROL_PLANE}/api/v1/challenge/`;
  const data = {
    image: http.file(binFile, "lemon.jpg"),
  };

  const res = http.post(url, data);

  check(res, {
    "is status 201:created": (r) => r.status === 201,
  });
}
