const http = require("http");
const fs = require("fs");
const path = require("path");

const HOST = process.env.SWITCH_VM_HOST || "127.0.0.1";
const PORT = Number(process.env.SWITCH_VM_PORT || 3000);
const ROOT = __dirname;

const mime = {
  ".html": "text/html; charset=utf-8",
  ".js": "text/javascript; charset=utf-8",
  ".css": "text/css; charset=utf-8",
  ".json": "application/json; charset=utf-8"
};

const server = http.createServer((req, res) => {
  const requestPath = req.url.split("?")[0];

  if (requestPath === "/api/status") {
    res.writeHead(200, { "Content-Type": mime[".json"] });
    res.end(JSON.stringify({
      service: "OmniNX Switch Virtual Machine",
      status: "ready",
      mode: "virtual-machine",
      platform: "Nintendo Switch-style",
      note: "Simulation interface; not Nintendo firmware or hardware emulation."
    }));
    return;
  }

  const relative = requestPath === "/" ? "index.html" : requestPath.replace(/^\/+/, "");
  const file = path.resolve(ROOT, relative);

  if (!file.startsWith(ROOT + path.sep)) {
    res.writeHead(403);
    res.end("Forbidden");
    return;
  }

  fs.readFile(file, (err, data) => {
    if (err) {
      res.writeHead(err.code === "ENOENT" ? 404 : 500, {"Content-Type":"text/plain; charset=utf-8"});
      res.end(err.code === "ENOENT" ? "Not found" : "Server error");
      return;
    }
    res.writeHead(200, {"Content-Type": mime[path.extname(file)] || "application/octet-stream"});
    res.end(data);
  });
});

server.listen(PORT, HOST, () => {
  console.log("OmniNX Switch Virtual Machine");
  console.log("Index: http://" + HOST + ":" + PORT + "/");
});
