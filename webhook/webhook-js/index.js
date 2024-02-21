const express = require("express");
const crypto = require("crypto");
const bodyParser = require("body-parser");

const app = express();

// What you shouldn't do (at least to validate the signature)
// app.use(bodyParser.json());
// bodyParser.json will JSON.Parse the body, thus altering the original byte array
// It actually works for most events, but JSON.Parse parses URLs differently
// than what we use, failing the signature validation
app.use(bodyParser.raw({ type: "*/*" }));

const SECRET = process.env.SECRET;

if (!SECRET) {
  throw new Error("Secret variable not set or empty.");
}

function isValidSignature(headerSignature) {
  const signature = Buffer.from(headerSignature, "hex");

  console.log("Received signature: " + headerSignature);

  const hasher = crypto.createHmac("sha256", SECRET);
  hasher.update(bodyBytes);

  const expected = hasher.digest();

  console.log("Expected value: " + expected.toString("hex"));

  if (!crypto.timingSafeEqual(expected, signature)) {
    console.log("Invalid signature");
    return false;
  }

  return true
}

app.post("/webhook", (req, res) => {
  const bodyBytes = req.body;
  const headerSignature = req.get("X-IFood-Signature");

  if (isValidSignature(headerSignature)) {
    const body = bodyBytes.toString("hex");
    console.log("Received message:", bodyBytes);

    messages[bodyBytes.eventId] = bodyBytes;

    res.status(202).set("Content-Type", "application/json").send(body);
  } else {
    res.status(401).set("Content-Type", "application/json").send({
      status: "UNAUTHORIZED",
      message: "Invalid signature.",
    });
  }
});

const port = 8080;
app.listen(port, () => {
  console.log(`Server running on port ${port}`);
});
