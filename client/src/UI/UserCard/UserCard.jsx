import React from "react";
import UserPhoto from "../../Images/userPhoto.svg";
import Credits from "../../Images/Credits.svg";
import Button, { ButtonType } from "../Button/Button";

export default function UserCard({
  username,
  email,
  credits,
  date,
  onBanClick,
}) {
  return (
    <div className="main-content-sidebar">
      <div className="card grey-container">
        <div className="main-user-photo">
          <img src={UserPhoto} alt="userPhoto" />
        </div>
        <h1 className=" car-name">{username}</h1>
        <div className="user-info-block">
          <div className="info-block">
            <p className="main-text">Email:</p>
            <p className="main-text">{email}</p>
          </div>
          <div className="info-block">
            <p className="main-text">Credits:</p>
            <img
              src={Credits}
              alt="credits"
              style={{ width: "20px; height: 20px" }}
            />
            <p className="main-text">{credits}</p>
          </div>
          <div className="info-block">
            <p className="main-text">Creation date:</p>
            <p className="main-text">{date}</p>
          </div>
        </div>
        <div className="btn-block" style={{ marginTop: "64px" }}>
          <Button text="Ban user" type={ButtonType.Ban} onClick={onBanClick} />
        </div>
      </div>
    </div>
  );
}
