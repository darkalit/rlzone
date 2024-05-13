import React from "react";
import "./Header.style.css";
import { useLocation, Link } from "react-router-dom";
import Button, { ButtonType } from "../../UI/Button/Button";

import LogoMain from "../../Images/LogoMain.svg";
import MagnifyingGlass from "../../Images/MagnifyingGlass.svg";
import ProfileIcon from "../../Images/profileIcon.svg";
import CartBlock from "../../Images/CartBlock.svg";
import Credits from "../../Images/Credits.svg";
// import { ReactComponent as LogoMain } from "../../Images/LogoMain.svg"

function UserView() {
  let path = useLocation().pathname;
  console.log(path);
  const selection = "header-section-selected";

  return (
    <>
      <div className="header-buy-sell">
        <div
          className={`header-section ${
            path.startsWith("/buy") ? selection : ""
          }`}
        >
          <Link to="/buy">Buy</Link>
        </div>
        <div
          className={`header-section ${
            path.startsWith("/inventory") ? selection : ""
          }`}
        >
          <Link to="/inventory">Inventory</Link>
        </div>
      </div>
      <div className="header-right-block">
        <div className="header-block-settings button-add-scale">
          <a className="header-svg-icons" href="basket.html">
            <img src={CartBlock} alt="Cart" />
          </a>
        </div>
        <div className="header-block-settings">
          <img
            src={Credits}
            alt="Credits"
            className="credits"
          />
          <p className="header-main-text">1290</p>
        </div>
        <div className="header-block-settings button-add-scale">
          <a className="header-svg-icons" href="profile.html">
            <img
              src={ProfileIcon}
              alt="profileIcon"
              className="profile-pic"
            />
          </a>
          <a href="profile.html">
            <p className="header-main-text">SampleName</p>
          </a>
        </div>
      </div>
    </>
  );
}

function AdminView() {
  let path = useLocation().pathname;
  const selection = "header-section-selected";

  return (
    <>
      <div className="header-buy-sell">
        <div
          className={`header-section ${
            path.startsWith("/listings") ? selection : ""
          }`}
        >
          <Link to="/listings">Edit listings</Link>
        </div>
        <div
          className={`header-section ${
            path.startsWith("/users") ? selection : ""
          }`}
        >
          <Link to="/users">User control</Link>
        </div>
      </div>
      <div className="header-right-block">
        <div className="header-block-settings button-add-scale">
          <Link to="/profile" className="header-svg-icons">
            <img
              src={ProfileIcon}
              alt="profileIcon"
              className="profile-pic"
            />
          </Link>
          <p className="header-main-text">AdminUser</p>
        </div>
      </div>
    </>
  );
}

function GuestView() {
  return (
    <div className="header-btns">
      <Button text="Login" type={ButtonType.Outline} width="106px" />
      <Button text="Sign Up" type={ButtonType.Fill} width="106px" />
    </div>
  );
}

function AuthView() {
  return (
    <img
      src={LogoMain}
      className="header-logo"
      alt="logo"
      style={{margin: "0 auto"}}
    />
  );
}

export default function Header(props) {
  let view = null;

  switch (props.role) {
    case "User":
      view = UserView();
      break;
    case "Admin":
      view = AdminView();
      break;
    case "Auth":
      view = AuthView();
      break;
    default:
      view = GuestView();
      break;
  }

  return (
    <header className="header">
      <div className="header-content">
        <a href="index.html">
          <img
            src={LogoMain}
            className="header-logo"
            alt="logo"
          />
        </a>
        <div className="header-block-search">
          <img src={MagnifyingGlass} alt="magnifyingGlass" />
          <input
            className="header-search"
            type="text"
            placeholder="Find an item..."
          />
          <button type="button" className="header-search-button">
            ➜
          </button>
        </div>
        {view}
      </div>
    </header>
  );
}
