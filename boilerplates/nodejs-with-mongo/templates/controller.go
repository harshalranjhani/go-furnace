package templates

const ControllerTemplate = `
const Start = require("../models/start");

module.exports.displayStartMessage = (req, res) => {
  res.send("Congratulations! Your go-furnace template for ExpressJS and MongoDB is ready and you can begin working!.");
};

module.exports.sendStartMessage = async (req, res) => {
	  const newStart = new Start();
	  await newStart.save();
	  res.json({ data: newStart})
};

`
