package templates

const RouteTemplate = `
const express = require("express");
const router = express.Router();
const Start = require("../models/start");
const start = require("../controllers/start");

router
  .route("/")
  .get(start.displayStartMessage)
  .post(start.sendStartMessage);

module.exports = router;

`
