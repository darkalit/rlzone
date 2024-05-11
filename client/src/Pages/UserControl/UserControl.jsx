import React from "react";
import "./UserControl.style.css";
import "../../UI/GreyContainer/GreyContainer.style.css";
import Button, { ButtonType } from "../../UI/Button/Button";

import MagnifyingGlass from "../../Images/MagnifyingGlass.svg";
import UserPhoto from "../../Images/userPhoto.svg";
import Credits from "../../Images/Credits.svg";
import { ArrayTable, ObjectTable } from "../../UI";

export default function UserControl() {
  const obj = {
    ID: 35,
    Name: "Honda Civic Type R",
    Type: "Body",
    Quality: "Limited",
    Hitbox: "Octane",
    Reactive: false,
    TradeIn: false,
    Paintable: true,
    Blueprints: false,
    Released: "7th September 2022",
    Platform: "All",
    Sideswipe: "Not available",
    Series: "-",
  };

  const head = {
    ID: function () {
      console.log("1");
    },
    Username: function () {
      console.log("2");
    },
    Email: function () {
      console.log("3");
    },
    Credits: function () {
      console.log("4");
    },
    "Creation date": function () {
      console.log("5");
    },
  };

  const users = [
    {
      ID: 1,
      Name: "Armadillo",
      Email: "examplemail@gmail.com",
      Credits: false,
      Created: "17th February 2016",
    },
    {
      ID: 2,
      Name: "Armadillo",
      Email: "examplemail@gmail.com",
      Credits: false,
      Created: "17th February 2016",
    },
    {
      ID: 3,
      Name: "Armadillo",
      Email: "examplemail@gmail.com",
      Credits: false,
      Created: "17th February 2016",
    },
    {
      ID: 4,
      Name: "Armadillo",
      Email: "examplemail@gmail.com",
      Credits: false,
      Created: "17th February 2016",
    },
  ];

  return (
    <>
      <main className="main">
        <div className="main-content">
          <div className="main-content-block-search">
            <div className="main-content-block-searchbar">
              <img src={MagnifyingGlass} alt="magnifyingGlass" />
              <input
                className="header-search"
                type="text"
                placeholder="Find user..."
              />
              <button type="button" className="header-search-button">
                ➜
              </button>
            </div>
          </div>
          <div className="main-content-block-main">
            <div className="main-content-block">
              <div className="main-description-table grey-container">
                <ArrayTable
                  data={users}
                  // side={side}
                  head={head}
                  className="main-table"
                />
              </div>
              <div className="main-description-table grey-container">
                <ObjectTable data={obj} />
              </div>
            </div>
            <div className="main-content-sidebar">
              <div className="card grey-container">
                <div className="main-user-photo">
                  <img src={UserPhoto} alt="userPhoto" />
                </div>
                <h1 className=" car-name">ExampleUserName</h1>
                <div className="user-info-block">
                  <div className="info-block">
                    <p className="main-text">Email:</p>
                    <p className="main-text">examplemail@gmail.com</p>
                  </div>
                  <div className="info-block">
                    <p className="main-text">Credits:</p>
                    <img
                      src={Credits}
                      alt="credits"
                      style={{ width: "20px; height: 20px" }}
                    />
                    <p className="main-text">1488</p>
                  </div>
                  <div className="info-block">
                    <p className="main-text">Creation date:</p>
                    <p className="main-text">2024-03-28 16:38:11</p>
                  </div>
                </div>
                <div className="btn-block" style={{ marginTop: "64px" }}>
                  <Button text="Ban user" type={ButtonType.Ban} />
                </div>
              </div>
            </div>
          </div>
        </div>
      </main>
    </>
  );
}
