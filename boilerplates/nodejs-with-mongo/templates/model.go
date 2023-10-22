package templates

const ModelTemplate = `
const mongoose = require("mongoose");
const Schema = mongoose.Schema;

const StartSchema = new Schema({
  name: {
	type: String,
	required: true,
	default: "go-furnace",
  },
  happyNumber: {
	type: Number,
	required: true,
	default: 42,
  }
});

module.exports = mongoose.model("Start", StartSchema);
`
