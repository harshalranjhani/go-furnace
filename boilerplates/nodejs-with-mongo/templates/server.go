package templates

const ServerTemplate = `
  const express = require("express");
  const app = express();
  const mongoose = require("mongoose");
  const startRoutes = require("./routes/start");

  const mongodb = require("mongodb");
  const dbUrl = "mongodb://localhost:27017/go-furnace";
  
  mongoose
	.connect(dbUrl, {
	  useNewUrlParser: true,
	  useUnifiedTopology: true,
	})
	.then(() => {
	  console.log("MONGO CONNECTION OPEN!!!");
	})
	.catch((err) => {
	  console.log("OH NO MONGO CONNECTION ERROR!!!!");
	  console.log(err);
	});
  
  // middleware
  app.use(express.urlencoded({ extended: true }));
    
  app.get("/", (req, res) => {
	res.send("go-furnace template for ExpressJS and MongoDB");
  });

  app.use("/start", startRoutes);
  
  const PORT = 8080;
  app.listen(PORT, () => {
	console.log("Your app is running on port 8080...");
  });`
