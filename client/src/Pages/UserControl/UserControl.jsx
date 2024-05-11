import React from "react";
import "./UserControl.style.css";
import "../../UI/GreyContainer/GreyContainer.style.css";
import Button, { ButtonType } from "../../UI/Button/Button";

import MagnifyingGlass from "../../Images/MagnifyingGlass.svg";
import UserPhoto from "../../Images/userPhoto.svg";
import Credits from "../../Images/Credits.svg";

export default function UserControl() {
  // const styles = {
  //   background: "red",
  //   height: "100px",
  //   width: "100px",
  // };

  return (
    <>
      <main className="main">
        <div className="main-content">
          <div className="main-content-block-search">
            <div className="main-content-block-searchbar">
              <img
                src={MagnifyingGlass}
                alt="magnifyingGlass"
              />
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
                <table className="main-table">
                  <thead>
                    <tr>
                      <th className="table-top-row">ID</th>
                      <th className="table-top-row">Username</th>
                      <th className="table-top-row">Email</th>
                      <th className="table-top-row">Credits</th>
                      <th className="top-right-row table-top-row">
                        Creation date
                      </th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr>
                      <td className="table-left-column">1</td>
                      <td>ExampleUserName</td>
                      <td>examplemail@gmail.com</td>
                      <td>10</td>
                      <td>2024-03-28</td>
                    </tr>
                    <tr>
                      <td className="table-left-column">2</td>
                      <td>ExampleUserName</td>
                      <td>examplemail@gmail.com</td>
                      <td>20</td>
                      <td>2024-03-28</td>
                    </tr>
                    <tr>
                      <td className="table-left-column">3</td>
                      <td>ExampleUserName</td>
                      <td>examplemail@gmail.com</td>
                      <td>15</td>
                      <td>2024-03-28</td>
                    </tr>
                    <tr>
                      <td className="table-left-column">4</td>
                      <td>ExampleUserName</td>
                      <td>examplemail@gmail.com</td>
                      <td>15</td>
                      <td>2024-03-28</td>
                    </tr>
                    <tr>
                      <td className="table-left-column">5</td>
                      <td>ExampleUserName</td>
                      <td>examplemail@gmail.com</td>
                      <td>15</td>
                      <td>2024-03-28</td>
                    </tr>
                    <tr>
                      <td className="table-left-column">6</td>
                      <td>ExampleUserName</td>
                      <td>examplemail@gmail.com</td>
                      <td>15</td>
                      <td>2024-03-28</td>
                    </tr>
                    <tr>
                      <td className="table-left-column">7</td>
                      <td>ExampleUserName</td>
                      <td>examplemail@gmail.com</td>
                      <td>15</td>
                      <td>2024-03-28</td>
                    </tr>
                    <tr>
                      <td className="table-left-column">8</td>
                      <td>ExampleUserName</td>
                      <td>examplemail@gmail.com</td>
                      <td>15</td>
                      <td>2024-03-28</td>
                    </tr>
                    <tr>
                      <td className="table-left-column">9</td>
                      <td>ExampleUserName</td>
                      <td>examplemail@gmail.com</td>
                      <td>15</td>
                      <td>2024-03-28</td>
                    </tr>
                    <tr>
                      <td className="table-left-column">10</td>
                      <td>ExampleUserName</td>
                      <td>examplemail@gmail.com</td>
                      <td>15</td>
                      <td>2024-03-28</td>
                    </tr>
                    <tr>
                      <td className="table-left-column">11</td>
                      <td>ExampleUserName</td>
                      <td>examplemail@gmail.com</td>
                      <td>15</td>
                      <td>2024-03-28</td>
                    </tr>
                    <tr>
                      <td className="table-left-column">12</td>
                      <td>ExampleUserName</td>
                      <td>examplemail@gmail.com</td>
                      <td>15</td>
                      <td>2024-03-28</td>
                    </tr>
                    <tr>
                      <td className="table-left-column">13</td>
                      <td>ExampleUserName</td>
                      <td>examplemail@gmail.com</td>
                      <td>15</td>
                      <td>2024-03-28</td>
                    </tr>
                    <tr>
                      <td className="table-left-column">14</td>
                      <td>ExampleUserName</td>
                      <td>examplemail@gmail.com</td>
                      <td>15</td>
                      <td>2024-03-28</td>
                    </tr>
                    <tr>
                      <td className="table-left-column">15</td>
                      <td>ExampleUserName</td>
                      <td>examplemail@gmail.com</td>
                      <td>15</td>
                      <td>2024-03-28</td>
                    </tr>
                  </tbody>
                </table>
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
                      style={{width: "20px; height: 20px"}}
                    />
                    <p className="main-text">1488</p>
                  </div>
                  <div className="info-block">
                    <p className="main-text">Creation date:</p>
                    <p className="main-text">2024-03-28 16:38:11</p>
                  </div>
                </div>
                <div className="btn-block" style={{marginTop: "64px"}}> 
                  <Button text="Ban user" type={ButtonType.Outline} border-color="var(--misc-failure)" />
                </div>
              </div>
            </div>
          </div>
        </div>
      </main>
    </>
  );
}
