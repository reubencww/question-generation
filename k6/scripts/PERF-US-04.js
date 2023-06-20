import { check } from "k6";
import http from "k6/http";
import { Trend } from "k6/metrics";
import ws from "k6/ws";

const get_request_response_time = new Trend("get_request_response_time", true);


export const options = {
  stages: [
    // { target: 20, duration: "1m" },
    // { target: 15, duration: "1m" },
    { target: 10, duration: "1m" },
    { target: 5, duration: "1m" },
    { target: 1, duration: "1m" },
  ],
  thresholds: {
   // http_reqs: ["count > 100"],
    http_req_failed: ["rate<0.01"], // http errors should be less than 1%
    //http_req_duration: ["p(99)<5000"], // 99% of requests should be below 3s
    get_request_response_time: ["p(99)<5000"], // 99% of requests should be below 15s
  },
};
const binFile = open("/scripts/assets/lemon.jpg", "b");

export default function () {
  const url = `http://${__ENV.CONTROL_PLANE}/api/v1/challenge`;
  const ws_url = `ws://${__ENV.CONTROL_PLANE}/ws`;
  const ws_res = ws.connect(ws_url, null, function (socket) {
    const data = {
      image: http.file(binFile, "lemon.jpg"),
    };
    const res = http.post(url, data);
    const res_json = res.json();
    socket.on("open", () => console.log("WebSocket connection established"));
    socket.on("message", (event) => {
      let data = JSON.parse(event);
      if (match_id(data.id)) {
        console.log(`${data.id} responded.`);
        const created_at = new Date(data.created_at).getTime();
        const updated_at = new Date(data.updated_at).getTime();
        // console.log(`Created at ${created_at}`);
        // console.log(`Updated at ${updated_at}`);
        get_request_response_time.add(updated_at - created_at);
        console.log(`Response Time at ${updated_at - created_at} for ${data.id}`);
        socket.close();
        check(res, {
          "is status 201:created": (r) => r.status === 201,
        });
      }
    });
    socket.on("close", () => console.log("disconnected"));
    socket.on("error", function (e) {
      if (e.error() != "websocket: close sent") {
        console.log("An unexpected error occured: ", e.error());
      }
    });
    const match_id = (id) => {
      return id === res_json.data.id;
    };
  });

  check(ws_res, { "status is 101": (r) => r && r.status === 101 });
}
